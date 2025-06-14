// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecPlacementPlatforms struct {
	// The architecture, e.g. `amd64`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#architecture Service#architecture}
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The operation system, e.g. `linux`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#os Service#os}
	Os *string `field:"required" json:"os" yaml:"os"`
}

