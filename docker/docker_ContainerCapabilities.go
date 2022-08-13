// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ContainerCapabilities struct {
	// List of linux capabilities to add.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#add Container#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// List of linux capabilities to drop.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#drop Container#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

