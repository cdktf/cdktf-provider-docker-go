// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceRollbackConfig struct {
	// Delay between task rollbacks (ns|us|ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#delay Service#delay}
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
	// Action on rollback failure: pause | continue. Defaults to `pause`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#failure_action Service#failure_action}
	FailureAction *string `field:"optional" json:"failureAction" yaml:"failureAction"`
	// Failure rate to tolerate during a rollback. Defaults to `0.0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#max_failure_ratio Service#max_failure_ratio}
	MaxFailureRatio *string `field:"optional" json:"maxFailureRatio" yaml:"maxFailureRatio"`
	// Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#monitor Service#monitor}
	Monitor *string `field:"optional" json:"monitor" yaml:"monitor"`
	// Rollback order: either 'stop-first' or 'start-first'. Defaults to `stop-first`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#order Service#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Maximum number of tasks to be rollbacked in one iteration. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#parallelism Service#parallelism}
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
}

