// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageConfig struct {
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
	// The name of the Docker image, including any tags or SHA256 repo digests.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#name Image#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// build block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#build Image#build}
	BuildAttribute *ImageBuild `field:"optional" json:"buildAttribute" yaml:"buildAttribute"`
	// If true, then the image is removed forcibly when the resource is destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#force_remove Image#force_remove}
	ForceRemove interface{} `field:"optional" json:"forceRemove" yaml:"forceRemove"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#id Image#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, then the Docker image won't be deleted on destroy operation.
	//
	// If this is false, it will delete the image from the docker local storage on destroy operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#keep_locally Image#keep_locally}
	KeepLocally interface{} `field:"optional" json:"keepLocally" yaml:"keepLocally"`
	// A value which cause an image pull when changed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#pull_trigger Image#pull_trigger}
	PullTrigger *string `field:"optional" json:"pullTrigger" yaml:"pullTrigger"`
	// List of values which cause an image pull when changed.
	//
	// This is used to store the image digest from the registry when using the [docker_registry_image](../data-sources/registry_image.md).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#pull_triggers Image#pull_triggers}
	PullTriggers *[]*string `field:"optional" json:"pullTriggers" yaml:"pullTriggers"`
	// A map of arbitrary strings that, when changed, will force the `docker_image` resource to be replaced.
	//
	// This can be used to rebuild an image when contents of source code folders change
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#triggers Image#triggers}
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}
