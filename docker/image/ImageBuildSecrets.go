// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package image


type ImageBuildSecrets struct {
	// ID of the secret. By default, secrets are mounted to /run/secrets/<id>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.3.0/docs/resources/image#id Image#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Environment variable source of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.3.0/docs/resources/image#env Image#env}
	Env *string `field:"optional" json:"env" yaml:"env"`
	// File source of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.3.0/docs/resources/image#src Image#src}
	Src *string `field:"optional" json:"src" yaml:"src"`
}

