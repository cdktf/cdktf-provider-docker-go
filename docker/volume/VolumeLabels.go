package volume


type VolumeLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/volume#label Volume#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/volume#value Volume#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

