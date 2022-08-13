// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type RegistryImageBuild struct {
	// The absolute path to the context folder. You can use the helper function '${path.cwd}/context-dir'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#context RegistryImage#context}
	Context *string `field:"required" json:"context" yaml:"context"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#auth_config RegistryImage#auth_config}
	AuthConfig interface{} `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Pairs for build-time variables in the form TODO.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#build_args RegistryImage#build_args}
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// BuildID is an optional identifier that can be passed together with the build request.
	//
	// The same identifier can be used to gracefully cancel the build with the cancel request.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#build_id RegistryImage#build_id}
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// Images to consider as cache sources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cache_from RegistryImage#cache_from}
	CacheFrom *[]*string `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Optional parent cgroup for the container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cgroup_parent RegistryImage#cgroup_parent}
	CgroupParent *string `field:"optional" json:"cgroupParent" yaml:"cgroupParent"`
	// The length of a CPU period in microseconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cpu_period RegistryImage#cpu_period}
	CpuPeriod *float64 `field:"optional" json:"cpuPeriod" yaml:"cpuPeriod"`
	// Microseconds of CPU time that the container can get in a CPU period.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cpu_quota RegistryImage#cpu_quota}
	CpuQuota *float64 `field:"optional" json:"cpuQuota" yaml:"cpuQuota"`
	// CPUs in which to allow execution (e.g., `0-3`, `0`, `1`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cpu_set_cpus RegistryImage#cpu_set_cpus}
	CpuSetCpus *string `field:"optional" json:"cpuSetCpus" yaml:"cpuSetCpus"`
	// MEMs in which to allow execution (`0-3`, `0`, `1`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cpu_set_mems RegistryImage#cpu_set_mems}
	CpuSetMems *string `field:"optional" json:"cpuSetMems" yaml:"cpuSetMems"`
	// CPU shares (relative weight).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#cpu_shares RegistryImage#cpu_shares}
	CpuShares *float64 `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// Dockerfile file. Defaults to `Dockerfile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#dockerfile RegistryImage#dockerfile}
	Dockerfile *string `field:"optional" json:"dockerfile" yaml:"dockerfile"`
	// A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form ["hostname:IP"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#extra_hosts RegistryImage#extra_hosts}
	ExtraHosts *[]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// Always remove intermediate containers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#force_remove RegistryImage#force_remove}
	ForceRemove interface{} `field:"optional" json:"forceRemove" yaml:"forceRemove"`
	// Isolation represents the isolation technology of a container. The supported values are.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#isolation RegistryImage#isolation}
	Isolation *string `field:"optional" json:"isolation" yaml:"isolation"`
	// User-defined key/value metadata.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#labels RegistryImage#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Set memory limit for build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#memory RegistryImage#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Total memory (memory + swap), -1 to enable unlimited swap.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#memory_swap RegistryImage#memory_swap}
	MemorySwap *float64 `field:"optional" json:"memorySwap" yaml:"memorySwap"`
	// Set the networking mode for the RUN instructions during build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#network_mode RegistryImage#network_mode}
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Do not use the cache when building the image.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#no_cache RegistryImage#no_cache}
	NoCache interface{} `field:"optional" json:"noCache" yaml:"noCache"`
	// Set platform if server is multi-platform capable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#platform RegistryImage#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Attempt to pull the image even if an older image exists locally.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#pull_parent RegistryImage#pull_parent}
	PullParent interface{} `field:"optional" json:"pullParent" yaml:"pullParent"`
	// A Git repository URI or HTTP/HTTPS context URI.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#remote_context RegistryImage#remote_context}
	RemoteContext *string `field:"optional" json:"remoteContext" yaml:"remoteContext"`
	// Remove intermediate containers after a successful build (default behavior).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#remove RegistryImage#remove}
	Remove interface{} `field:"optional" json:"remove" yaml:"remove"`
	// The security options.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#security_opt RegistryImage#security_opt}
	SecurityOpt *[]*string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// Set an ID for the build session.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#session_id RegistryImage#session_id}
	SessionId *string `field:"optional" json:"sessionId" yaml:"sessionId"`
	// Size of /dev/shm in bytes. The size must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#shm_size RegistryImage#shm_size}
	ShmSize *float64 `field:"optional" json:"shmSize" yaml:"shmSize"`
	// If true the new layers are squashed into a new image with a single new layer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#squash RegistryImage#squash}
	Squash interface{} `field:"optional" json:"squash" yaml:"squash"`
	// Suppress the build output and print image ID on success.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#suppress_output RegistryImage#suppress_output}
	SuppressOutput interface{} `field:"optional" json:"suppressOutput" yaml:"suppressOutput"`
	// Set the target build stage to build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#target RegistryImage#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// ulimit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#ulimit RegistryImage#ulimit}
	Ulimit interface{} `field:"optional" json:"ulimit" yaml:"ulimit"`
	// Version of the underlying builder to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#version RegistryImage#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

