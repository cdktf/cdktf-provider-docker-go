// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecRestartPolicy struct {
	// Condition for restart.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#condition Service#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Delay between restart attempts (ms|s|m|h).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#delay Service#delay}
	Delay *string `field:"optional" json:"delay" yaml:"delay"`
	// Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#max_attempts Service#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#window Service#window}
	Window *string `field:"optional" json:"window" yaml:"window"`
}

