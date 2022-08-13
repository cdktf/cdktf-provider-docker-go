// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ServiceTaskSpecContainerSpecPrivileges struct {
	// credential_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#credential_spec Service#credential_spec}
	CredentialSpec *ServiceTaskSpecContainerSpecPrivilegesCredentialSpec `field:"optional" json:"credentialSpec" yaml:"credentialSpec"`
	// se_linux_context block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/service#se_linux_context Service#se_linux_context}
	SeLinuxContext *ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext `field:"optional" json:"seLinuxContext" yaml:"seLinuxContext"`
}

