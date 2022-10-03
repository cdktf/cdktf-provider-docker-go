package container


type ContainerMountsTmpfsOptions struct {
	// The permission mode for the tmpfs mount in an integer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#mode Container#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	// The size for the tmpfs mount in bytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#size_bytes Container#size_bytes}
	SizeBytes *float64 `field:"optional" json:"sizeBytes" yaml:"sizeBytes"`
}

