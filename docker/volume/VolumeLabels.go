// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/volume#label Volume#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/volume#value Volume#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

