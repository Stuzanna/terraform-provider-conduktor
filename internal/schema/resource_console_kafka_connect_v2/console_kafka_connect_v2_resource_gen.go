// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_console_kafka_connect_v2

import (
	"context"
	"fmt"
	"github.com/conduktor/terraform-provider-conduktor/internal/schema/validation"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ConsoleKafkaConnectV2ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster": schema.StringAttribute{
				Required:            true,
				Description:         "Kafka cluster name linked with Kafka current connect server. Must exist in Conduktor Console",
				MarkdownDescription: "Kafka cluster name linked with Kafka current connect server. Must exist in Conduktor Console",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-z\\_\\-.]+$"), ""),
				},
			},
			"labels": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "Kafka connect server labels",
				MarkdownDescription: "Kafka connect server labels",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Kafka connect server name, must be unique, acts as an ID for import",
				MarkdownDescription: "Kafka connect server name, must be unique, acts as an ID for import",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[0-9a-z\\_\\-.]+$"), ""),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"spec": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"display_name": schema.StringAttribute{
						Required:            true,
						Description:         "Kafka connect server display name",
						MarkdownDescription: "Kafka connect server display name",
					},
					"headers": schema.MapAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Description:         "Key-Value HTTP headers to add to requests",
						MarkdownDescription: "Key-Value HTTP headers to add to requests",
					},
					"ignore_untrusted_certificate": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						Description:         "Ignore untrusted certificate for Kafka connect server requests",
						MarkdownDescription: "Ignore untrusted certificate for Kafka connect server requests",
						Default:             booldefault.StaticBool(false),
					},
					"security": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"certificate_chain": schema.StringAttribute{
								Optional:            true,
								Description:         "Kafka connect server mTLS auth certificate chain PEM. Required if security type is `SSLAuth`",
								MarkdownDescription: "Kafka connect server mTLS auth certificate chain PEM. Required if security type is `SSLAuth`",
							},
							"key": schema.StringAttribute{
								Optional:            true,
								Sensitive:           true,
								Description:         "Kafka connect server mTLS auth private key PEM. Required if security type is `SSLAuth`",
								MarkdownDescription: "Kafka connect server mTLS auth private key PEM. Required if security type is `SSLAuth`",
							},
							"password": schema.StringAttribute{
								Optional:            true,
								Sensitive:           true,
								Description:         "Kafka connect server basic auth password. Required if security type is `BasicAuth`",
								MarkdownDescription: "Kafka connect server basic auth password. Required if security type is `BasicAuth`",
							},
							"token": schema.StringAttribute{
								Optional:            true,
								Sensitive:           true,
								Description:         "Kafka connect server bearer token. Required if security type is `BearerToken`",
								MarkdownDescription: "Kafka connect server bearer token. Required if security type is `BearerToken`",
							},
							"type": schema.StringAttribute{
								Required:            true,
								Description:         "Kafka connect server security type. Either `BasicAuth`, `BearerToken`, `SSLAuth`\n\n More detail on our [documentation](https://docs.conduktor.io/platform/reference/resource-reference/console/#kafkaconnectcluster)",
								MarkdownDescription: "Kafka connect server security type. Either `BasicAuth`, `BearerToken`, `SSLAuth`\n\n More detail on our [documentation](https://docs.conduktor.io/platform/reference/resource-reference/console/#kafkaconnectcluster)",
								Validators: []validator.String{
									stringvalidator.OneOf(validation.ValidKafkaConnectSecurityTypes...),
								},
							},
							"username": schema.StringAttribute{
								Optional:            true,
								Description:         "Kafka connect server basic auth username. Required if security type is `BasicAuth`",
								MarkdownDescription: "Kafka connect server basic auth username. Required if security type is `BasicAuth`",
							},
						},
						CustomType: SecurityType{
							ObjectType: types.ObjectType{
								AttrTypes: SecurityValue{}.AttributeTypes(ctx),
							},
						},
						Optional:            true,
						Description:         "Kafka connect server security configuration. One of `BasicAuth`, `BearerToken`, `SSLAuth`",
						MarkdownDescription: "Kafka connect server security configuration. One of `BasicAuth`, `BearerToken`, `SSLAuth`",
					},
					"urls": schema.StringAttribute{
						Required:            true,
						Description:         "URL of a Kafka Connect cluster. **Multiple URLs are not supported for now**",
						MarkdownDescription: "URL of a Kafka Connect cluster. **Multiple URLs are not supported for now**",
					},
				},
				CustomType: SpecType{
					ObjectType: types.ObjectType{
						AttrTypes: SpecValue{}.AttributeTypes(ctx),
					},
				},
			},
		},
	}
}

