// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DockerProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#alias DockerProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// PEM-encoded content of Docker host CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#ca_material DockerProvider#ca_material}
	CaMaterial *string `field:"optional" json:"caMaterial" yaml:"caMaterial"`
	// PEM-encoded content of Docker client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#cert_material DockerProvider#cert_material}
	CertMaterial *string `field:"optional" json:"certMaterial" yaml:"certMaterial"`
	// Path to directory with Docker TLS config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#cert_path DockerProvider#cert_path}
	CertPath *string `field:"optional" json:"certPath" yaml:"certPath"`
	// The name of the Docker context to use.
	//
	// Can also be set via `DOCKER_CONTEXT` environment variable. Overrides the `host` if set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#context DockerProvider#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// If set to `true`, the provider will not check if the Docker daemon is running.
	//
	// This is useful for resources/data_sourcess that do not require a running Docker daemon, such as the data source `docker_registry_image`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#disable_docker_daemon_check DockerProvider#disable_docker_daemon_check}
	DisableDockerDaemonCheck interface{} `field:"optional" json:"disableDockerDaemonCheck" yaml:"disableDockerDaemonCheck"`
	// The Docker daemon address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#host DockerProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// PEM-encoded content of Docker client private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#key_material DockerProvider#key_material}
	KeyMaterial *string `field:"optional" json:"keyMaterial" yaml:"keyMaterial"`
	// registry_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#registry_auth DockerProvider#registry_auth}
	RegistryAuth interface{} `field:"optional" json:"registryAuth" yaml:"registryAuth"`
	// Additional SSH option flags to be appended when using `ssh://` protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.2/docs#ssh_opts DockerProvider#ssh_opts}
	SshOpts *[]*string `field:"optional" json:"sshOpts" yaml:"sshOpts"`
}

