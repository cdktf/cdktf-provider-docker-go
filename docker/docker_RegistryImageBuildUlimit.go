// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type RegistryImageBuildUlimit struct {
	// soft limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#hard RegistryImage#hard}
	Hard *float64 `field:"required" json:"hard" yaml:"hard"`
	// type of ulimit, e.g. `nofile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#name RegistryImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// hard limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#soft RegistryImage#soft}
	Soft *float64 `field:"required" json:"soft" yaml:"soft"`
}

