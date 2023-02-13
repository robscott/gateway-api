/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	"sigs.k8s.io/gateway-api/apis/v1beta1"
	"sigs.k8s.io/gateway-api/conformance/utils/config"
)

// GatewayRef is a tiny type for specifying an HTTP Route ParentRef without
// relying on a specific API Version.
type GatewayRef struct {
	types.NamespacedName
	listenerNames []*v1beta1.SectionName
}

type RouteRef struct {
	types.NamespacedName
	metav1.Object
}

// NewGatewayRef creates a GatewayRef resource. ListenerNames are optional.
func NewGatewayRef(nn types.NamespacedName, listenerNames ...string) GatewayRef {
	var listeners []*v1beta1.SectionName

	if len(listenerNames) == 0 {
		listenerNames = append(listenerNames, "")
	}

	for _, listener := range listenerNames {
		sectionName := v1beta1.SectionName(listener)
		listeners = append(listeners, &sectionName)
	}
	return GatewayRef{
		NamespacedName: nn,
		listenerNames:  listeners,
	}
}

type RouteWaitOptions struct {
	timeoutConfig  config.TimeoutConfig
	controllerName string
	gw             GatewayRef
	routeRefs      []RouteRef
	routeNNs       []types.NamespacedName
	RouteKinds     []string
	Programmed     bool
}

func WaitForGatewayAndRoutes(t *testing.T, c client.Client, w RouteWaitOptions) (string, []string) {
	t.Helper()

	hostnames := []string{}

	gwAddr, err := WaitForGatewayAddress(t, c, w.timeoutConfig, w.gw.NamespacedName)
	require.NoErrorf(t, err, "timed out waiting for Gateway address to be assigned")

	ns := v1beta1.Namespace(w.gw.Namespace)
	kind := v1beta1.Kind("Gateway")

	for _, routeNN := range w.routeNNs {
		namespaceRequired := true
		if routeNN.Namespace == w.gw.Namespace {
			namespaceRequired = false
		}

		var parents []v1beta1.RouteParentStatus
		for _, listener := range w.gw.listenerNames {
			parents = append(parents, v1beta1.RouteParentStatus{
				ParentRef: v1beta1.ParentReference{
					Group:       (*v1beta1.Group)(&v1beta1.GroupVersion.Group),
					Kind:        &kind,
					Name:        v1beta1.ObjectName(w.gw.Name),
					Namespace:   &ns,
					SectionName: listener,
				},
				ControllerName: v1beta1.GatewayController(w.controllerName),
				Conditions: []metav1.Condition{
					{
						Type:   string(v1beta1.RouteConditionAccepted),
						Status: metav1.ConditionTrue,
						Reason: string(v1beta1.RouteReasonAccepted),
					},
				},
			})
		}
		RouteMustHaveParents(t, c, w.timeoutConfig, routeNN, parents, namespaceRequired)
	}

	return gwAddr, hostnames
}

// RouteMustHaveParents waits for the specified Route to have parents
// in status that match the expected parents. This will cause the test to halt
// if the specified timeout is exceeded.
func RouteMustHaveParents[R v1beta1.HTTPRoute | v1alpha2.TLSRoute](t *testing.T, client client.Client, timeoutConfig config.TimeoutConfig, routeName types.NamespacedName, parents []v1beta1.RouteParentStatus, namespaceRequired bool) {
	t.Helper()

	var actual []v1beta1.RouteParentStatus
	waitErr := wait.PollImmediate(1*time.Second, timeoutConfig.RouteMustHaveParents, func() (bool, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		route := &v1beta1.HTTPRoute{}
		err := client.Get(ctx, routeName, route)
		if err != nil {
			return false, fmt.Errorf("error fetching Route: %w", err)
		}

		for _, parent := range actual {
			if err := ConditionsHaveLatestObservedGeneration(route, parent.Conditions); err != nil {
				t.Logf("Route(controller=%v,ref=%#v) %v", parent.ControllerName, parent, err)
				return false, nil
			}
		}

		actual = route.Status.Parents
		return parentsForRouteMatch(t, routeName, parents, actual, namespaceRequired), nil
	})
	require.NoErrorf(t, waitErr, "error waiting for HTTPRoute to have parents matching expectations")
}

// func GatewayAndTLSRoutesMustBeAccepted(t *testing.T, c client.Client, timeoutConfig config.TimeoutConfig, controllerName string, gw GatewayRef, routeNNs ...types.NamespacedName) (string, []v1beta1.Hostname) {
