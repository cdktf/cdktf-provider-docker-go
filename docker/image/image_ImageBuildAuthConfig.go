package image


type ImageBuildAuthConfig struct {
	// hostname of the registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#host_name Image#host_name}
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// the auth token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#auth Image#auth}
	Auth *string `field:"optional" json:"auth" yaml:"auth"`
	// the user emal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#email Image#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// the identity token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#identity_token Image#identity_token}
	IdentityToken *string `field:"optional" json:"identityToken" yaml:"identityToken"`
	// the registry password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#password Image#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// the registry token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#registry_token Image#registry_token}
	RegistryToken *string `field:"optional" json:"registryToken" yaml:"registryToken"`
	// the server address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#server_address Image#server_address}
	ServerAddress *string `field:"optional" json:"serverAddress" yaml:"serverAddress"`
	// the registry user name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/image#user_name Image#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}