type ConsoleKafkaConnectV2Model struct {
	Cluster types.String `tfsdk:"cluster"`
	Labels  types.Map    `tfsdk:"labels"`
	Name    types.String `tfsdk:"name"`
	Spec    SpecValue    `tfsdk:"spec"`
}

var _ basetypes.ObjectTypable = SpecType{}

type SpecType struct {
	basetypes.ObjectType
}

func (t SpecType) Equal(o attr.Type) bool {
	other, ok := o.(SpecType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t SpecType) String() string {
	return "SpecType"
}

func (t SpecType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	displayNameAttribute, ok := attributes["display_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`display_name is missing from object`)

		return nil, diags
	}

	displayNameVal, ok := displayNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`display_name expected to be basetypes.StringValue, was: %T`, displayNameAttribute))
	}

	headersAttribute, ok := attributes["headers"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`headers is missing from object`)

		return nil, diags
	}

	headersVal, ok := headersAttribute.(basetypes.MapValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`headers expected to be basetypes.MapValue, was: %T`, headersAttribute))
	}

	ignoreUntrustedCertificateAttribute, ok := attributes["ignore_untrusted_certificate"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ignore_untrusted_certificate is missing from object`)

		return nil, diags
	}

	ignoreUntrustedCertificateVal, ok := ignoreUntrustedCertificateAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ignore_untrusted_certificate expected to be basetypes.BoolValue, was: %T`, ignoreUntrustedCertificateAttribute))
	}

	securityAttribute, ok := attributes["security"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`security is missing from object`)

		return nil, diags
	}

	securityVal, ok := securityAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`security expected to be basetypes.ObjectValue, was: %T`, securityAttribute))
	}

	urlsAttribute, ok := attributes["urls"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`urls is missing from object`)

		return nil, diags
	}

	urlsVal, ok := urlsAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`urls expected to be basetypes.StringValue, was: %T`, urlsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SpecValue{
		DisplayName:                displayNameVal,
		Headers:                    headersVal,
		IgnoreUntrustedCertificate: ignoreUntrustedCertificateVal,
		Security:                   securityVal,
		Urls:                       urlsVal,
		state:                      attr.ValueStateKnown,
	}, diags
}

func NewSpecValueNull() SpecValue {
	return SpecValue{
		state: attr.ValueStateNull,
	}
}

func NewSpecValueUnknown() SpecValue {
	return SpecValue{
		state: attr.ValueStateUnknown,
	}
}

func NewSpecValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (SpecValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing SpecValue Attribute Value",
				"While creating a SpecValue value, a missing attribute value was detected. "+
					"A SpecValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SpecValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid SpecValue Attribute Type",
				"While creating a SpecValue value, an invalid attribute value was detected. "+
					"A SpecValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SpecValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("SpecValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra SpecValue Attribute Value",
				"While creating a SpecValue value, an extra attribute value was detected. "+
					"A SpecValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra SpecValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewSpecValueUnknown(), diags
	}

	displayNameAttribute, ok := attributes["display_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`display_name is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	displayNameVal, ok := displayNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`display_name expected to be basetypes.StringValue, was: %T`, displayNameAttribute))
	}

	headersAttribute, ok := attributes["headers"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`headers is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	headersVal, ok := headersAttribute.(basetypes.MapValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`headers expected to be basetypes.MapValue, was: %T`, headersAttribute))
	}

	ignoreUntrustedCertificateAttribute, ok := attributes["ignore_untrusted_certificate"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ignore_untrusted_certificate is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	ignoreUntrustedCertificateVal, ok := ignoreUntrustedCertificateAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ignore_untrusted_certificate expected to be basetypes.BoolValue, was: %T`, ignoreUntrustedCertificateAttribute))
	}

	securityAttribute, ok := attributes["security"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`security is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	securityVal, ok := securityAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`security expected to be basetypes.ObjectValue, was: %T`, securityAttribute))
	}

	urlsAttribute, ok := attributes["urls"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`urls is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	urlsVal, ok := urlsAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`urls expected to be basetypes.StringValue, was: %T`, urlsAttribute))
	}

	if diags.HasError() {
		return NewSpecValueUnknown(), diags
	}

	return SpecValue{
		DisplayName:                displayNameVal,
		Headers:                    headersVal,
		IgnoreUntrustedCertificate: ignoreUntrustedCertificateVal,
		Security:                   securityVal,
		Urls:                       urlsVal,
		state:                      attr.ValueStateKnown,
	}, diags
}

func NewSpecValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) SpecValue {
	object, diags := NewSpecValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewSpecValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t SpecType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewSpecValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewSpecValueUnknown(), nil
	}

	if in.IsNull() {
		return NewSpecValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewSpecValueMust(SpecValue{}.AttributeTypes(ctx), attributes), nil
}

func (t SpecType) ValueType(ctx context.Context) attr.Value {
	return SpecValue{}
}

var _ basetypes.ObjectValuable = SpecValue{}

type SpecValue struct {
	DisplayName                basetypes.StringValue `tfsdk:"display_name"`
	Headers                    basetypes.MapValue    `tfsdk:"headers"`
	IgnoreUntrustedCertificate basetypes.BoolValue   `tfsdk:"ignore_untrusted_certificate"`
	Security                   basetypes.ObjectValue `tfsdk:"security"`
	Urls                       basetypes.StringValue `tfsdk:"urls"`
	state                      attr.ValueState
}

func (v SpecValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 5)

	var val tftypes.Value
	var err error

	attrTypes["display_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["headers"] = basetypes.MapType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["ignore_untrusted_certificate"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["security"] = basetypes.ObjectType{
		AttrTypes: SecurityValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)
	attrTypes["urls"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 5)

		val, err = v.DisplayName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["display_name"] = val

		val, err = v.Headers.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["headers"] = val

		val, err = v.IgnoreUntrustedCertificate.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ignore_untrusted_certificate"] = val

		val, err = v.Security.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["security"] = val

		val, err = v.Urls.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["urls"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v SpecValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v SpecValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v SpecValue) String() string {
	return "SpecValue"
}

func (v SpecValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var security basetypes.ObjectValue

	if v.Security.IsNull() {
		security = types.ObjectNull(
			SecurityValue{}.AttributeTypes(ctx),
		)
	}

	if v.Security.IsUnknown() {
		security = types.ObjectUnknown(
			SecurityValue{}.AttributeTypes(ctx),
		)
	}

	if !v.Security.IsNull() && !v.Security.IsUnknown() {
		security = types.ObjectValueMust(
			SecurityValue{}.AttributeTypes(ctx),
			v.Security.Attributes(),
		)
	}

	var headersVal basetypes.MapValue
	switch {
	case v.Headers.IsUnknown():
		headersVal = types.MapUnknown(types.StringType)
	case v.Headers.IsNull():
		headersVal = types.MapNull(types.StringType)
	default:
		var d diag.Diagnostics
		headersVal, d = types.MapValue(types.StringType, v.Headers.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"display_name": basetypes.StringType{},
			"headers": basetypes.MapType{
				ElemType: types.StringType,
			},
			"ignore_untrusted_certificate": basetypes.BoolType{},
			"security": basetypes.ObjectType{
				AttrTypes: SecurityValue{}.AttributeTypes(ctx),
			},
			"urls": basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"display_name": basetypes.StringType{},
		"headers": basetypes.MapType{
			ElemType: types.StringType,
		},
		"ignore_untrusted_certificate": basetypes.BoolType{},
		"security": basetypes.ObjectType{
			AttrTypes: SecurityValue{}.AttributeTypes(ctx),
		},
		"urls": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"display_name":                 v.DisplayName,
			"headers":                      headersVal,
			"ignore_untrusted_certificate": v.IgnoreUntrustedCertificate,
			"security":                     security,
			"urls":                         v.Urls,
		})

	return objVal, diags
}

