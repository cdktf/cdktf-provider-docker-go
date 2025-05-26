// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registryimage


type RegistryImageAuthConfig struct {
	// The address of the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.0/docs/resources/registry_image#address RegistryImage#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The password for the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.0/docs/resources/registry_image#password RegistryImage#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username for the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.0/docs/resources/registry_image#username RegistryImage#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

