// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecContainerSpecHosts struct {
	// The name of the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#host Service#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The ip of the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs/resources/service#ip Service#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

