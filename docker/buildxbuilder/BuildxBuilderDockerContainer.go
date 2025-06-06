// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder


type BuildxBuilderDockerContainer struct {
	// Sets the cgroup parent of the container if Docker is using the "cgroupfs" driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cgroup_parent BuildxBuilder#cgroup_parent}
	CgroupParent *string `field:"optional" json:"cgroupParent" yaml:"cgroupParent"`
	// Sets the CPU CFS scheduler period for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cpu_period BuildxBuilder#cpu_period}
	CpuPeriod *string `field:"optional" json:"cpuPeriod" yaml:"cpuPeriod"`
	// Imposes a CPU CFS quota on the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cpu_quota BuildxBuilder#cpu_quota}
	CpuQuota *string `field:"optional" json:"cpuQuota" yaml:"cpuQuota"`
	// Limits the set of CPU cores the container can use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cpuset_cpus BuildxBuilder#cpuset_cpus}
	CpusetCpus *string `field:"optional" json:"cpusetCpus" yaml:"cpusetCpus"`
	// Limits the set of CPU memory nodes the container can use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cpuset_mems BuildxBuilder#cpuset_mems}
	CpusetMems *string `field:"optional" json:"cpusetMems" yaml:"cpusetMems"`
	// Configures CPU shares (relative weight) of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#cpu_shares BuildxBuilder#cpu_shares}
	CpuShares *string `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// Automatically load images to the Docker Engine image store. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#default_load BuildxBuilder#default_load}
	DefaultLoad interface{} `field:"optional" json:"defaultLoad" yaml:"defaultLoad"`
	// Sets environment variables in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#env BuildxBuilder#env}
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Sets the BuildKit image to use for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#image BuildxBuilder#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Sets the amount of memory the container can use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#memory BuildxBuilder#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Sets the memory swap limit for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#memory_swap BuildxBuilder#memory_swap}
	MemorySwap *string `field:"optional" json:"memorySwap" yaml:"memorySwap"`
	// Sets the network mode for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#network BuildxBuilder#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Sets the container's restart policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder#restart_policy BuildxBuilder#restart_policy}
	RestartPolicy *string `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
}

