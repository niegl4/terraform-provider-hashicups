package provider

import (
	_init_ "cdk.tf/go/stack/generated/edu/hashicups/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/edu/hashicups/provider/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://hashicorp.com/providers/edu/hashicups/latest/docs hashicups}.
type HashicupsProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetHost()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetUsername()
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

// The jsii proxy struct for HashicupsProvider
type jsiiProxy_HashicupsProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_HashicupsProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HashicupsProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://hashicorp.com/providers/edu/hashicups/latest/docs hashicups} Resource.
func NewHashicupsProvider(scope constructs.Construct, id *string, config *HashicupsProviderConfig) HashicupsProvider {
	_init_.Initialize()

	if err := validateNewHashicupsProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HashicupsProvider{}

	_jsii_.Create(
		"hashicups.provider.HashicupsProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://hashicorp.com/providers/edu/hashicups/latest/docs hashicups} Resource.
func NewHashicupsProvider_Override(h HashicupsProvider, scope constructs.Construct, id *string, config *HashicupsProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicups.provider.HashicupsProvider",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HashicupsProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_HashicupsProvider)SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_HashicupsProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HashicupsProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a HashicupsProvider resource upon running "cdktf plan <stack-name>".
func HashicupsProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHashicupsProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"hashicups.provider.HashicupsProvider",
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
func HashicupsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHashicupsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"hashicups.provider.HashicupsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HashicupsProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHashicupsProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"hashicups.provider.HashicupsProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HashicupsProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHashicupsProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"hashicups.provider.HashicupsProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HashicupsProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicups.provider.HashicupsProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HashicupsProvider) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HashicupsProvider) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HashicupsProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		h,
		"resetAlias",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HashicupsProvider) ResetHost() {
	_jsii_.InvokeVoid(
		h,
		"resetHost",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HashicupsProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HashicupsProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HashicupsProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		h,
		"resetUsername",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HashicupsProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HashicupsProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HashicupsProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HashicupsProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HashicupsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HashicupsProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

