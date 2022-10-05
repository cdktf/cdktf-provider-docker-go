package image

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v3/jsii"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/image/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImageBuildOutputReference interface {
	cdktf.ComplexObject
	BuildArg() *map[string]*string
	SetBuildArg(val *map[string]*string)
	BuildArgInput() *map[string]*string
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
	Dockerfile() *string
	SetDockerfile(val *string)
	DockerfileInput() *string
	ForceRemove() interface{}
	SetForceRemove(val interface{})
	ForceRemoveInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ImageBuild
	SetInternalValue(val *ImageBuild)
	Label() *map[string]*string
	SetLabel(val *map[string]*string)
	LabelInput() *map[string]*string
	NoCache() interface{}
	SetNoCache(val interface{})
	NoCacheInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Remove() interface{}
	SetRemove(val interface{})
	RemoveInput() interface{}
	Tag() *[]*string
	SetTag(val *[]*string)
	TagInput() *[]*string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetBuildArg()
	ResetDockerfile()
	ResetForceRemove()
	ResetLabel()
	ResetNoCache()
	ResetRemove()
	ResetTag()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ImageBuildOutputReference
type jsiiProxy_ImageBuildOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ImageBuildOutputReference) BuildArg() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildArg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) BuildArgInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildArgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Dockerfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) DockerfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) ForceRemove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) ForceRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) InternalValue() *ImageBuild {
	var returns *ImageBuild
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Label() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) LabelInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) NoCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) NoCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Remove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) RemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Tag() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) TagInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageBuildOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewImageBuildOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ImageBuildOutputReference {
	_init_.Initialize()

	if err := validateNewImageBuildOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImageBuildOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-docker.image.ImageBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewImageBuildOutputReference_Override(i ImageBuildOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.image.ImageBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetBuildArg(val *map[string]*string) {
	if err := j.validateSetBuildArgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildArg",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetDockerfile(val *string) {
	if err := j.validateSetDockerfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfile",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetForceRemove(val interface{}) {
	if err := j.validateSetForceRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRemove",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetInternalValue(val *ImageBuild) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetLabel(val *map[string]*string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetNoCache(val interface{}) {
	if err := j.validateSetNoCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCache",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetRemove(val interface{}) {
	if err := j.validateSetRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remove",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetTag(val *[]*string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ImageBuildOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetBuildArg() {
	_jsii_.InvokeVoid(
		i,
		"resetBuildArg",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetDockerfile() {
	_jsii_.InvokeVoid(
		i,
		"resetDockerfile",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetForceRemove() {
	_jsii_.InvokeVoid(
		i,
		"resetForceRemove",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetLabel() {
	_jsii_.InvokeVoid(
		i,
		"resetLabel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetNoCache() {
	_jsii_.InvokeVoid(
		i,
		"resetNoCache",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetRemove() {
	_jsii_.InvokeVoid(
		i,
		"resetRemove",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		i,
		"resetTag",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		i,
		"resetTarget",
		nil, // no parameters
	)
}

func (i *jsiiProxy_ImageBuildOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageBuildOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

