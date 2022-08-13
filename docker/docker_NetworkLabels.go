// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type NetworkLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/network#label Network#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/network#value Network#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

