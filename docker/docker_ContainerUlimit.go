// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ContainerUlimit struct {
	// The hard limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#hard Container#hard}
	Hard *float64 `field:"required" json:"hard" yaml:"hard"`
	// The name of the ulimit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#name Container#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The soft limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#soft Container#soft}
	Soft *float64 `field:"required" json:"soft" yaml:"soft"`
}

