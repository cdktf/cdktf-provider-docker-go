package container


type ContainerMountsVolumeOptionsLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#label Container#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#value Container#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

