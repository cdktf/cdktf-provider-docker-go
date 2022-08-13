// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type DockerProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#alias DockerProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// PEM-encoded content of Docker host CA certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#ca_material DockerProvider#ca_material}
	CaMaterial *string `field:"optional" json:"caMaterial" yaml:"caMaterial"`
	// PEM-encoded content of Docker client certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#cert_material DockerProvider#cert_material}
	CertMaterial *string `field:"optional" json:"certMaterial" yaml:"certMaterial"`
	// Path to directory with Docker TLS config.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#cert_path DockerProvider#cert_path}
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// The Docker daemon address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#host DockerProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// PEM-encoded content of Docker client private key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#key_material DockerProvider#key_material}
	KeyMaterial *string `field:"optional" json:"keyMaterial" yaml:"keyMaterial"`
	// registry_auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#registry_auth DockerProvider#registry_auth}
	RegistryAuth interface{} `field:"optional" json:"registryAuth" yaml:"registryAuth"`
	// Additional SSH option flags to be appended when using `ssh://` protocol.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker#ssh_opts DockerProvider#ssh_opts}
	SshOpts *[]*string `field:"optional" json:"sshOpts" yaml:"sshOpts"`
}

