package service


type ServiceTaskSpecContainerSpecHosts struct {
	// The name of the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#host Service#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// The ip of the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#ip Service#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

