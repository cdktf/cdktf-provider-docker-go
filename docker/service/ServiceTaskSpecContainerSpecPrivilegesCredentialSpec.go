// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecContainerSpecPrivilegesCredentialSpec struct {
	// Load credential spec from this file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.4.0/docs/resources/service#file Service#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// Load credential spec from this value in the Windows registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.4.0/docs/resources/service#registry Service#registry}
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
}

