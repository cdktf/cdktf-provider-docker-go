package service


type ServiceTaskSpecContainerSpecPrivilegesCredentialSpec struct {
	// Load credential spec from this file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#file Service#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// Load credential spec from this value in the Windows registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#registry Service#registry}
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
}

