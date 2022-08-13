// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceEndpointSpec struct {
	// The mode of resolution to use for internal load balancing between tasks.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#mode Service#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#ports Service#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

