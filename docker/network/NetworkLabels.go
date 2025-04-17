// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network


type NetworkLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.2.0/docs/resources/network#label Network#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.2.0/docs/resources/network#value Network#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

