package container


type ContainerMountsVolumeOptions struct {
	// Name of the driver to use to create the volume.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#driver_name Container#driver_name}
	DriverName *string `field:"optional" json:"driverName" yaml:"driverName"`
	// key/value map of driver specific options.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#driver_options Container#driver_options}
	DriverOptions *map[string]*string `field:"optional" json:"driverOptions" yaml:"driverOptions"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#labels Container#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Populate volume with data from the target.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#no_copy Container#no_copy}
	NoCopy interface{} `field:"optional" json:"noCopy" yaml:"noCopy"`
}
