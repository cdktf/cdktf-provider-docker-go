//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerUploadList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerUploadList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerUploadList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerUploadList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerUploadList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerUploadList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerUploadListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
