/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewaySpecAddressApplyConfiguration represents a declarative configuration of the GatewaySpecAddress type for use
// with apply.
type GatewaySpecAddressApplyConfiguration struct {
	Type  *v1.AddressType `json:"type,omitempty"`
	Value *string         `json:"value,omitempty"`
}

// GatewaySpecAddressApplyConfiguration constructs a declarative configuration of the GatewaySpecAddress type for use with
// apply.
func GatewaySpecAddress() *GatewaySpecAddressApplyConfiguration {
	return &GatewaySpecAddressApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *GatewaySpecAddressApplyConfiguration) WithType(value v1.AddressType) *GatewaySpecAddressApplyConfiguration {
	b.Type = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *GatewaySpecAddressApplyConfiguration) WithValue(value string) *GatewaySpecAddressApplyConfiguration {
	b.Value = &value
	return b
}
