// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceMode struct {
	// The global service mode. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#global Service#global}
	Global interface{} `field:"optional" json:"global" yaml:"global"`
	// replicated block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#replicated Service#replicated}
	Replicated *ServiceModeReplicated `field:"optional" json:"replicated" yaml:"replicated"`
}

