//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package registryimage

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryImageBuildAuthConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildAuthConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegistryImageBuildAuthConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

