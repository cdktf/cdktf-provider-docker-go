// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecContainerSpecMountsVolumeOptionsLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#label Service#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#value Service#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}
