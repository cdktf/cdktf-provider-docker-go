// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecPlacement struct {
	// An array of constraints. e.g.: `node.role==manager`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#constraints Service#constraints}
	Constraints *[]*string `field:"optional" json:"constraints" yaml:"constraints"`
	// Maximum number of replicas for per node (default value is `0`, which is unlimited).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#max_replicas Service#max_replicas}
	MaxReplicas *float64 `field:"optional" json:"maxReplicas" yaml:"maxReplicas"`
	// platforms block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#platforms Service#platforms}
	Platforms interface{} `field:"optional" json:"platforms" yaml:"platforms"`
	// Preferences provide a way to make the scheduler aware of factors such as topology.
	//
	// They are provided in order from highest to lowest precedence, e.g.: `spread=node.role.manager`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#prefs Service#prefs}
	Prefs *[]*string `field:"optional" json:"prefs" yaml:"prefs"`
}
