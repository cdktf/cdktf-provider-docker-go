package image

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v3/image/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/docker/r/image docker_image}.
type Image interface {
	cdktf.TerraformResource
	BuildAttribute() ImageBuildOutputReference
	BuildAttributeInput() *ImageBuild
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ForceRemove() interface{}
	SetForceRemove(val interface{})
	ForceRemoveInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	ImageId() *string
	KeepLocally() interface{}
	SetKeepLocally(val interface{})
	KeepLocallyInput() interface{}
	Latest() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Output() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PullTrigger() *string
	SetPullTrigger(val *string)
	PullTriggerInput() *string
	PullTriggers() *[]*string
	SetPullTriggers(val *[]*string)
	PullTriggersInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RepoDigest() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Triggers() *map[string]*string
	SetTriggers(val *map[string]*string)
	TriggersInput() *map[string]*string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBuildAttribute(value *ImageBuild)
	ResetBuildAttribute()
	ResetForceRemove()
	ResetKeepLocally()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPullTrigger()
	ResetPullTriggers()
	ResetTriggers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Image
type jsiiProxy_Image struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Image) BuildAttribute() ImageBuildOutputReference {
	var returns ImageBuildOutputReference
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) BuildAttributeInput() *ImageBuild {
	var returns *ImageBuild
	_jsii_.Get(
		j,
		"buildAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) ForceRemove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) ForceRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) KeepLocally() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLocally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) KeepLocallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepLocallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Latest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Output() *string {
	var returns *string
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) PullTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) PullTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pullTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) PullTriggers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pullTriggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) PullTriggersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pullTriggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) RepoDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Image) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/docker/r/image docker_image} Resource.
func NewImage(scope constructs.Construct, id *string, config *ImageConfig) Image {
	_init_.Initialize()

	if err := validateNewImageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Image{}

	_jsii_.Create(
		"@cdktf/provider-docker.image.Image",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/docker/r/image docker_image} Resource.
func NewImage_Override(i Image, scope constructs.Construct, id *string, config *ImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.image.Image",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_Image)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Image)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Image)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Image)SetForceRemove(val interface{}) {
	if err := j.validateSetForceRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRemove",
		val,
	)
}

func (j *jsiiProxy_Image)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Image)SetKeepLocally(val interface{}) {
	if err := j.validateSetKeepLocallyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepLocally",
		val,
	)
}

func (j *jsiiProxy_Image)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Image)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Image)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Image)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Image)SetPullTrigger(val *string) {
	if err := j.validateSetPullTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullTrigger",
		val,
	)
}

func (j *jsiiProxy_Image)SetPullTriggers(val *[]*string) {
	if err := j.validateSetPullTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullTriggers",
		val,
	)
}

func (j *jsiiProxy_Image)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Image_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-docker.image.Image",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Image_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-docker.image.Image",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_Image) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_Image) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_Image) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_Image) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_Image) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_Image) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_Image) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_Image) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_Image) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_Image) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_Image) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Image) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_Image) PutBuildAttribute(value *ImageBuild) {
	if err := i.validatePutBuildAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBuildAttribute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Image) ResetBuildAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetBuildAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetForceRemove() {
	_jsii_.InvokeVoid(
		i,
		"resetForceRemove",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetKeepLocally() {
	_jsii_.InvokeVoid(
		i,
		"resetKeepLocally",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetPullTrigger() {
	_jsii_.InvokeVoid(
		i,
		"resetPullTrigger",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetPullTriggers() {
	_jsii_.InvokeVoid(
		i,
		"resetPullTriggers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) ResetTriggers() {
	_jsii_.InvokeVoid(
		i,
		"resetTriggers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Image) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Image) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Image) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Image) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

