// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecResourcesLimits struct {
	// The amounf of memory in bytes the container allocates.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#memory_bytes Service#memory_bytes}
	MemoryBytes *float64 `field:"optional" json:"memoryBytes" yaml:"memoryBytes"`
	// CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#nano_cpus Service#nano_cpus}
	NanoCpus *float64 `field:"optional" json:"nanoCpus" yaml:"nanoCpus"`
}
