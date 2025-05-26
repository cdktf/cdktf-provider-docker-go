// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v11/jsii"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v11/buildxbuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BuildxBuilderDockerContainerOutputReference interface {
	cdktf.ComplexObject
	CgroupParent() *string
	SetCgroupParent(val *string)
	CgroupParentInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CpuPeriod() *string
	SetCpuPeriod(val *string)
	CpuPeriodInput() *string
	CpuQuota() *string
	SetCpuQuota(val *string)
	CpuQuotaInput() *string
	CpusetCpus() *string
	SetCpusetCpus(val *string)
	CpusetCpusInput() *string
	CpusetMems() *string
	SetCpusetMems(val *string)
	CpusetMemsInput() *string
	CpuShares() *string
	SetCpuShares(val *string)
	CpuSharesInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultLoad() interface{}
	SetDefaultLoad(val interface{})
	DefaultLoadInput() interface{}
	Env() *map[string]*string
	SetEnv(val *map[string]*string)
	EnvInput() *map[string]*string
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() *BuildxBuilderDockerContainer
	SetInternalValue(val *BuildxBuilderDockerContainer)
	Memory() *string
	SetMemory(val *string)
	MemoryInput() *string
	MemorySwap() *string
	SetMemorySwap(val *string)
	MemorySwapInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	RestartPolicy() *string
	SetRestartPolicy(val *string)
	RestartPolicyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCgroupParent()
	ResetCpuPeriod()
	ResetCpuQuota()
	ResetCpusetCpus()
	ResetCpusetMems()
	ResetCpuShares()
	ResetDefaultLoad()
	ResetEnv()
	ResetImage()
	ResetMemory()
	ResetMemorySwap()
	ResetNetwork()
	ResetRestartPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BuildxBuilderDockerContainerOutputReference
type jsiiProxy_BuildxBuilderDockerContainerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CgroupParent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CgroupParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuQuota() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuQuotaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpusetCpus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpusetCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpusetCpusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpusetCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpusetMems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpusetMems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpusetMemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpusetMemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuShares() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CpuSharesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) DefaultLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) DefaultLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) EnvInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) InternalValue() *BuildxBuilderDockerContainer {
	var returns *BuildxBuilderDockerContainer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) MemorySwap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memorySwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) MemorySwapInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memorySwapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) RestartPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) RestartPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restartPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBuildxBuilderDockerContainerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BuildxBuilderDockerContainerOutputReference {
	_init_.Initialize()

	if err := validateNewBuildxBuilderDockerContainerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuildxBuilderDockerContainerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderDockerContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBuildxBuilderDockerContainerOutputReference_Override(b BuildxBuilderDockerContainerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderDockerContainerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCgroupParent(val *string) {
	if err := j.validateSetCgroupParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupParent",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCpuPeriod(val *string) {
	if err := j.validateSetCpuPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPeriod",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCpuQuota(val *string) {
	if err := j.validateSetCpuQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuQuota",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCpusetCpus(val *string) {
	if err := j.validateSetCpusetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpusetCpus",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCpusetMems(val *string) {
	if err := j.validateSetCpusetMemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpusetMems",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetCpuShares(val *string) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetDefaultLoad(val interface{}) {
	if err := j.validateSetDefaultLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLoad",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetEnv(val *map[string]*string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetInternalValue(val *BuildxBuilderDockerContainer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetMemorySwap(val *string) {
	if err := j.validateSetMemorySwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySwap",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetRestartPolicy(val *string) {
	if err := j.validateSetRestartPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartPolicy",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderDockerContainerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCgroupParent() {
	_jsii_.InvokeVoid(
		b,
		"resetCgroupParent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCpuPeriod() {
	_jsii_.InvokeVoid(
		b,
		"resetCpuPeriod",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCpuQuota() {
	_jsii_.InvokeVoid(
		b,
		"resetCpuQuota",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCpusetCpus() {
	_jsii_.InvokeVoid(
		b,
		"resetCpusetCpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCpusetMems() {
	_jsii_.InvokeVoid(
		b,
		"resetCpusetMems",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetCpuShares() {
	_jsii_.InvokeVoid(
		b,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetDefaultLoad() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultLoad",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		b,
		"resetEnv",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		b,
		"resetMemory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetMemorySwap() {
	_jsii_.InvokeVoid(
		b,
		"resetMemorySwap",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		b,
		"resetNetwork",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ResetRestartPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetRestartPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderDockerContainerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

