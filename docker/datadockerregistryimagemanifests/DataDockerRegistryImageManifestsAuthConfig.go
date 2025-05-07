// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadockerregistryimagemanifests


type DataDockerRegistryImageManifestsAuthConfig struct {
	// The address of the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.5.0/docs/data-sources/registry_image_manifests#address DataDockerRegistryImageManifests#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// The password for the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.5.0/docs/data-sources/registry_image_manifests#password DataDockerRegistryImageManifests#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username for the Docker registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.5.0/docs/data-sources/registry_image_manifests#username DataDockerRegistryImageManifests#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

