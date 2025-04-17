// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceMode struct {
	// The global service mode. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.2.0/docs/resources/service#global Service#global}
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// replicated block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.2.0/docs/resources/service#replicated Service#replicated}
	Replicated *ServiceModeReplicated `field:"optional" json:"replicated" yaml:"replicated"`
}

