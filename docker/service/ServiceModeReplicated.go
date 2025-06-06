// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceModeReplicated struct {
	// The amount of replicas of the service. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/service#replicas Service#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

