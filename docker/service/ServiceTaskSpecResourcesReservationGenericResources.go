package service


type ServiceTaskSpecResourcesReservationGenericResources struct {
	// The Integer resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#discrete_resources_spec Service#discrete_resources_spec}
	DiscreteResourcesSpec *[]*string `field:"optional" json:"discreteResourcesSpec" yaml:"discreteResourcesSpec"`
	// The String resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#named_resources_spec Service#named_resources_spec}
	NamedResourcesSpec *[]*string `field:"optional" json:"namedResourcesSpec" yaml:"namedResourcesSpec"`
}

