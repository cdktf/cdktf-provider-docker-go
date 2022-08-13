// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDockerNetworkConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Docker network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/d/network#name DataDockerNetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

