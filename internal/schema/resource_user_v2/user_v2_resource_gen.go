// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_user_v2

import (
	"context"
	"fmt"
	"github.com/conduktor/terraform-provider-conduktor/internal/schema/validation"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func UserV2ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "User email, must be unique, act as ID for import",
				MarkdownDescription: "User email, must be unique, act as ID for import",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^([\\w\\-_.]*[^.])(@\\w+)(\\.\\w+(\\.\\w+)?[^.\\W])$"), ""),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"spec": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"firstname": schema.StringAttribute{
						Optional:            true,
						Description:         "User firstname",
						MarkdownDescription: "User firstname",
					},
					"lastname": schema.StringAttribute{
						Optional:            true,
						Description:         "User lastname",
						MarkdownDescription: "User lastname",
					},
					"permissions": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster": schema.StringAttribute{
									Optional:            true,
									Description:         "Name of the cluster to apply permission, only required if resource_type is TOPIC, SUBJECT, CONSUMER_GROUP, KAFKA_CONNECT, KSQLDB",
									MarkdownDescription: "Name of the cluster to apply permission, only required if resource_type is TOPIC, SUBJECT, CONSUMER_GROUP, KAFKA_CONNECT, KSQLDB",
								},
								"kafka_connect": schema.StringAttribute{
									Optional:            true,
									Description:         "Name of the Kafka Connect to apply permission, only required if resource_type is KAFKA_CONNECT",
									MarkdownDescription: "Name of the Kafka Connect to apply permission, only required if resource_type is KAFKA_CONNECT",
								},
								"name": schema.StringAttribute{
									Optional:            true,
									Description:         "Name of the resource to apply permission could be a topic, a cluster, a consumer group, etc. depending on resource_type",
									MarkdownDescription: "Name of the resource to apply permission could be a topic, a cluster, a consumer group, etc. depending on resource_type",
								},
								"pattern_type": schema.StringAttribute{
									Optional:            true,
									Description:         "Type of the pattern to apply permission on valid values are: LITERAL, PREFIXED",
									MarkdownDescription: "Type of the pattern to apply permission on valid values are: LITERAL, PREFIXED",
									Validators: []validator.String{
										stringvalidator.OneOf(validation.ValidPermissionPatternTypes...),
									},
								},
								"permissions": schema.ListAttribute{
									ElementType:         types.StringType,
									Required:            true,
									Description:         "Set of all permissions to apply on the resource. See https://docs.conduktor.io/platform/reference/resource-reference/console/#permissions for more details",
									MarkdownDescription: "Set of all permissions to apply on the resource. See https://docs.conduktor.io/platform/reference/resource-reference/console/#permissions for more details",
									Validators: []validator.List{
										listvalidator.UniqueValues(),
										listvalidator.ValueStringsAre(stringvalidator.OneOf(validation.ValidPermissions...)),
									},
								},
								"resource_type": schema.StringAttribute{
									Required:            true,
									Description:         "Type of the resource to apply permission on valid values are: CLUSTER, CONSUMER_GROUP, KAFKA_CONNECT, KSQLDB, PLATFORM, SUBJECT, TOPIC",
									MarkdownDescription: "Type of the resource to apply permission on valid values are: CLUSTER, CONSUMER_GROUP, KAFKA_CONNECT, KSQLDB, PLATFORM, SUBJECT, TOPIC",
									Validators: []validator.String{
										stringvalidator.OneOf(validation.ValidPermissionTypes...),
									},
								},
							},
							CustomType: PermissionsType{
								ObjectType: types.ObjectType{
									AttrTypes: PermissionsValue{}.AttributeTypes(ctx),
								},
							},
						},
						Optional:            true,
						Computed:            true,
						Description:         "Set of all user permissions",
						MarkdownDescription: "Set of all user permissions",
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

