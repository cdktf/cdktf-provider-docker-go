// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package buildxbuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v12/buildxbuilder/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder docker_buildx_builder}.
type BuildxBuilder interface {
	cdktf.TerraformResource
	Append() interface{}
	SetAppend(val interface{})
	AppendInput() interface{}
	Bootstrap() interface{}
	SetBootstrap(val interface{})
	BootstrapInput() interface{}
	BuildkitConfig() *string
	SetBuildkitConfig(val *string)
	BuildkitConfigInput() *string
	BuildkitFlags() *string
	SetBuildkitFlags(val *string)
	BuildkitFlagsInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerContainer() BuildxBuilderDockerContainerOutputReference
	DockerContainerInput() *BuildxBuilderDockerContainer
	Driver() *string
	SetDriver(val *string)
	DriverInput() *string
	DriverOptions() *map[string]*string
	SetDriverOptions(val *map[string]*string)
	DriverOptionsInput() *map[string]*string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Kubernetes() BuildxBuilderKubernetesOutputReference
	KubernetesInput() *BuildxBuilderKubernetes
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeAttribute() *string
	SetNodeAttribute(val *string)
	NodeAttributeInput() *string
	Platform() *[]*string
	SetPlatform(val *[]*string)
	PlatformInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Remote() BuildxBuilderRemoteOutputReference
	RemoteInput() *BuildxBuilderRemote
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Use() interface{}
	SetUse(val interface{})
	UseInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDockerContainer(value *BuildxBuilderDockerContainer)
	PutKubernetes(value *BuildxBuilderKubernetes)
	PutRemote(value *BuildxBuilderRemote)
	ResetAppend()
	ResetBootstrap()
	ResetBuildkitConfig()
	ResetBuildkitFlags()
	ResetDockerContainer()
	ResetDriver()
	ResetDriverOptions()
	ResetEndpoint()
	ResetId()
	ResetKubernetes()
	ResetName()
	ResetNodeAttribute()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRemote()
	ResetUse()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BuildxBuilder
type jsiiProxy_BuildxBuilder struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BuildxBuilder) Append() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"append",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) AppendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Bootstrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) BootstrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) BuildkitConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildkitConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) BuildkitConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildkitConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) BuildkitFlags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildkitFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) BuildkitFlagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildkitFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DockerContainer() BuildxBuilderDockerContainerOutputReference {
	var returns BuildxBuilderDockerContainerOutputReference
	_jsii_.Get(
		j,
		"dockerContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DockerContainerInput() *BuildxBuilderDockerContainer {
	var returns *BuildxBuilderDockerContainer
	_jsii_.Get(
		j,
		"dockerContainerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DriverOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"driverOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) DriverOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"driverOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Kubernetes() BuildxBuilderKubernetesOutputReference {
	var returns BuildxBuilderKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) KubernetesInput() *BuildxBuilderKubernetes {
	var returns *BuildxBuilderKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) NodeAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) NodeAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Platform() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) PlatformInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Remote() BuildxBuilderRemoteOutputReference {
	var returns BuildxBuilderRemoteOutputReference
	_jsii_.Get(
		j,
		"remote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) RemoteInput() *BuildxBuilderRemote {
	var returns *BuildxBuilderRemote
	_jsii_.Get(
		j,
		"remoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) Use() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildxBuilder) UseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder docker_buildx_builder} Resource.
func NewBuildxBuilder(scope constructs.Construct, id *string, config *BuildxBuilderConfig) BuildxBuilder {
	_init_.Initialize()

	if err := validateNewBuildxBuilderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuildxBuilder{}

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/kreuzwerker/docker/3.6.1/docs/resources/buildx_builder docker_buildx_builder} Resource.
func NewBuildxBuilder_Override(b BuildxBuilder, scope constructs.Construct, id *string, config *BuildxBuilderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetAppend(val interface{}) {
	if err := j.validateSetAppendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"append",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetBootstrap(val interface{}) {
	if err := j.validateSetBootstrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrap",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetBuildkitConfig(val *string) {
	if err := j.validateSetBuildkitConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildkitConfig",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetBuildkitFlags(val *string) {
	if err := j.validateSetBuildkitFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildkitFlags",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetDriverOptions(val *map[string]*string) {
	if err := j.validateSetDriverOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverOptions",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetNodeAttribute(val *string) {
	if err := j.validateSetNodeAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeAttribute",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetPlatform(val *[]*string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BuildxBuilder)SetUse(val interface{}) {
	if err := j.validateSetUseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use",
		val,
	)
}

// Generates CDKTF code for importing a BuildxBuilder resource upon running "cdktf plan <stack-name>".
func BuildxBuilder_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBuildxBuilder_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func BuildxBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBuildxBuilder_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BuildxBuilder_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBuildxBuilder_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BuildxBuilder_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBuildxBuilder_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BuildxBuilder_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-docker.buildxBuilder.BuildxBuilder",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BuildxBuilder) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BuildxBuilder) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BuildxBuilder) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BuildxBuilder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BuildxBuilder) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BuildxBuilder) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BuildxBuilder) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BuildxBuilder) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BuildxBuilder) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BuildxBuilder) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BuildxBuilder) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BuildxBuilder) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BuildxBuilder) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BuildxBuilder) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BuildxBuilder) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BuildxBuilder) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BuildxBuilder) PutDockerContainer(value *BuildxBuilderDockerContainer) {
	if err := b.validatePutDockerContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDockerContainer",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilder) PutKubernetes(value *BuildxBuilderKubernetes) {
	if err := b.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilder) PutRemote(value *BuildxBuilderRemote) {
	if err := b.validatePutRemoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRemote",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetAppend() {
	_jsii_.InvokeVoid(
		b,
		"resetAppend",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetBootstrap() {
	_jsii_.InvokeVoid(
		b,
		"resetBootstrap",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetBuildkitConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetBuildkitConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetBuildkitFlags() {
	_jsii_.InvokeVoid(
		b,
		"resetBuildkitFlags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetDockerContainer() {
	_jsii_.InvokeVoid(
		b,
		"resetDockerContainer",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetDriver() {
	_jsii_.InvokeVoid(
		b,
		"resetDriver",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetDriverOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetDriverOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetKubernetes() {
	_jsii_.InvokeVoid(
		b,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetNodeAttribute() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeAttribute",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetPlatform() {
	_jsii_.InvokeVoid(
		b,
		"resetPlatform",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetRemote() {
	_jsii_.InvokeVoid(
		b,
		"resetRemote",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) ResetUse() {
	_jsii_.InvokeVoid(
		b,
		"resetUse",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildxBuilder) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildxBuilder) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

