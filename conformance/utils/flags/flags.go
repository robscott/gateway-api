/*
Copyright 2022 The Kubernetes Authors.

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

// flags contains command-line flag definitions for the conformance
// tests. They're in this package so they can be shared among the
// various suites that are all run by the same Makefile invocation.
package flags

import (
	"flag"
)

const (
	// DefaultMode is the operating mode to default to in case no mode is specified.
	DefaultMode = "default"
)

var (
	GatewayClassName           = flag.String("gateway-class", "gateway-conformance", "Name of GatewayClass to use for tests")
	ShowDebug                  = flag.Bool("debug", false, "Whether to print debug logs")
	CleanupBaseResources       = flag.Bool("cleanup-base-resources", true, "Whether to cleanup base test resources after the run")
	SupportedFeatures          = flag.String("supported-features", "", "Supported features included in conformance tests suites")
	SkipTests                  = flag.String("skip-tests", "", "Comma-separated list of tests to skip")
	RunTest                    = flag.String("run-test", "", "Name of a single test to run, instead of the whole suite")
	ExemptFeatures             = flag.String("exempt-features", "", "Exempt Features excluded from conformance tests suites")
	EnableAllSupportedFeatures = flag.Bool("all-features", false, "Whether to enable all supported features for conformance tests")
	NamespaceLabels            = flag.String("namespace-labels", "", "Comma-separated list of name=value labels to add to test namespaces")
	NamespaceAnnotations       = flag.String("namespace-annotations", "", "Comma-separated list of name=value annotations to add to test namespaces")
	ImplementationOrganization = flag.String("organization", "", "Implementation's Organization")
	ImplementationProject      = flag.String("project", "", "Implementation's project")
	ImplementationURL          = flag.String("url", "", "Implementation's url")
	ImplementationVersion      = flag.String("version", "", "Implementation's version")
	ImplementationContact      = flag.String("contact", "", "Comma-separated list of contact information for the maintainers")
	Mode                       = flag.String("mode", DefaultMode, "The operating mode of the implementation.")
	AllowCRDsMismatch          = flag.Bool("allow-crds-mismatch", false, "Flag to allow the suite not to fail in case there is a mismatch between CRDs versions and channels.")
	ConformanceProfiles        = flag.String("conformance-profiles", "", "Comma-separated list of the conformance profiles to run")
	ReportOutput               = flag.String("report-output", "", "The file where to write the conformance report")
	SkipProvisionalTests       = flag.Bool("skip-provisional-tests", false, "Whether to skip provisional tests")
)