type UserV2Model struct {
	Name types.String `tfsdk:"name"`
	Spec SpecValue    `tfsdk:"spec"`
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

	firstnameAttribute, ok := attributes["firstname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`firstname is missing from object`)

		return nil, diags
	}

	firstnameVal, ok := firstnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`firstname expected to be basetypes.StringValue, was: %T`, firstnameAttribute))
	}

	lastnameAttribute, ok := attributes["lastname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`lastname is missing from object`)

		return nil, diags
	}

	lastnameVal, ok := lastnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`lastname expected to be basetypes.StringValue, was: %T`, lastnameAttribute))
	}

	permissionsAttribute, ok := attributes["permissions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permissions is missing from object`)

		return nil, diags
	}

	permissionsVal, ok := permissionsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permissions expected to be basetypes.ListValue, was: %T`, permissionsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SpecValue{
		Firstname:   firstnameVal,
		Lastname:    lastnameVal,
		Permissions: permissionsVal,
		state:       attr.ValueStateKnown,
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

	firstnameAttribute, ok := attributes["firstname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`firstname is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	firstnameVal, ok := firstnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`firstname expected to be basetypes.StringValue, was: %T`, firstnameAttribute))
	}

	lastnameAttribute, ok := attributes["lastname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`lastname is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	lastnameVal, ok := lastnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`lastname expected to be basetypes.StringValue, was: %T`, lastnameAttribute))
	}

	permissionsAttribute, ok := attributes["permissions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permissions is missing from object`)

		return NewSpecValueUnknown(), diags
	}

	permissionsVal, ok := permissionsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permissions expected to be basetypes.ListValue, was: %T`, permissionsAttribute))
	}

	if diags.HasError() {
		return NewSpecValueUnknown(), diags
	}

	return SpecValue{
		Firstname:   firstnameVal,
		Lastname:    lastnameVal,
		Permissions: permissionsVal,
		state:       attr.ValueStateKnown,
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
	Firstname   basetypes.StringValue `tfsdk:"firstname"`
	Lastname    basetypes.StringValue `tfsdk:"lastname"`
	Permissions basetypes.ListValue   `tfsdk:"permissions"`
	state       attr.ValueState
}

func (v SpecValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["firstname"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["lastname"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["permissions"] = basetypes.ListType{
		ElemType: PermissionsValue{}.Type(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.Firstname.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["firstname"] = val

		val, err = v.Lastname.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["lastname"] = val

		val, err = v.Permissions.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["permissions"] = val

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

	permissions := types.ListValueMust(
		PermissionsType{
			basetypes.ObjectType{
				AttrTypes: PermissionsValue{}.AttributeTypes(ctx),
			},
		},
		v.Permissions.Elements(),
	)

	if v.Permissions.IsNull() {
		permissions = types.ListNull(
			PermissionsType{
				basetypes.ObjectType{
					AttrTypes: PermissionsValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	if v.Permissions.IsUnknown() {
		permissions = types.ListUnknown(
			PermissionsType{
				basetypes.ObjectType{
					AttrTypes: PermissionsValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	attributeTypes := map[string]attr.Type{
		"firstname": basetypes.StringType{},
		"lastname":  basetypes.StringType{},
		"permissions": basetypes.ListType{
			ElemType: PermissionsValue{}.Type(ctx),
		},
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
			"firstname":   v.Firstname,
			"lastname":    v.Lastname,
			"permissions": permissions,
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

	if !v.Firstname.Equal(other.Firstname) {
		return false
	}

	if !v.Lastname.Equal(other.Lastname) {
		return false
	}

	if !v.Permissions.Equal(other.Permissions) {
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
		"firstname": basetypes.StringType{},
		"lastname":  basetypes.StringType{},
		"permissions": basetypes.ListType{
			ElemType: PermissionsValue{}.Type(ctx),
		},
	}
}

var _ basetypes.ObjectTypable = PermissionsType{}

type PermissionsType struct {
	basetypes.ObjectType
}

func (t PermissionsType) Equal(o attr.Type) bool {
	other, ok := o.(PermissionsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t PermissionsType) String() string {
	return "PermissionsType"
}

func (t PermissionsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	clusterAttribute, ok := attributes["cluster"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cluster is missing from object`)

		return nil, diags
	}

	clusterVal, ok := clusterAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cluster expected to be basetypes.StringValue, was: %T`, clusterAttribute))
	}

	kafkaConnectAttribute, ok := attributes["kafka_connect"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kafka_connect is missing from object`)

		return nil, diags
	}

	kafkaConnectVal, ok := kafkaConnectAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kafka_connect expected to be basetypes.StringValue, was: %T`, kafkaConnectAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	patternTypeAttribute, ok := attributes["pattern_type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`pattern_type is missing from object`)

		return nil, diags
	}

	patternTypeVal, ok := patternTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`pattern_type expected to be basetypes.StringValue, was: %T`, patternTypeAttribute))
	}

	permissionsAttribute, ok := attributes["permissions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permissions is missing from object`)

		return nil, diags
	}

	permissionsVal, ok := permissionsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permissions expected to be basetypes.ListValue, was: %T`, permissionsAttribute))
	}

	resourceTypeAttribute, ok := attributes["resource_type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`resource_type is missing from object`)

		return nil, diags
	}

	resourceTypeVal, ok := resourceTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`resource_type expected to be basetypes.StringValue, was: %T`, resourceTypeAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return PermissionsValue{
		Cluster:      clusterVal,
		KafkaConnect: kafkaConnectVal,
		Name:         nameVal,
		PatternType:  patternTypeVal,
		Permissions:  permissionsVal,
		ResourceType: resourceTypeVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewPermissionsValueNull() PermissionsValue {
	return PermissionsValue{
		state: attr.ValueStateNull,
	}
}

func NewPermissionsValueUnknown() PermissionsValue {
	return PermissionsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewPermissionsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (PermissionsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing PermissionsValue Attribute Value",
				"While creating a PermissionsValue value, a missing attribute value was detected. "+
					"A PermissionsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PermissionsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid PermissionsValue Attribute Type",
				"While creating a PermissionsValue value, an invalid attribute value was detected. "+
					"A PermissionsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("PermissionsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("PermissionsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra PermissionsValue Attribute Value",
				"While creating a PermissionsValue value, an extra attribute value was detected. "+
					"A PermissionsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra PermissionsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewPermissionsValueUnknown(), diags
	}

	clusterAttribute, ok := attributes["cluster"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`cluster is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	clusterVal, ok := clusterAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`cluster expected to be basetypes.StringValue, was: %T`, clusterAttribute))
	}

	kafkaConnectAttribute, ok := attributes["kafka_connect"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kafka_connect is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	kafkaConnectVal, ok := kafkaConnectAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kafka_connect expected to be basetypes.StringValue, was: %T`, kafkaConnectAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	patternTypeAttribute, ok := attributes["pattern_type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`pattern_type is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	patternTypeVal, ok := patternTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`pattern_type expected to be basetypes.StringValue, was: %T`, patternTypeAttribute))
	}

	permissionsAttribute, ok := attributes["permissions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permissions is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	permissionsVal, ok := permissionsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permissions expected to be basetypes.ListValue, was: %T`, permissionsAttribute))
	}

	resourceTypeAttribute, ok := attributes["resource_type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`resource_type is missing from object`)

		return NewPermissionsValueUnknown(), diags
	}

	resourceTypeVal, ok := resourceTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`resource_type expected to be basetypes.StringValue, was: %T`, resourceTypeAttribute))
	}

	if diags.HasError() {
		return NewPermissionsValueUnknown(), diags
	}

	return PermissionsValue{
		Cluster:      clusterVal,
		KafkaConnect: kafkaConnectVal,
		Name:         nameVal,
		PatternType:  patternTypeVal,
		Permissions:  permissionsVal,
		ResourceType: resourceTypeVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewPermissionsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) PermissionsValue {
	object, diags := NewPermissionsValue(attributeTypes, attributes)

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

		panic("NewPermissionsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t PermissionsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewPermissionsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewPermissionsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewPermissionsValueNull(), nil
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

	return NewPermissionsValueMust(PermissionsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t PermissionsType) ValueType(ctx context.Context) attr.Value {
	return PermissionsValue{}
}

var _ basetypes.ObjectValuable = PermissionsValue{}

type PermissionsValue struct {
	Cluster      basetypes.StringValue `tfsdk:"cluster"`
	KafkaConnect basetypes.StringValue `tfsdk:"kafka_connect"`
	Name         basetypes.StringValue `tfsdk:"name"`
	PatternType  basetypes.StringValue `tfsdk:"pattern_type"`
	Permissions  basetypes.ListValue   `tfsdk:"permissions"`
	ResourceType basetypes.StringValue `tfsdk:"resource_type"`
	state        attr.ValueState
}

func (v PermissionsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 6)

	var val tftypes.Value
	var err error

	attrTypes["cluster"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["kafka_connect"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["pattern_type"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["permissions"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["resource_type"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 6)

		val, err = v.Cluster.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["cluster"] = val

		val, err = v.KafkaConnect.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["kafka_connect"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.PatternType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["pattern_type"] = val

		val, err = v.Permissions.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["permissions"] = val

		val, err = v.ResourceType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["resource_type"] = val

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

func (v PermissionsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v PermissionsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v PermissionsValue) String() string {
	return "PermissionsValue"
}

func (v PermissionsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var permissionsVal basetypes.ListValue
	switch {
	case v.Permissions.IsUnknown():
		permissionsVal = types.ListUnknown(types.StringType)
	case v.Permissions.IsNull():
		permissionsVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		permissionsVal, d = types.ListValue(types.StringType, v.Permissions.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"cluster":       basetypes.StringType{},
			"kafka_connect": basetypes.StringType{},
			"name":          basetypes.StringType{},
			"pattern_type":  basetypes.StringType{},
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"resource_type": basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"cluster":       basetypes.StringType{},
		"kafka_connect": basetypes.StringType{},
		"name":          basetypes.StringType{},
		"pattern_type":  basetypes.StringType{},
		"permissions": basetypes.ListType{
			ElemType: types.StringType,
		},
		"resource_type": basetypes.StringType{},
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
			"cluster":       v.Cluster,
			"kafka_connect": v.KafkaConnect,
			"name":          v.Name,
			"pattern_type":  v.PatternType,
			"permissions":   permissionsVal,
			"resource_type": v.ResourceType,
		})

	return objVal, diags
}

func (v PermissionsValue) Equal(o attr.Value) bool {
	other, ok := o.(PermissionsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Cluster.Equal(other.Cluster) {
		return false
	}

	if !v.KafkaConnect.Equal(other.KafkaConnect) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.PatternType.Equal(other.PatternType) {
		return false
	}

	if !v.Permissions.Equal(other.Permissions) {
		return false
	}

	if !v.ResourceType.Equal(other.ResourceType) {
		return false
	}

	return true
}

func (v PermissionsValue) Type(ctx context.Context) attr.Type {
	return PermissionsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v PermissionsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"cluster":       basetypes.StringType{},
		"kafka_connect": basetypes.StringType{},
		"name":          basetypes.StringType{},
		"pattern_type":  basetypes.StringType{},
		"permissions": basetypes.ListType{
			ElemType: types.StringType,
		},
		"resource_type": basetypes.StringType{},
	}
}
