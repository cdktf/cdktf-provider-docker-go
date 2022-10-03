package image


type ImageBuild struct {
	// Context path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#path Image#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Set build-time variables.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#build_arg Image#build_arg}
	BuildArg *map[string]*string `field:"optional" json:"buildArg" yaml:"buildArg"`
	// Name of the Dockerfile. Defaults to `Dockerfile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#dockerfile Image#dockerfile}
	Dockerfile *string `field:"optional" json:"dockerfile" yaml:"dockerfile"`
	// Always remove intermediate containers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#force_remove Image#force_remove}
	ForceRemove interface{} `field:"optional" json:"forceRemove" yaml:"forceRemove"`
	// Set metadata for an image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#label Image#label}
	Label *map[string]*string `field:"optional" json:"label" yaml:"label"`
	// Do not use cache when building the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#no_cache Image#no_cache}
	NoCache interface{} `field:"optional" json:"noCache" yaml:"noCache"`
	// Remove intermediate containers after a successful build. Defaults to  `true`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#remove Image#remove}
	Remove interface{} `field:"optional" json:"remove" yaml:"remove"`
	// Name and optionally a tag in the 'name:tag' format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#tag Image#tag}
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
	// Set the target build stage to build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#target Image#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

