// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder


type BuildxBuilderKubernetes struct {
	// Sets additional annotations on the deployments and pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#annotations BuildxBuilder#annotations}
	Annotations *string `field:"optional" json:"annotations" yaml:"annotations"`
	// Automatically load images to the Docker Engine image store. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#default_load BuildxBuilder#default_load}
	DefaultLoad interface{} `field:"optional" json:"defaultLoad" yaml:"defaultLoad"`
	// Sets the image to use for running BuildKit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#image BuildxBuilder#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Sets additional labels on the deployments and pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#labels BuildxBuilder#labels}
	Labels *string `field:"optional" json:"labels" yaml:"labels"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#limits BuildxBuilder#limits}
	Limits *BuildxBuilderKubernetesLimits `field:"optional" json:"limits" yaml:"limits"`
	// Load-balancing strategy (sticky or random).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#loadbalance BuildxBuilder#loadbalance}
	Loadbalance *string `field:"optional" json:"loadbalance" yaml:"loadbalance"`
	// Sets the Kubernetes namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#namespace BuildxBuilder#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Sets the pod's nodeSelector label(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#nodeselector BuildxBuilder#nodeselector}
	Nodeselector *string `field:"optional" json:"nodeselector" yaml:"nodeselector"`
	// qemu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#qemu BuildxBuilder#qemu}
	Qemu *BuildxBuilderKubernetesQemu `field:"optional" json:"qemu" yaml:"qemu"`
	// Sets the number of Pod replicas to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#replicas BuildxBuilder#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#requests BuildxBuilder#requests}
	Requests *BuildxBuilderKubernetesRequests `field:"optional" json:"requests" yaml:"requests"`
	// Run the container as a non-root user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#rootless BuildxBuilder#rootless}
	Rootless interface{} `field:"optional" json:"rootless" yaml:"rootless"`
	// Sets the scheduler responsible for scheduling the pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#schedulername BuildxBuilder#schedulername}
	Schedulername *string `field:"optional" json:"schedulername" yaml:"schedulername"`
	// Sets the pod's serviceAccountName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#serviceaccount BuildxBuilder#serviceaccount}
	Serviceaccount *string `field:"optional" json:"serviceaccount" yaml:"serviceaccount"`
	// Set the timeout limit for pod provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#timeout BuildxBuilder#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Configures the pod's taint toleration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#tolerations BuildxBuilder#tolerations}
	Tolerations *string `field:"optional" json:"tolerations" yaml:"tolerations"`
}

