// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceEndpointSpec struct {
	// The mode of resolution to use for internal load balancing between tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#mode Service#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#ports Service#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

