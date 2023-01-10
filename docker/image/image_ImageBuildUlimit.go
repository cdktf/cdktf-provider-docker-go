package image


type ImageBuildUlimit struct {
	// soft limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#hard Image#hard}
	Hard *float64 `field:"required" json:"hard" yaml:"hard"`
	// type of ulimit, e.g. `nofile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#name Image#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// hard limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#soft Image#soft}
	Soft *float64 `field:"required" json:"soft" yaml:"soft"`
}

