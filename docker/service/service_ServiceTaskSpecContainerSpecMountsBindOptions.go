package service


type ServiceTaskSpecContainerSpecMountsBindOptions struct {
	// Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount.
	//
	// See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#propagation Service#propagation}
	Propagation *string `field:"optional" json:"propagation" yaml:"propagation"`
}

