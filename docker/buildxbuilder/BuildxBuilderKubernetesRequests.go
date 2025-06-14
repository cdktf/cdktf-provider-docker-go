// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder


type BuildxBuilderKubernetesRequests struct {
	// CPU limit for the Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#cpu BuildxBuilder#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Ephemeral storage limit for the Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#ephemeral_storage BuildxBuilder#ephemeral_storage}
	EphemeralStorage *string `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Memory limit for the Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/buildx_builder#memory BuildxBuilder#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

