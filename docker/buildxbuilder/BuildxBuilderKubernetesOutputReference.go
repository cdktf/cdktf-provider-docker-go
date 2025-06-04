// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v12/jsii"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v12/buildxbuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BuildxBuilderKubernetesOutputReference interface {
	cdktf.ComplexObject
	Annotations() *string
	SetAnnotations(val *string)
	AnnotationsInput() *string
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
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() *BuildxBuilderKubernetes
	SetInternalValue(val *BuildxBuilderKubernetes)
	Labels() *string
	SetLabels(val *string)
	LabelsInput() *string
	Limits() BuildxBuilderKubernetesLimitsOutputReference
	LimitsInput() *BuildxBuilderKubernetesLimits
	Loadbalance() *string
	SetLoadbalance(val *string)
	LoadbalanceInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Nodeselector() *string
	SetNodeselector(val *string)
	NodeselectorInput() *string
	Qemu() BuildxBuilderKubernetesQemuOutputReference
	QemuInput() *BuildxBuilderKubernetesQemu
	Replicas() *float64
	SetReplicas(val *float64)
	ReplicasInput() *float64
	Requests() BuildxBuilderKubernetesRequestsOutputReference
	RequestsInput() *BuildxBuilderKubernetesRequests
	Rootless() interface{}
	SetRootless(val interface{})
	RootlessInput() interface{}
	Schedulername() *string
	SetSchedulername(val *string)
	SchedulernameInput() *string
	Serviceaccount() *string
	SetServiceaccount(val *string)
	ServiceaccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Tolerations() *string
	SetTolerations(val *string)
	TolerationsInput() *string
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
	PutLimits(value *BuildxBuilderKubernetesLimits)
	PutQemu(value *BuildxBuilderKubernetesQemu)
	PutRequests(value *BuildxBuilderKubernetesRequests)
	ResetAnnotations()
	ResetDefaultLoad()
	ResetImage()
	ResetLabels()
	ResetLimits()
	ResetLoadbalance()
	ResetNamespace()
	ResetNodeselector()
	ResetQemu()
	ResetReplicas()
	ResetRequests()
	ResetRootless()
	ResetSchedulername()
	ResetServiceaccount()
	ResetTimeout()
	ResetTolerations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BuildxBuilderKubernetesOutputReference
type jsiiProxy_BuildxBuilderKubernetesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Annotations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) AnnotationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) DefaultLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) DefaultLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) InternalValue() *BuildxBuilderKubernetes {
	var returns *BuildxBuilderKubernetes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Labels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) LabelsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Limits() BuildxBuilderKubernetesLimitsOutputReference {
	var returns BuildxBuilderKubernetesLimitsOutputReference
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) LimitsInput() *BuildxBuilderKubernetesLimits {
	var returns *BuildxBuilderKubernetesLimits
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Loadbalance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) LoadbalanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadbalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Nodeselector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeselector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) NodeselectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeselectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Qemu() BuildxBuilderKubernetesQemuOutputReference {
	var returns BuildxBuilderKubernetesQemuOutputReference
	_jsii_.Get(
		j,
		"qemu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) QemuInput() *BuildxBuilderKubernetesQemu {
	var returns *BuildxBuilderKubernetesQemu
	_jsii_.Get(
		j,
		"qemuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Replicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) ReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Requests() BuildxBuilderKubernetesRequestsOutputReference {
	var returns BuildxBuilderKubernetesRequestsOutputReference
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) RequestsInput() *BuildxBuilderKubernetesRequests {
	var returns *BuildxBuilderKubernetesRequests
	_jsii_.Get(
		j,
		"requestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Rootless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) RootlessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootlessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Schedulername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) SchedulernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Serviceaccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceaccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) ServiceaccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceaccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) Tolerations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tolerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference) TolerationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tolerationsInput",
		&returns,
	)
	return returns
}


func NewBuildxBuilderKubernetesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BuildxBuilderKubernetesOutputReference {
	_init_.Initialize()

	if err := validateNewBuildxBuilderKubernetesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuildxBuilderKubernetesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderKubernetesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBuildxBuilderKubernetesOutputReference_Override(b BuildxBuilderKubernetesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilderKubernetesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetAnnotations(val *string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetDefaultLoad(val interface{}) {
	if err := j.validateSetDefaultLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLoad",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetInternalValue(val *BuildxBuilderKubernetes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetLabels(val *string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetLoadbalance(val *string) {
	if err := j.validateSetLoadbalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadbalance",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetNodeselector(val *string) {
	if err := j.validateSetNodeselectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeselector",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetReplicas(val *float64) {
	if err := j.validateSetReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetRootless(val interface{}) {
	if err := j.validateSetRootlessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootless",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetSchedulername(val *string) {
	if err := j.validateSetSchedulernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulername",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetServiceaccount(val *string) {
	if err := j.validateSetServiceaccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceaccount",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilderKubernetesOutputReference)SetTolerations(val *string) {
	if err := j.validateSetTolerationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tolerations",
		val,
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) PutLimits(value *BuildxBuilderKubernetesLimits) {
	if err := b.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLimits",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) PutQemu(value *BuildxBuilderKubernetesQemu) {
	if err := b.validatePutQemuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putQemu",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) PutRequests(value *BuildxBuilderKubernetesRequests) {
	if err := b.validatePutRequestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequests",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetAnnotations() {
	_jsii_.InvokeVoid(
		b,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetDefaultLoad() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultLoad",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetLimits() {
	_jsii_.InvokeVoid(
		b,
		"resetLimits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetLoadbalance() {
	_jsii_.InvokeVoid(
		b,
		"resetLoadbalance",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		b,
		"resetNamespace",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetNodeselector() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeselector",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetQemu() {
	_jsii_.InvokeVoid(
		b,
		"resetQemu",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetReplicas() {
	_jsii_.InvokeVoid(
		b,
		"resetReplicas",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetRequests() {
	_jsii_.InvokeVoid(
		b,
		"resetRequests",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetRootless() {
	_jsii_.InvokeVoid(
		b,
		"resetRootless",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetSchedulername() {
	_jsii_.InvokeVoid(
		b,
		"resetSchedulername",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetServiceaccount() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceaccount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeout",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ResetTolerations() {
	_jsii_.InvokeVoid(
		b,
		"resetTolerations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BuildxBuilderKubernetesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

