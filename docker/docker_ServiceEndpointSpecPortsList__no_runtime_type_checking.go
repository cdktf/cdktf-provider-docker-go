//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceEndpointSpecPortsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceEndpointSpecPortsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceEndpointSpecPortsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceEndpointSpecPortsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceEndpointSpecPortsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceEndpointSpecPortsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceEndpointSpecPortsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

