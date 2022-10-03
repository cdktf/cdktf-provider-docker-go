package secret


type SecretLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/secret#label Secret#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/secret#value Secret#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

