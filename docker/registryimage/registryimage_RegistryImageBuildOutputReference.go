package registryimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-docker-go/docker/v4/jsii"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v4/registryimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistryImageBuildOutputReference interface {
	cdktf.ComplexObject
	AuthConfig() RegistryImageBuildAuthConfigList
	AuthConfigInput() interface{}
	BuildArgs() *map[string]*string
	SetBuildArgs(val *map[string]*string)
	BuildArgsInput() *map[string]*string
	BuildId() *string
	SetBuildId(val *string)
	BuildIdInput() *string
	CacheFrom() *[]*string
	SetCacheFrom(val *[]*string)
	CacheFromInput() *[]*string
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
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	CpuPeriod() *float64
	SetCpuPeriod(val *float64)
	CpuPeriodInput() *float64
	CpuQuota() *float64
	SetCpuQuota(val *float64)
	CpuQuotaInput() *float64
	CpuSetCpus() *string
	SetCpuSetCpus(val *string)
	CpuSetCpusInput() *string
	CpuSetMems() *string
	SetCpuSetMems(val *string)
	CpuSetMemsInput() *string
	CpuShares() *float64
	SetCpuShares(val *float64)
	CpuSharesInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dockerfile() *string
	SetDockerfile(val *string)
	DockerfileInput() *string
	ExtraHosts() *[]*string
	SetExtraHosts(val *[]*string)
	ExtraHostsInput() *[]*string
	ForceRemove() interface{}
	SetForceRemove(val interface{})
	ForceRemoveInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RegistryImageBuild
	SetInternalValue(val *RegistryImageBuild)
	Isolation() *string
	SetIsolation(val *string)
	IsolationInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Memory() *float64
	SetMemory(val *float64)
	MemoryInput() *float64
	MemorySwap() *float64
	SetMemorySwap(val *float64)
	MemorySwapInput() *float64
	NetworkMode() *string
	SetNetworkMode(val *string)
	NetworkModeInput() *string
	NoCache() interface{}
	SetNoCache(val interface{})
	NoCacheInput() interface{}
	Platform() *string
	SetPlatform(val *string)
	PlatformInput() *string
	PullParent() interface{}
	SetPullParent(val interface{})
	PullParentInput() interface{}
	RemoteContext() *string
	SetRemoteContext(val *string)
	RemoteContextInput() *string
	Remove() interface{}
	SetRemove(val interface{})
	RemoveInput() interface{}
	SecurityOpt() *[]*string
	SetSecurityOpt(val *[]*string)
	SecurityOptInput() *[]*string
	SessionId() *string
	SetSessionId(val *string)
	SessionIdInput() *string
	ShmSize() *float64
	SetShmSize(val *float64)
	ShmSizeInput() *float64
	Squash() interface{}
	SetSquash(val interface{})
	SquashInput() interface{}
	SuppressOutput() interface{}
	SetSuppressOutput(val interface{})
	SuppressOutputInput() interface{}
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
	Ulimit() RegistryImageBuildUlimitList
	UlimitInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutAuthConfig(value interface{})
	PutUlimit(value interface{})
	ResetAuthConfig()
	ResetBuildArgs()
	ResetBuildId()
	ResetCacheFrom()
	ResetCgroupParent()
	ResetCpuPeriod()
	ResetCpuQuota()
	ResetCpuSetCpus()
	ResetCpuSetMems()
	ResetCpuShares()
	ResetDockerfile()
	ResetExtraHosts()
	ResetForceRemove()
	ResetIsolation()
	ResetLabels()
	ResetMemory()
	ResetMemorySwap()
	ResetNetworkMode()
	ResetNoCache()
	ResetPlatform()
	ResetPullParent()
	ResetRemoteContext()
	ResetRemove()
	ResetSecurityOpt()
	ResetSessionId()
	ResetShmSize()
	ResetSquash()
	ResetSuppressOutput()
	ResetTarget()
	ResetUlimit()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RegistryImageBuildOutputReference
