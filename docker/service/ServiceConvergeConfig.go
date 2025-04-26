// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceConvergeConfig struct {
	// The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.4.0/docs/resources/service#delay Service#delay}
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
	// The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.4.0/docs/resources/service#timeout Service#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

