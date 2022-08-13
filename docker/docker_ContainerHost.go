// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ContainerHost struct {
	// Hostname to add.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#host Container#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// IP address this hostname should resolve to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#ip Container#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
}

