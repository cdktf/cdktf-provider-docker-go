// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecContainerSpecConfigs struct {
	// ID of the specific config that we're referencing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#config_id Service#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// Represents the final filename in the filesystem.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#file_name Service#file_name}
	FileName *string `field:"required" json:"fileName" yaml:"fileName"`
	// Name of the config that this references, but this is just provided for lookup/display purposes.
	//
	// The config in the reference will be identified by its ID
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#config_name Service#config_name}
	ConfigName *string `field:"optional" json:"configName" yaml:"configName"`
	// Represents the file GID. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#file_gid Service#file_gid}
	FileGid *string `field:"optional" json:"fileGid" yaml:"fileGid"`
	// Represents represents the FileMode of the file. Defaults to `0o444`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#file_mode Service#file_mode}
	FileMode *float64 `field:"optional" json:"fileMode" yaml:"fileMode"`
	// Represents the file UID. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#file_uid Service#file_uid}
	FileUid *string `field:"optional" json:"fileUid" yaml:"fileUid"`
}