func (v SpecValue) Equal(o attr.Value) bool {
	other, ok := o.(SpecValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.DisplayName.Equal(other.DisplayName) {
		return false
	}

	if !v.Headers.Equal(other.Headers) {
		return false
	}

	if !v.IgnoreUntrustedCertificate.Equal(other.IgnoreUntrustedCertificate) {
		return false
	}

	if !v.Security.Equal(other.Security) {
		return false
	}

	if !v.Urls.Equal(other.Urls) {
		return false
	}

	return true
}

func (v SpecValue) Type(ctx context.Context) attr.Type {
	return SpecType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v SpecValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"display_name": basetypes.StringType{},
		"headers": basetypes.MapType{
			ElemType: types.StringType,
		},
		"ignore_untrusted_certificate": basetypes.BoolType{},
		"security": basetypes.ObjectType{
			AttrTypes: SecurityValue{}.AttributeTypes(ctx),
		},
		"urls": basetypes.StringType{},
	}
}

var _ basetypes.ObjectTypable = SecurityType{}

type SecurityType struct {
	basetypes.ObjectType
}

func (t SecurityType) Equal(o attr.Type) bool {
	other, ok := o.(SecurityType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t SecurityType) String() string {
	return "SecurityType"
}

func (t SecurityType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	certificateChainAttribute, ok := attributes["certificate_chain"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`certificate_chain is missing from object`)

		return nil, diags
	}

	certificateChainVal, ok := certificateChainAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`certificate_chain expected to be basetypes.StringValue, was: %T`, certificateChainAttribute))
	}

	keyAttribute, ok := attributes["key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`key is missing from object`)

		return nil, diags
	}

	keyVal, ok := keyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`key expected to be basetypes.StringValue, was: %T`, keyAttribute))
	}

	passwordAttribute, ok := attributes["password"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`password is missing from object`)

		return nil, diags
	}

	passwordVal, ok := passwordAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`password expected to be basetypes.StringValue, was: %T`, passwordAttribute))
	}

	tokenAttribute, ok := attributes["token"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`token is missing from object`)

		return nil, diags
	}

	tokenVal, ok := tokenAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`token expected to be basetypes.StringValue, was: %T`, tokenAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return nil, diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	usernameAttribute, ok := attributes["username"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`username is missing from object`)

		return nil, diags
	}

	usernameVal, ok := usernameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`username expected to be basetypes.StringValue, was: %T`, usernameAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SecurityValue{
		CertificateChain: certificateChainVal,
		Key:              keyVal,
		Password:         passwordVal,
		Token:            tokenVal,
		SecurityType:     typeVal,
		Username:         usernameVal,
		state:            attr.ValueStateKnown,
	}, diags
}

func NewSecurityValueNull() SecurityValue {
	return SecurityValue{
		state: attr.ValueStateNull,
	}
}

func NewSecurityValueUnknown() SecurityValue {
	return SecurityValue{
		state: attr.ValueStateUnknown,
	}
}

func NewSecurityValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (SecurityValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing SecurityValue Attribute Value",
				"While creating a SecurityValue value, a missing attribute value was detected. "+
					"A SecurityValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SecurityValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid SecurityValue Attribute Type",
				"While creating a SecurityValue value, an invalid attribute value was detected. "+
					"A SecurityValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SecurityValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("SecurityValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra SecurityValue Attribute Value",
				"While creating a SecurityValue value, an extra attribute value was detected. "+
					"A SecurityValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra SecurityValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewSecurityValueUnknown(), diags
	}

	certificateChainAttribute, ok := attributes["certificate_chain"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`certificate_chain is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	certificateChainVal, ok := certificateChainAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`certificate_chain expected to be basetypes.StringValue, was: %T`, certificateChainAttribute))
	}

	keyAttribute, ok := attributes["key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`key is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	keyVal, ok := keyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`key expected to be basetypes.StringValue, was: %T`, keyAttribute))
	}

	passwordAttribute, ok := attributes["password"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`password is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	passwordVal, ok := passwordAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`password expected to be basetypes.StringValue, was: %T`, passwordAttribute))
	}

	tokenAttribute, ok := attributes["token"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`token is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	tokenVal, ok := tokenAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`token expected to be basetypes.StringValue, was: %T`, tokenAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	typeVal, ok := typeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, typeAttribute))
	}

	usernameAttribute, ok := attributes["username"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`username is missing from object`)

		return NewSecurityValueUnknown(), diags
	}

	usernameVal, ok := usernameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`username expected to be basetypes.StringValue, was: %T`, usernameAttribute))
	}

	if diags.HasError() {
		return NewSecurityValueUnknown(), diags
	}

	return SecurityValue{
		CertificateChain: certificateChainVal,
		Key:              keyVal,
		Password:         passwordVal,
		Token:            tokenVal,
		SecurityType:     typeVal,
		Username:         usernameVal,
		state:            attr.ValueStateKnown,
	}, diags
}

func NewSecurityValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) SecurityValue {
	object, diags := NewSecurityValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewSecurityValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t SecurityType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewSecurityValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewSecurityValueUnknown(), nil
	}

	if in.IsNull() {
		return NewSecurityValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewSecurityValueMust(SecurityValue{}.AttributeTypes(ctx), attributes), nil
}

func (t SecurityType) ValueType(ctx context.Context) attr.Value {
	return SecurityValue{}
}

var _ basetypes.ObjectValuable = SecurityValue{}

type SecurityValue struct {
	CertificateChain basetypes.StringValue `tfsdk:"certificate_chain"`
	Key              basetypes.StringValue `tfsdk:"key"`
	Password         basetypes.StringValue `tfsdk:"password"`
	Token            basetypes.StringValue `tfsdk:"token"`
	SecurityType     basetypes.StringValue `tfsdk:"type"`
	Username         basetypes.StringValue `tfsdk:"username"`
	state            attr.ValueState
}

func (v SecurityValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 6)

	var val tftypes.Value
	var err error

	attrTypes["certificate_chain"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["key"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["password"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["token"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["type"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["username"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 6)

		val, err = v.CertificateChain.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["certificate_chain"] = val

		val, err = v.Key.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["key"] = val

		val, err = v.Password.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["password"] = val

		val, err = v.Token.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["token"] = val

		val, err = v.SecurityType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["type"] = val

		val, err = v.Username.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["username"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v SecurityValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v SecurityValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v SecurityValue) String() string {
	return "SecurityValue"
}

func (v SecurityValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"certificate_chain": basetypes.StringType{},
		"key":               basetypes.StringType{},
		"password":          basetypes.StringType{},
		"token":             basetypes.StringType{},
		"type":              basetypes.StringType{},
		"username":          basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"certificate_chain": v.CertificateChain,
			"key":               v.Key,
			"password":          v.Password,
			"token":             v.Token,
			"type":              v.SecurityType,
			"username":          v.Username,
		})

	return objVal, diags
}

func (v SecurityValue) Equal(o attr.Value) bool {
	other, ok := o.(SecurityValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CertificateChain.Equal(other.CertificateChain) {
		return false
	}

	if !v.Key.Equal(other.Key) {
		return false
	}

	if !v.Password.Equal(other.Password) {
		return false
	}

	if !v.Token.Equal(other.Token) {
		return false
	}

	if !v.SecurityType.Equal(other.SecurityType) {
		return false
	}

	if !v.Username.Equal(other.Username) {
		return false
	}

	return true
}

func (v SecurityValue) Type(ctx context.Context) attr.Type {
	return SecurityType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v SecurityValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"certificate_chain": basetypes.StringType{},
		"key":               basetypes.StringType{},
		"password":          basetypes.StringType{},
		"token":             basetypes.StringType{},
		"type":              basetypes.StringType{},
		"username":          basetypes.StringType{},
	}
}
