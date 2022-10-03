package container


type ContainerMountsBindOptions struct {
	// A propagation mode with the value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#propagation Container#propagation}
	Propagation *string `field:"optional" json:"propagation" yaml:"propagation"`
}

