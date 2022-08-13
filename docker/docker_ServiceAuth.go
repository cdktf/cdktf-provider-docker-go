// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceAuth struct {
	// The address of the server for the authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#server_address Service#server_address}
	ServerAddress *string `field:"required" json:"serverAddress" yaml:"serverAddress"`
	// The password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#password Service#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#username Service#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