type jsiiProxy_RegistryImageBuildOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) AuthConfig() RegistryImageBuildAuthConfigList {
	var returns RegistryImageBuildAuthConfigList
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) AuthConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) BuildArgs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) BuildArgsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) BuildId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) BuildIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CacheFrom() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CacheFromInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CgroupParent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CgroupParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuQuota() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuQuotaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuSetCpus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuSetCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuSetCpusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuSetCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuSetMems() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuSetMems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuSetMemsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuSetMemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuShares() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuShares",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CpuSharesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuSharesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Dockerfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) DockerfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ExtraHosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extraHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ExtraHostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extraHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ForceRemove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ForceRemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) InternalValue() *RegistryImageBuild {
	var returns *RegistryImageBuild
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Isolation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) IsolationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Memory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) MemoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) MemorySwap() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySwap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) MemorySwapInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySwapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) NetworkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) NoCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) NoCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) PullParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pullParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) PullParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pullParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) RemoteContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) RemoteContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Remove() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) RemoveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SecurityOpt() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOpt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SecurityOptInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityOptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SessionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SessionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ShmSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) ShmSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Squash() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"squash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SquashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"squashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SuppressOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) SuppressOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suppressOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Ulimit() RegistryImageBuildUlimitList {
	var returns RegistryImageBuildUlimitList
	_jsii_.Get(
		j,
		"ulimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) UlimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ulimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RegistryImageBuildOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewRegistryImageBuildOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RegistryImageBuildOutputReference {
	_init_.Initialize()

	if err := validateNewRegistryImageBuildOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RegistryImageBuildOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-docker.registryImage.RegistryImageBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRegistryImageBuildOutputReference_Override(r RegistryImageBuildOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-docker.registryImage.RegistryImageBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetBuildArgs(val *map[string]*string) {
	if err := j.validateSetBuildArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildArgs",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetBuildId(val *string) {
	if err := j.validateSetBuildIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildId",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCacheFrom(val *[]*string) {
	if err := j.validateSetCacheFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheFrom",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCgroupParent(val *string) {
	if err := j.validateSetCgroupParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupParent",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCpuPeriod(val *float64) {
	if err := j.validateSetCpuPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuPeriod",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCpuQuota(val *float64) {
	if err := j.validateSetCpuQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuQuota",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCpuSetCpus(val *string) {
	if err := j.validateSetCpuSetCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuSetCpus",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCpuSetMems(val *string) {
	if err := j.validateSetCpuSetMemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuSetMems",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetCpuShares(val *float64) {
	if err := j.validateSetCpuSharesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuShares",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetDockerfile(val *string) {
	if err := j.validateSetDockerfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfile",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetExtraHosts(val *[]*string) {
	if err := j.validateSetExtraHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraHosts",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetForceRemove(val interface{}) {
	if err := j.validateSetForceRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceRemove",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetInternalValue(val *RegistryImageBuild) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetIsolation(val *string) {
	if err := j.validateSetIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isolation",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetMemory(val *float64) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetMemorySwap(val *float64) {
	if err := j.validateSetMemorySwapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySwap",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetNetworkMode(val *string) {
	if err := j.validateSetNetworkModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetNoCache(val interface{}) {
	if err := j.validateSetNoCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noCache",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetPullParent(val interface{}) {
	if err := j.validateSetPullParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullParent",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetRemoteContext(val *string) {
	if err := j.validateSetRemoteContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteContext",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetRemove(val interface{}) {
	if err := j.validateSetRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remove",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetSecurityOpt(val *[]*string) {
	if err := j.validateSetSecurityOptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityOpt",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetSessionId(val *string) {
	if err := j.validateSetSessionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionId",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetShmSize(val *float64) {
	if err := j.validateSetShmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shmSize",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetSquash(val interface{}) {
	if err := j.validateSetSquashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squash",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetSuppressOutput(val interface{}) {
	if err := j.validateSetSuppressOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressOutput",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RegistryImageBuildOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) PutAuthConfig(value interface{}) {
	if err := r.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) PutUlimit(value interface{}) {
	if err := r.validatePutUlimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putUlimit",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetBuildArgs() {
	_jsii_.InvokeVoid(
		r,
		"resetBuildArgs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetBuildId() {
	_jsii_.InvokeVoid(
		r,
		"resetBuildId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCacheFrom() {
	_jsii_.InvokeVoid(
		r,
		"resetCacheFrom",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCgroupParent() {
	_jsii_.InvokeVoid(
		r,
		"resetCgroupParent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCpuPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCpuQuota() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuQuota",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCpuSetCpus() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuSetCpus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCpuSetMems() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuSetMems",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetCpuShares() {
	_jsii_.InvokeVoid(
		r,
		"resetCpuShares",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetDockerfile() {
	_jsii_.InvokeVoid(
		r,
		"resetDockerfile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetExtraHosts() {
	_jsii_.InvokeVoid(
		r,
		"resetExtraHosts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetForceRemove() {
	_jsii_.InvokeVoid(
		r,
		"resetForceRemove",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetIsolation() {
	_jsii_.InvokeVoid(
		r,
		"resetIsolation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		r,
		"resetLabels",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		r,
		"resetMemory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetMemorySwap() {
	_jsii_.InvokeVoid(
		r,
		"resetMemorySwap",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetNetworkMode() {
	_jsii_.InvokeVoid(
		r,
		"resetNetworkMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetNoCache() {
	_jsii_.InvokeVoid(
		r,
		"resetNoCache",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetPlatform() {
	_jsii_.InvokeVoid(
		r,
		"resetPlatform",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetPullParent() {
	_jsii_.InvokeVoid(
		r,
		"resetPullParent",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetRemoteContext() {
	_jsii_.InvokeVoid(
		r,
		"resetRemoteContext",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetRemove() {
	_jsii_.InvokeVoid(
		r,
		"resetRemove",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetSecurityOpt() {
	_jsii_.InvokeVoid(
		r,
		"resetSecurityOpt",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetSessionId() {
	_jsii_.InvokeVoid(
		r,
		"resetSessionId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetShmSize() {
	_jsii_.InvokeVoid(
		r,
		"resetShmSize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetSquash() {
	_jsii_.InvokeVoid(
		r,
		"resetSquash",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetSuppressOutput() {
	_jsii_.InvokeVoid(
		r,
		"resetSuppressOutput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		r,
		"resetTarget",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetUlimit() {
	_jsii_.InvokeVoid(
		r,
		"resetUlimit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RegistryImageBuildOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

