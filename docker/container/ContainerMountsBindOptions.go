// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerMountsBindOptions struct {
	// A propagation mode with the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.0/docs/resources/container#propagation Container#propagation}
	Propagation *string `field:"optional" json:"propagation" yaml:"propagation"`
}

