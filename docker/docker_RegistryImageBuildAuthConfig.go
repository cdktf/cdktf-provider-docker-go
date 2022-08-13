// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type RegistryImageBuildAuthConfig struct {
	// hostname of the registry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#host_name RegistryImage#host_name}
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// the auth token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#auth RegistryImage#auth}
	Auth *string `field:"optional" json:"auth" yaml:"auth"`
	// the user emal.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#email RegistryImage#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// the identity token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#identity_token RegistryImage#identity_token}
	IdentityToken *string `field:"optional" json:"identityToken" yaml:"identityToken"`
	// the registry password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#password RegistryImage#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// the registry token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#registry_token RegistryImage#registry_token}
	RegistryToken *string `field:"optional" json:"registryToken" yaml:"registryToken"`
	// the server address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#server_address RegistryImage#server_address}
	ServerAddress *string `field:"optional" json:"serverAddress" yaml:"serverAddress"`
	// the registry user name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/registry_image#user_name RegistryImage#user_name}
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

