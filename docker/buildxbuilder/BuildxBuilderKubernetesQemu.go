// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder


type BuildxBuilderKubernetesQemu struct {
	// Sets the QEMU emulation image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#image BuildxBuilder#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Install QEMU emulation for multi-platform support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#install BuildxBuilder#install}
	Install interface{} `field:"optional" json:"install" yaml:"install"`
}

