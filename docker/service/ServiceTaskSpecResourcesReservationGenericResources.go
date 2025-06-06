// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecResourcesReservationGenericResources struct {
	// The Integer resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/service#discrete_resources_spec Service#discrete_resources_spec}
	DiscreteResourcesSpec *[]*string `field:"optional" json:"discreteResourcesSpec" yaml:"discreteResourcesSpec"`
	// The String resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/service#named_resources_spec Service#named_resources_spec}
	NamedResourcesSpec *[]*string `field:"optional" json:"namedResourcesSpec" yaml:"namedResourcesSpec"`
}

