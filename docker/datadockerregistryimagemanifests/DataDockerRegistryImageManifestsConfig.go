// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadockerregistryimagemanifests

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDockerRegistryImageManifestsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Docker image, including any tags. e.g. `alpine:latest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/data-sources/registry_image_manifests#name DataDockerRegistryImageManifests#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/data-sources/registry_image_manifests#auth_config DataDockerRegistryImageManifests#auth_config}
	AuthConfig *DataDockerRegistryImageManifestsAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/data-sources/registry_image_manifests#id DataDockerRegistryImageManifests#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/data-sources/registry_image_manifests#insecure_skip_verify DataDockerRegistryImageManifests#insecure_skip_verify}
	InsecureSkipVerify interface{} `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
}

