// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecContainerSpecMountsTmpfsOptions struct {
	// The permission mode for the tmpfs mount in an integer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#mode Service#mode}
	Mode *float64 `field:"optional" json:"mode" yaml:"mode"`
	// The size for the tmpfs mount in bytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#size_bytes Service#size_bytes}
	SizeBytes *float64 `field:"optional" json:"sizeBytes" yaml:"sizeBytes"`
}

