// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceModeReplicated struct {
	// The amount of replicas of the service. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#replicas Service#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

