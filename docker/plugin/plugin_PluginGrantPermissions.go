package plugin


type PluginGrantPermissions struct {
	// The name of the permission.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/plugin#name Plugin#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the permission.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/plugin#value Plugin#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

