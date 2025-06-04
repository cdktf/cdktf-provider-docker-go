// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v12/jsii"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v12/buildxbuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BuildxBuilderRemoteOutputReference interface {
	cdktf.ComplexObject
	Cacert() *string
	SetCacert(val *string)
	CacertInput() *string
	Cert() *string
	SetCert(val *string)
	CertInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultLoad() interface{}
	SetDefaultLoad(val interface{})
	DefaultLoadInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BuildxBuilderRemote
	SetInternalValue(val *BuildxBuilderRemote)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Servername() *string
	SetServername(val *string)
	ServernameInput() *string
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
	ResetCacert()
	ResetCert()
	ResetDefaultLoad()
	ResetKey()
	ResetServername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BuildxBuilderRemoteOutputReference
type jsiiProxy_BuildxBuilderRemoteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) Cacert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) CacertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) DefaultLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) DefaultLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) InternalValue() *BuildxBuilderRemote {
	var returns *BuildxBuilderRemote
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) Servername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) ServernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBuildxBuilderRemoteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BuildxBuilderRemoteOutputReference {
	_init_.Initialize()

	if err := validateNewBuildxBuilderRemoteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuildxBuilderRemoteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderRemoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBuildxBuilderRemoteOutputReference_Override(b BuildxBuilderRemoteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderRemoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetCacert(val *string) {
	if err := j.validateSetCacertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacert",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetCert(val *string) {
	if err := j.validateSetCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cert",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetDefaultLoad(val interface{}) {
	if err := j.validateSetDefaultLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLoad",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetInternalValue(val *BuildxBuilderRemote) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetServername(val *string) {
	if err := j.validateSetServernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servername",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderRemoteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ResetCacert() {
	_jsii_.InvokeVoid(
		b,
		"resetCacert",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ResetCert() {
	_jsii_.InvokeVoid(
		b,
		"resetCert",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ResetDefaultLoad() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultLoad",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		b,
		"resetKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ResetServername() {
	_jsii_.InvokeVoid(
		b,
		"resetServername",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BuildxBuilderRemoteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

