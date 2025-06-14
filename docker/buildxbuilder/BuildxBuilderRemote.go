// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder


type BuildxBuilderRemote struct {
	// Absolute path to the TLS certificate authority used for validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#cacert BuildxBuilder#cacert}
	Cacert *string `field:"optional" json:"cacert" yaml:"cacert"`
	// Absolute path to the TLS client certificate to present to buildkitd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#cert BuildxBuilder#cert}
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// Automatically load images to the Docker Engine image store. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#default_load BuildxBuilder#default_load}
	DefaultLoad interface{} `field:"optional" json:"defaultLoad" yaml:"defaultLoad"`
	// Sets the TLS client key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#key BuildxBuilder#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// TLS server name used in requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#servername BuildxBuilder#servername}
	Servername *string `field:"optional" json:"servername" yaml:"servername"`
}

