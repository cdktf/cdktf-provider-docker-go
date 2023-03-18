package service


type ServiceTaskSpecResources struct {
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#limits Service#limits}
	Limits *ServiceTaskSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// reservation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#reservation Service#reservation}
	Reservation *ServiceTaskSpecResourcesReservation `field:"optional" json:"reservation" yaml:"reservation"`
}

