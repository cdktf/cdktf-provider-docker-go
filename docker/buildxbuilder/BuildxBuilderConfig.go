// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BuildxBuilderConfig struct {
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
	// Append a node to builder instead of changing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#append BuildxBuilder#append}
	Append interface{} `field:"optional" json:"append" yaml:"append"`
	// Automatically boot the builder after creation. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#bootstrap BuildxBuilder#bootstrap}
	Bootstrap interface{} `field:"optional" json:"bootstrap" yaml:"bootstrap"`
	// BuildKit daemon config file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#buildkit_config BuildxBuilder#buildkit_config}
	BuildkitConfig *string `field:"optional" json:"buildkitConfig" yaml:"buildkitConfig"`
	// BuildKit flags to set for the builder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#buildkit_flags BuildxBuilder#buildkit_flags}
	BuildkitFlags *string `field:"optional" json:"buildkitFlags" yaml:"buildkitFlags"`
	// docker_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#docker_container BuildxBuilder#docker_container}
	DockerContainer *BuildxBuilderDockerContainer `field:"optional" json:"dockerContainer" yaml:"dockerContainer"`
	// The driver to use for the Buildx builder (e.g., docker-container, kubernetes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#driver BuildxBuilder#driver}
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Additional options for the Buildx driver in the form of `key=value,...`. These options are driver-specific.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#driver_options BuildxBuilder#driver_options}
	DriverOptions *map[string]*string `field:"optional" json:"driverOptions" yaml:"driverOptions"`
	// The endpoint or context to use for the Buildx builder, where context is the name of a context from docker context ls and endpoint is the address for Docker socket (eg.
	//
	// DOCKER_HOST value). By default, the current Docker configuration is used for determining the context/endpoint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#endpoint BuildxBuilder#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#id BuildxBuilder#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#kubernetes BuildxBuilder#kubernetes}
	Kubernetes *BuildxBuilderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// The name of the Buildx builder. IF not specified, a random name will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#name BuildxBuilder#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Create/modify node with given name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#node BuildxBuilder#node}
	NodeAttribute *string `field:"optional" json:"nodeAttribute" yaml:"nodeAttribute"`
	// Fixed platforms for current node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#platform BuildxBuilder#platform}
	Platform *[]*string `field:"optional" json:"platform" yaml:"platform"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#remote BuildxBuilder#remote}
	Remote *BuildxBuilderRemote `field:"optional" json:"remote" yaml:"remote"`
	// Set the current builder instance as the default for the current context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#use BuildxBuilder#use}
	Use interface{} `field:"optional" json:"use" yaml:"use"`
}

