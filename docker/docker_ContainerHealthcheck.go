// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ContainerHealthcheck struct {
	// Command to run to check health.
	//
	// For example, to run `curl -f localhost/health` set the command to be `["CMD", "curl", "-f", "localhost/health"]`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#test Container#test}
	Test *[]*string `field:"required" json:"test" yaml:"test"`
	// Time between running the check (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#interval Container#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Consecutive failures needed to report unhealthy. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#retries Container#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#start_period Container#start_period}
	StartPeriod *string `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#timeout Container#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

