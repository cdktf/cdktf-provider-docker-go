package service


type ServiceTaskSpec struct {
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#container_spec Service#container_spec}
	ContainerSpec *ServiceTaskSpecContainerSpec `field:"required" json:"containerSpec" yaml:"containerSpec"`
	// A counter that triggers an update even if no relevant parameters have been changed. See the [spec](https://github.com/docker/swarmkit/blob/master/api/specs.proto#L126).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#force_update Service#force_update}
	ForceUpdate *float64 `field:"optional" json:"forceUpdate" yaml:"forceUpdate"`
	// log_driver block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#log_driver Service#log_driver}
	LogDriver *ServiceTaskSpecLogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// Ids of the networks in which the  container will be put in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#networks Service#networks}
	Networks *[]*string `field:"optional" json:"networks" yaml:"networks"`
	// networks_advanced block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#networks_advanced Service#networks_advanced}
	NetworksAdvanced interface{} `field:"optional" json:"networksAdvanced" yaml:"networksAdvanced"`
	// placement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#placement Service#placement}
	Placement *ServiceTaskSpecPlacement `field:"optional" json:"placement" yaml:"placement"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#resources Service#resources}
	Resources *ServiceTaskSpecResources `field:"optional" json:"resources" yaml:"resources"`
	// restart_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#restart_policy Service#restart_policy}
	RestartPolicy *ServiceTaskSpecRestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// Runtime is the type of runtime specified for the task executor. See the [types](https://github.com/moby/moby/blob/master/api/types/swarm/runtime.go).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#runtime Service#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
}

