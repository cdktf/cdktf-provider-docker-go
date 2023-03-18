package service


type ServiceTaskSpecContainerSpecMounts struct {
	// Container path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#target Service#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The mount type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#type Service#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// bind_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#bind_options Service#bind_options}
	BindOptions *ServiceTaskSpecContainerSpecMountsBindOptions `field:"optional" json:"bindOptions" yaml:"bindOptions"`
	// Whether the mount should be read-only.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#read_only Service#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Mount source (e.g. a volume name, a host path).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#source Service#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// tmpfs_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#tmpfs_options Service#tmpfs_options}
	TmpfsOptions *ServiceTaskSpecContainerSpecMountsTmpfsOptions `field:"optional" json:"tmpfsOptions" yaml:"tmpfsOptions"`
	// volume_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#volume_options Service#volume_options}
	VolumeOptions *ServiceTaskSpecContainerSpecMountsVolumeOptions `field:"optional" json:"volumeOptions" yaml:"volumeOptions"`
}

