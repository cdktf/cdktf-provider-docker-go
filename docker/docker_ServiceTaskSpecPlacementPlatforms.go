// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecPlacementPlatforms struct {
	// The architecture, e.g. `amd64`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#architecture Service#architecture}
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
	// The operation system, e.g. `linux`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#os Service#os}
	Os *string `field:"required" json:"os" yaml:"os"`
}

