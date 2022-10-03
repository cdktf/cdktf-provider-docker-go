package service


type ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext struct {
	// Disable SELinux.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#disable Service#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// SELinux level label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#level Service#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// SELinux role label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#role Service#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// SELinux type label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#type Service#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// SELinux user label.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#user Service#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

