// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package buildxbuilder

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BuildxBuilder) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateImportFromParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateMoveToIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validatePutDockerContainerParameters(value *BuildxBuilderDockerContainer) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validatePutKubernetesParameters(value *BuildxBuilderKubernetes) error {
	return nil
}

func (b *jsiiProxy_BuildxBuilder) validatePutRemoteParameters(value *BuildxBuilderRemote) error {
	return nil
}

func validateBuildxBuilder_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateBuildxBuilder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBuildxBuilder_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBuildxBuilder_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetAppendParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetBootstrapParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetBuildkitConfigParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetBuildkitFlagsParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetDriverParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetDriverOptionsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetEndpointParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetNodeAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetPlatformParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_BuildxBuilder) validateSetUseParameters(val interface{}) error {
	return nil
}

func validateNewBuildxBuilderParameters(scope constructs.Construct, id *string, config *BuildxBuilderConfig) error {
	return nil
}

