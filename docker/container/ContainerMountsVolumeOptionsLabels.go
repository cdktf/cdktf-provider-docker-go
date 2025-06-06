// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerMountsVolumeOptionsLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/container#label Container#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/container#value Container#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

