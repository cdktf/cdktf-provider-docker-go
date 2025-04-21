// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecResources struct {
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.3.0/docs/resources/service#limits Service#limits}
	Limits *ServiceTaskSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.3.0/docs/resources/service#reservation Service#reservation}
	Reservation *ServiceTaskSpecResourcesReservation `field:"optional" json:"reservation" yaml:"reservation"`
}

