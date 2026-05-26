package converter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"

	protoc_gen_jsonschema "github.com/ikstewa/protoc-gen-jsonschema"
	protoc_gen_validate "github.com/envoyproxy/protoc-gen-validate/validate"
)

// JSON Schema primitive type vocabulary (Draft 2020-12). Defined locally so we
// don't depend on github.com/xeipuuv/gojsonschema just for these string
// constants. These values are part of the stable JSON Schema spec.
const (
	jsTypeNull    = "null"
	jsTypeBoolean = "boolean"
	jsTypeObject  = "object"
	jsTypeArray   = "array"
	jsTypeNumber  = "number"
	jsTypeString  = "string"
	jsTypeInteger = "integer"
)

var (
	globalPkg = newProtoPackage(nil, "")

	wellKnownTypes = map[string]bool{
		"BoolValue":   true,
		"BytesValue":  true,
		"DoubleValue": true,
		"Duration":    true,
		"FloatValue":  true,
		"Int32Value":  true,
		"Int64Value":  true,
		"ListValue":   true,
		"StringValue": true,
		"Struct":      true,
		"UInt32Value": true,
		"UInt64Value": true,
		"Value":       true,
	}
)

func (c *Converter) registerEnum(pkgName string, enum *descriptor.EnumDescriptorProto) {
	pkg := globalPkg
	if pkgName != "" {
		for _, node := range strings.Split(pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}
			child, ok := pkg.children[node]
			if !ok {
				child = newProtoPackage(pkg, node)
				pkg.children[node] = child
			}
			pkg = child
		}
	}
	pkg.enums[enum.GetName()] = enum
}

func (c *Converter) registerType(pkgName string, msgDesc *descriptor.DescriptorProto) {
	pkg := globalPkg
	if pkgName != "" {
		for _, node := range strings.Split(pkgName, ".") {
			if pkg == globalPkg && node == "" {
				// Skips leading "."
				continue
			}
			child, ok := pkg.children[node]
			if !ok {
				child = newProtoPackage(pkg, node)
				pkg.children[node] = child
			}
			pkg = child
		}
	}
	pkg.types[msgDesc.GetName()] = msgDesc
}

// Convert a proto "field" (essentially a type-switch with some recursion):
func (c *Converter) convertField(curPkg *ProtoPackage, desc *descriptor.FieldDescriptorProto, msgDesc *descriptor.DescriptorProto, duplicatedMessages map[*descriptor.DescriptorProto]string, messageFlags ConverterFlags) (*jsonschema.Schema, error) {

	// Prepare a new jsonschema.Schema for our eventual return value:
	jsonSchemaType := &jsonschema.Schema{}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetField(desc); src != nil {
		jsonSchemaType.Title, jsonSchemaType.Description = c.formatTitleAndDescription(nil, src)
	}

	// Switch the types, and pick a JSONSchema equivalent:
	switch desc.GetType() {

	// Float32:
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if messageFlags.AllowNullValues {
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeNumber, Format: "float"},
				}
			} else {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeNumber},
				}
			}
		} else {
			jsonSchemaType.Type = jsTypeNumber
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.Format = "float"
			}
		}
		if c.Flags.IncludeNumericBounds {
			applyNumericBounds(jsonSchemaType, desc.GetType())
		}

	// Double:
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		if messageFlags.AllowNullValues {
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeNumber, Format: "double"},
				}
			} else {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeNumber},
				}
			}
		} else {
			jsonSchemaType.Type = jsTypeNumber
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.Format = "double"
			}
		}
		if c.Flags.IncludeNumericBounds {
			applyNumericBounds(jsonSchemaType, desc.GetType())
		}

	// Int32:
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SINT32:
		if messageFlags.AllowNullValues {
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeInteger, Format: "int32"},
				}
			} else {
				jsonSchemaType.OneOf = []*jsonschema.Schema{
					{Type: jsTypeNull},
					{Type: jsTypeInteger},
				}
			}
		} else {
			jsonSchemaType.Type = jsTypeInteger
			if c.Flags.IncludeNumericFormat {
				jsonSchemaType.Format = "int32"
			}
		}
		if c.Flags.IncludeNumericBounds {
			applyNumericBounds(jsonSchemaType, desc.GetType())
		}

	// Int64:
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT64:

		// As integer:
		if c.Flags.DisallowBigIntsAsStrings {
			if messageFlags.AllowNullValues {
				if c.Flags.IncludeNumericFormat {
					jsonSchemaType.OneOf = []*jsonschema.Schema{
						{Type: jsTypeInteger, Format: "int64"},
						{Type: jsTypeNull},
					}
				} else {
					jsonSchemaType.OneOf = []*jsonschema.Schema{
						{Type: jsTypeInteger},
						{Type: jsTypeNull},
					}
				}
			} else {
				jsonSchemaType.Type = jsTypeInteger
				if c.Flags.IncludeNumericFormat {
					jsonSchemaType.Format = "int64"
				}
			}
		}

		// As string:
		if !c.Flags.DisallowBigIntsAsStrings {
			if messageFlags.AllowNullValues {
				if c.Flags.IncludeNumericFormat {
					jsonSchemaType.OneOf = []*jsonschema.Schema{
						{Type: jsTypeString, Format: "int64"},
						{Type: jsTypeNull},
					}
				} else {
					jsonSchemaType.OneOf = []*jsonschema.Schema{
						{Type: jsTypeString},
						{Type: jsTypeNull},
					}
				}
			} else {
				jsonSchemaType.Type = jsTypeString
				if c.Flags.IncludeNumericFormat {
					jsonSchemaType.Format = "int64"
				}
			}
		}
		if c.Flags.IncludeNumericBounds {
			applyNumericBounds(jsonSchemaType, desc.GetType())
		}

	// String:
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		stringDef := &jsonschema.Schema{Type: jsTypeString}

		// Custom field options from protoc-gen-jsonschema:
		if opt := proto.GetExtension(desc.GetOptions(), protoc_gen_jsonschema.E_FieldOptions); opt != nil {
			if fieldOptions, ok := opt.(*protoc_gen_jsonschema.FieldOptions); ok {
				stringDef.MinLength = ptrUint64(uint64(fieldOptions.GetMinLength()))
				stringDef.MaxLength = ptrUint64(uint64(fieldOptions.GetMaxLength()))
				stringDef.Pattern = fieldOptions.GetPattern()
			}
		}

		// Custom field options from protoc-gen-validate:
		if opt := proto.GetExtension(desc.GetOptions(), protoc_gen_validate.E_Rules); opt != nil {
			if fieldRules, ok := opt.(*protoc_gen_validate.FieldRules); fieldRules != nil && ok {
				if stringRules := fieldRules.GetString_(); stringRules != nil {
					stringDef.MaxLength = ptrUint64(uint64(stringRules.GetMaxLen()))
					stringDef.MinLength = ptrUint64(uint64(stringRules.GetMinLen()))
					stringDef.Pattern = stringRules.GetPattern()
				}
			}
		}

		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeNull},
				stringDef,
			}
		} else {
			jsonSchemaType.Type = stringDef.Type
			jsonSchemaType.MinLength = stringDef.MinLength
			jsonSchemaType.MaxLength = stringDef.MaxLength
			jsonSchemaType.Pattern = stringDef.Pattern
		}

	// Bytes:
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeNull},
				{
					Type:           jsTypeString,
					Format:         "binary",
					ContentEncoding: "base64",
				},
			}
		} else {
			jsonSchemaType.Type = jsTypeString
			jsonSchemaType.Format = "binary"
			jsonSchemaType.ContentEncoding = "base64"
		}

	// ENUM:
	case descriptor.FieldDescriptorProto_TYPE_ENUM:

		// Go through all the enums we have, see if we can match any to this field.
		fullEnumIdentifier := strings.TrimPrefix(desc.GetTypeName(), ".")
		matchedEnum, _, ok := c.lookupEnum(curPkg, fullEnumIdentifier)
		if !ok {
			return nil, fmt.Errorf("unable to resolve enum type: %s", desc.GetType().String())
		}

		// We already have a converter for standalone ENUMs, so just use that:
		enumSchema, err := c.convertEnumType(matchedEnum, messageFlags)
		if err != nil {
			switch err {
			case errIgnored:
			default:
				return nil, err
			}
		}

		jsonSchemaType = &enumSchema

	// Bool:
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeNull},
				{Type: jsTypeBoolean},
			}
		} else {
			jsonSchemaType.Type = jsTypeBoolean
		}

	// Group (object):
	case descriptor.FieldDescriptorProto_TYPE_GROUP, descriptor.FieldDescriptorProto_TYPE_MESSAGE:

		switch desc.GetTypeName() {
		// Make sure that durations match a particular string pattern (eg 3.4s):
		case ".google.protobuf.Duration":
			jsonSchemaType.Type = jsTypeString
			jsonSchemaType.Format = "regex"
			jsonSchemaType.Pattern = `^([0-9]+\.?[0-9]*|\.[0-9]+)s$`
		case ".google.protobuf.Timestamp":
			jsonSchemaType.Type = jsTypeString
			jsonSchemaType.Format = "date-time"
		case ".google.protobuf.Value", ".google.protobuf.Struct":
			jsonSchemaType.Type = jsTypeObject
			jsonSchemaType.AdditionalProperties = schemaAllowAny()
		default:
			jsonSchemaType.Type = jsTypeObject
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL {
				jsonSchemaType.AdditionalProperties = schemaAllowAny()
			}
			if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
				jsonSchemaType.AdditionalProperties = schemaAllowNone()
			}
			if messageFlags.DisallowAdditionalProperties {
				jsonSchemaType.AdditionalProperties = schemaAllowNone()
			}
		}

	default:
		return nil, fmt.Errorf("unrecognized field type: %s", desc.GetType().String())
	}

	// Recurse basic array:
	if desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED && jsonSchemaType.Type != jsTypeObject {
		jsonSchemaType.Items = &jsonschema.Schema{}

		// Custom field options from protoc-gen-validate:
		if opt := proto.GetExtension(desc.GetOptions(), protoc_gen_validate.E_Rules); opt != nil {
			if fieldRules, ok := opt.(*protoc_gen_validate.FieldRules); fieldRules != nil && ok {
				if repeatedRules := fieldRules.GetRepeated(); repeatedRules != nil {
					jsonSchemaType.MaxItems = ptrUint64(uint64(repeatedRules.GetMaxItems()))
					jsonSchemaType.MinItems = ptrUint64(uint64(repeatedRules.GetMinItems()))
				}
			}
		}

		if len(jsonSchemaType.Enum) > 0 {
			jsonSchemaType.Items.Enum = jsonSchemaType.Enum
			jsonSchemaType.Enum = nil
			jsonSchemaType.Items.OneOf = nil
		} else {
			jsonSchemaType.Items.Type = jsonSchemaType.Type
			jsonSchemaType.Items.OneOf = jsonSchemaType.OneOf
			jsonSchemaType.Items.Format = jsonSchemaType.Format
			jsonSchemaType.Items.Minimum = jsonSchemaType.Minimum
			jsonSchemaType.Items.Maximum = jsonSchemaType.Maximum
			jsonSchemaType.Format = ""
			jsonSchemaType.Minimum = ""
			jsonSchemaType.Maximum = ""
		}

		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeNull},
				{Type: jsTypeArray},
			}
		} else {
			jsonSchemaType.Type = jsTypeArray
			jsonSchemaType.OneOf = []*jsonschema.Schema{}
		}
		return jsonSchemaType, nil
	}

	// Recurse nested objects / arrays of objects (if necessary):
	if jsonSchemaType.Type == jsTypeObject {

		recordType, pkgName, ok := c.lookupType(curPkg, desc.GetTypeName())
		if !ok {
			return nil, fmt.Errorf("no such message type named %s", desc.GetTypeName())
		}

		// Recurse the recordType:
		recursedJSONSchemaType, err := c.recursiveConvertMessageType(curPkg, recordType, pkgName, duplicatedMessages, false)
		if err != nil {
			return nil, err
		}

		// Maps, arrays, and objects are structured in different ways:
		switch {

		// Maps:
		case recordType.Options.GetMapEntry():
			c.logger.
				WithField("field_name", recordType.GetName()).
				WithField("msgDesc_name", *msgDesc.Name).
				Tracef("Is a map")

			if recursedJSONSchemaType.Properties == nil {
				return nil, fmt.Errorf("Unable to find properties of MAP type")
			}

			// Make sure we have a "value":
			value, valuePresent := recursedJSONSchemaType.Properties.Get("value")
			if !valuePresent {
				return nil, fmt.Errorf("Unable to find 'value' property of MAP type")
			}

			// Pass the "value" schema through as AdditionalProperties (now typed
			// *jsonschema.Schema in invopop, no JSON-marshal indirection needed):
			jsonSchemaType.AdditionalProperties = value

		// Arrays:
		case desc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED:
			jsonSchemaType.Items = recursedJSONSchemaType
			jsonSchemaType.Type = jsTypeArray

			// Build up the list of required fields:
			if messageFlags.AllFieldsRequired && len(recursedJSONSchemaType.OneOf) == 0 && recursedJSONSchemaType.Properties != nil {
				for pair := recursedJSONSchemaType.Properties.Oldest(); pair != nil; pair = pair.Next() {
					jsonSchemaType.Items.Required = append(jsonSchemaType.Items.Required, pair.Key)
				}
			}
			jsonSchemaType.Items.Required = dedupe(jsonSchemaType.Items.Required)

		// Not maps, not arrays:
		default:

			// If we've got optional types then just take those:
			if recursedJSONSchemaType.OneOf != nil {
				return recursedJSONSchemaType, nil
			}

			// If we're not an object then set the type from whatever we recursed:
			if recursedJSONSchemaType.Type != jsTypeObject {
				jsonSchemaType.Type = recursedJSONSchemaType.Type
			}

			// Assume the attrbutes of the recursed value:
			jsonSchemaType.Properties = recursedJSONSchemaType.Properties
			jsonSchemaType.Ref = recursedJSONSchemaType.Ref
			jsonSchemaType.Required = recursedJSONSchemaType.Required

			// In Draft 2020-12, sibling keywords on a $ref schema are applied
			// (unlike Draft 4-7 where $ref was exclusive). The wrapping schema
			// here previously set Type=object and AdditionalProperties=...
			// based on the field's label; those become real constraints under
			// 2020-12 (e.g. additionalProperties: {not: true} forbids all
			// properties on the referenced object) and corrupt validation.
			// Clear them so the $ref alone delegates to the target schema.
			if jsonSchemaType.Ref != "" {
				jsonSchemaType.Type = ""
				jsonSchemaType.AdditionalProperties = nil
			}

			// Build up the list of required fields:
			if messageFlags.AllFieldsRequired && len(recursedJSONSchemaType.OneOf) == 0 && recursedJSONSchemaType.Properties != nil {
				for pair := recursedJSONSchemaType.Properties.Oldest(); pair != nil; pair = pair.Next() {
					jsonSchemaType.Required = append(jsonSchemaType.Required, pair.Key)
				}
			}
		}

		// Optionally allow NULL values:
		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeNull},
				{Type: jsonSchemaType.Type, Items: jsonSchemaType.Items},
			}
			jsonSchemaType.Type = ""
			jsonSchemaType.Items = nil
		}
	}

	jsonSchemaType.Required = dedupe(jsonSchemaType.Required)

	return jsonSchemaType, nil
}

// Converts a proto "MESSAGE" into a JSON-Schema:
func (c *Converter) convertMessageType(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto) (*jsonschema.Schema, error) {

	// Get a list of any nested messages in our schema:
	duplicatedMessages, err := c.findNestedMessages(curPkg, msgDesc)
	if err != nil {
		return nil, err
	}

	// Build up a list of JSONSchema type definitions for every message:
	definitions := jsonschema.Definitions{}
	for refmsgDesc, nameWithPackage := range duplicatedMessages {
		var typeName string
		if c.Flags.TypeNamesWithNoPackage {
			typeName = refmsgDesc.GetName();
		} else {
			typeName = nameWithPackage;
		}
		refType, err := c.recursiveConvertMessageType(curPkg, refmsgDesc, "", duplicatedMessages, true)
		if err != nil {
			return nil, err
		}

		// Add the schema to our definitions:
		definitions[typeName] = refType
	}

	// Put together a JSON schema with our discovered definitions, and a $ref for the root type:
	newJSONSchema := &jsonschema.Schema{
		Ref:         fmt.Sprintf("%s%s", c.refPrefix, msgDesc.GetName()),
		Version:     c.schemaVersion,
		Definitions: definitions,
	}

	return newJSONSchema, nil
}

// findNestedMessages takes a message, and returns a map mapping pointers to messages nested within it:
// these messages become definitions which can be referenced (instead of repeating them every time they're used)
func (c *Converter) findNestedMessages(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto) (map[*descriptor.DescriptorProto]string, error) {

	// Get a list of all nested messages, and how often they occur:
	nestedMessages := make(map[*descriptor.DescriptorProto]string)
	if err := c.recursiveFindNestedMessages(curPkg, msgDesc, msgDesc.GetName(), nestedMessages); err != nil {
		return nil, err
	}

	// Now filter them:
	result := make(map[*descriptor.DescriptorProto]string)
	for message, messageName := range nestedMessages {
		if !message.GetOptions().GetMapEntry() && !strings.HasPrefix(messageName, ".google.protobuf.") {
			result[message] = strings.TrimLeft(messageName, ".")
		}
	}

	return result, nil
}

func (c *Converter) recursiveFindNestedMessages(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto, typeName string, nestedMessages map[*descriptor.DescriptorProto]string) error {
	if _, present := nestedMessages[msgDesc]; present {
		return nil
	}
	nestedMessages[msgDesc] = typeName

	for _, desc := range msgDesc.GetField() {
		descType := desc.GetType()
		if descType != descriptor.FieldDescriptorProto_TYPE_MESSAGE && descType != descriptor.FieldDescriptorProto_TYPE_GROUP {
			// no nested messages
			continue
		}

		typeName := desc.GetTypeName()
		recordType, _, ok := c.lookupType(curPkg, typeName)
		if !ok {
			return fmt.Errorf("no such message type named %s", typeName)
		}
		if err := c.recursiveFindNestedMessages(curPkg, recordType, typeName, nestedMessages); err != nil {
			return err
		}
	}

	return nil
}

func (c *Converter) recursiveConvertMessageType(curPkg *ProtoPackage, msgDesc *descriptor.DescriptorProto, pkgName string, duplicatedMessages map[*descriptor.DescriptorProto]string, ignoreDuplicatedMessages bool) (*jsonschema.Schema, error) {

	// Prepare a new jsonschema:
	jsonSchemaType := new(jsonschema.Schema)

	// Set some per-message flags from config and options:
	messageFlags := c.Flags

	// Custom message options from protoc-gen-jsonschema:
	if opt := proto.GetExtension(msgDesc.GetOptions(), protoc_gen_jsonschema.E_MessageOptions); opt != nil {
		if messageOptions, ok := opt.(*protoc_gen_jsonschema.MessageOptions); ok {

			// AllFieldsRequired:
			if messageOptions.GetAllFieldsRequired() {
				messageFlags.AllFieldsRequired = true
			}

			// AllowNullValues:
			if messageOptions.GetAllowNullValues() {
				messageFlags.AllowNullValues = true
			}

			// DisallowAdditionalProperties:
			if messageOptions.GetDisallowAdditionalProperties() {
				messageFlags.DisallowAdditionalProperties = true
			}

			// ENUMs as constants:
			if messageOptions.GetEnumsAsConstants() {
				messageFlags.EnumsAsConstants = true
			}
		}
	}

	// Generate a description from src comments (if available)
	if src := c.sourceInfo.GetMessage(msgDesc); src != nil {
		jsonSchemaType.Title, jsonSchemaType.Description = c.formatTitleAndDescription(strPtr(msgDesc.GetName()), src)
	}

	// Handle google's well-known types:
	if msgDesc.Name != nil && wellKnownTypes[*msgDesc.Name] && pkgName == ".google.protobuf" {
		switch *msgDesc.Name {
		case "DoubleValue", "FloatValue":
			jsonSchemaType.Type = jsTypeNumber
		case "Int32Value", "UInt32Value":
			jsonSchemaType.Type = jsTypeInteger
		case "Int64Value", "UInt64Value":
			// BigInt as ints
			if messageFlags.DisallowBigIntsAsStrings {
				jsonSchemaType.Type = jsTypeInteger
			} else {

				// BigInt as strings
				jsonSchemaType.Type = jsTypeString
			}

		case "BoolValue":
			jsonSchemaType.Type = jsTypeBoolean
		case "BytesValue", "StringValue":
			jsonSchemaType.Type = jsTypeString
		case "Value":
			jsonSchemaType.OneOf = []*jsonschema.Schema{
				{Type: jsTypeArray},
				{Type: jsTypeBoolean},
				{Type: jsTypeNumber},
				{Type: jsTypeObject},
				{Type: jsTypeString},
			}
			// jsonSchemaType.AdditionalProperties = []byte("true")
		case "Duration":
			jsonSchemaType.Type = jsTypeString
		case "Struct":
			jsonSchemaType.Type = jsTypeObject
			// jsonSchemaType.AdditionalProperties = []byte("true")
		case "ListValue":
			jsonSchemaType.Type = jsTypeArray
		}

		// If we're allowing nulls then prepare a OneOf:
		if messageFlags.AllowNullValues {
			jsonSchemaType.OneOf = append(jsonSchemaType.OneOf, &jsonschema.Schema{Type: jsTypeNull}, &jsonschema.Schema{Type: jsonSchemaType.Type})
			// and clear the Type that was previously set.
			jsonSchemaType.Type = ""
			return jsonSchemaType, nil
		}

		// Otherwise just return this simple type:
		return jsonSchemaType, nil
	}

	// Set defaults:
	jsonSchemaType.Properties = orderedmap.New[string, *jsonschema.Schema]()

	// Look up references:
	if nameWithPackage, ok := duplicatedMessages[msgDesc]; ok && !ignoreDuplicatedMessages {
		var typeName string
		if c.Flags.TypeNamesWithNoPackage {
			typeName = msgDesc.GetName();
		} else {
			typeName = nameWithPackage;
		}
		return &jsonschema.Schema{
			Ref: fmt.Sprintf("%s%s", c.refPrefix, typeName),
		}, nil
	}

	// Optionally allow NULL values:
	if messageFlags.AllowNullValues {
		jsonSchemaType.OneOf = []*jsonschema.Schema{
			{Type: jsTypeNull},
			{Type: jsTypeObject},
		}
	} else {
		jsonSchemaType.Type = jsTypeObject
	}

	// disallowAdditionalProperties will prevent validation where extra fields are found (outside of the schema):
	if messageFlags.DisallowAdditionalProperties {
		jsonSchemaType.AdditionalProperties = schemaAllowNone()
	} else {
		jsonSchemaType.AdditionalProperties = schemaAllowAny()
	}

	c.logger.WithField("message_str", msgDesc.String()).Trace("Converting message")
	for _, fieldDesc := range msgDesc.GetField() {

		// Custom field options from protoc-gen-jsonschema:
		if opt := proto.GetExtension(fieldDesc.GetOptions(), protoc_gen_jsonschema.E_FieldOptions); opt != nil {
			if fieldOptions, ok := opt.(*protoc_gen_jsonschema.FieldOptions); ok {

				// "Ignored" fields are simply skipped:
				if fieldOptions.GetIgnore() {
					c.logger.WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Debug("Skipping ignored field")
					continue
				}

				// "Required" fields are added to the list of required attributes in our schema:
				if fieldOptions.GetRequired() {
					c.logger.WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Debug("Marking required field")
					if c.Flags.UseJSONFieldnamesOnly {
						jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetJsonName())
					} else {
						jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetName())
					}
				}
			}
		}

		// Convert the field into a JSONSchema type:
		recursedJSONSchemaType, err := c.convertField(curPkg, fieldDesc, msgDesc, duplicatedMessages, messageFlags)
		if err != nil {
			c.logger.WithError(err).WithField("field_name", fieldDesc.GetName()).WithField("message_name", msgDesc.GetName()).Error("Failed to convert field")
			return nil, err
		}
		c.logger.WithField("field_name", fieldDesc.GetName()).WithField("type", recursedJSONSchemaType.Type).Trace("Converted field")

		// If this field is part of a OneOf declaration then build that here:
		if c.Flags.EnforceOneOf && fieldDesc.OneofIndex != nil && !fieldDesc.GetProto3Optional() {
			for {
				if *fieldDesc.OneofIndex < int32(len(jsonSchemaType.AllOf)) {
					break
				}
				var notAnyOf = &jsonschema.Schema{Not: &jsonschema.Schema{AnyOf: []*jsonschema.Schema{}}}
				jsonSchemaType.AllOf = append(jsonSchemaType.AllOf, &jsonschema.Schema{OneOf: []*jsonschema.Schema{notAnyOf}})
			}
			if c.Flags.UseJSONFieldnamesOnly {
				jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf = append(jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf, &jsonschema.Schema{Required: []string{fieldDesc.GetJsonName()}})
				jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf[0].Not.AnyOf = append(jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf[0].Not.AnyOf, &jsonschema.Schema{Required: []string{fieldDesc.GetJsonName()}})
			} else {
				jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf = append(jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf, &jsonschema.Schema{Required: []string{fieldDesc.GetName()}})
				jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf[0].Not.AnyOf = append(jsonSchemaType.AllOf[*fieldDesc.OneofIndex].OneOf[0].Not.AnyOf, &jsonschema.Schema{Required: []string{fieldDesc.GetName()}})
			}
		}

		// Figure out which field names we want to use:
		switch {
		case c.Flags.UseJSONFieldnamesOnly:
			jsonSchemaType.Properties.Set(fieldDesc.GetJsonName(), recursedJSONSchemaType)
		case c.Flags.UseProtoAndJSONFieldNames:
			jsonSchemaType.Properties.Set(fieldDesc.GetName(), recursedJSONSchemaType)
			jsonSchemaType.Properties.Set(fieldDesc.GetJsonName(), recursedJSONSchemaType)
		default:
			jsonSchemaType.Properties.Set(fieldDesc.GetName(), recursedJSONSchemaType)
		}

		// Enforce all_fields_required:
		if messageFlags.AllFieldsRequired {
			if fieldDesc.OneofIndex == nil && !fieldDesc.GetProto3Optional() && fieldDesc.GetLabel() != descriptor.FieldDescriptorProto_LABEL_REPEATED {

				if c.Flags.UseJSONFieldnamesOnly {
					jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetJsonName())
				} else {
					jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetName())
				}
			}
		}

		// Look for required fields by the proto2 "required" flag:
		if fieldDesc.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED && fieldDesc.OneofIndex == nil {
			if c.Flags.UseJSONFieldnamesOnly {
				jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetJsonName())
			} else {
				jsonSchemaType.Required = append(jsonSchemaType.Required, fieldDesc.GetName())
			}
		}
	}

	// Remove empty properties to keep the final output as clean as possible:
	if jsonSchemaType.Properties.Len() == 0 {
		jsonSchemaType.Properties = nil
	}

	// Dedupe required fields:
	jsonSchemaType.Required = dedupe(jsonSchemaType.Required)

	return jsonSchemaType, nil
}

func dedupe(inputStrings []string) []string {
	appended := make(map[string]bool)
	outputStrings := []string{}

	for _, inputString := range inputStrings {
		if !appended[inputString] {
			outputStrings = append(outputStrings, inputString)
			appended[inputString] = true
		}
	}
	return outputStrings
}

// ptrUint64 returns a pointer to a uint64 with the given value, for use with
// invopop/jsonschema fields like MaxLength / MaxItems which are typed *uint64.
// A zero value is treated as "unset" and returns nil so the field is omitted
// from the rendered schema rather than emitted as `0` — proto field options
// like min_length / max_length default to 0 meaning "no constraint", and
// emitting `maxLength: 0` would forbid any non-empty string.
func ptrUint64(v uint64) *uint64 {
	if v == 0 {
		return nil
	}
	return &v
}

// schemaAllowAny returns the canonical "matches anything" schema. invopop
// exports a package-level singleton whose MarshalJSON emits literal `true`.
// Use for additionalProperties when the converter wants to explicitly state
// "any additional properties allowed."
func schemaAllowAny() *jsonschema.Schema {
	return jsonschema.TrueSchema
}

// schemaAllowNone returns the canonical "matches nothing" schema. invopop
// exports a package-level singleton whose MarshalJSON emits literal `false`.
// Use for additionalProperties when the converter wants to disallow any
// additional properties.
func schemaAllowNone() *jsonschema.Schema {
	return jsonschema.FalseSchema
}

// numericBounds returns the inclusive (minimum, maximum) pair for a protobuf
// numeric scalar type. The bool indicates whether the type is numeric at all.
//
// Bound values are encoded as json.Number string literals so that uint64's
// 2^64-1 and float64's ±1.8e308 round-trip without lossy conversion through a
// Go int (which would overflow on 32-bit hosts) or a float64 (which can only
// approximate the 64-bit integer extremes).
func numericBounds(protoType descriptor.FieldDescriptorProto_Type) (min, max json.Number, ok bool) {
	switch protoType {
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "-3.4028234663852886e+38", "3.4028234663852886e+38", true
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "-1.7976931348623157e+308", "1.7976931348623157e+308", true
	case descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		return "-2147483648", "2147483647", true
	case descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED32:
		return "0", "4294967295", true
	case descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return "-9223372036854775808", "9223372036854775807", true
	case descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return "0", "18446744073709551615", true
	}
	return "", "", false
}

// applyNumericBounds attaches minimum/maximum to a schema produced for a
// protobuf numeric scalar. For nullable fields the schema is a OneOf of null +
// the typed branch, and bounds are applied to the typed branch. The string
// variant of int64-class fields (proto3's default JSON encoding) is skipped:
// JSON Schema numeric bounds apply only to integer/number-typed schemas.
func applyNumericBounds(schema *jsonschema.Schema, protoType descriptor.FieldDescriptorProto_Type) {
	min, max, ok := numericBounds(protoType)
	if !ok {
		return
	}
	if len(schema.OneOf) > 0 {
		for _, branch := range schema.OneOf {
			if branch.Type == jsTypeInteger || branch.Type == jsTypeNumber {
				branch.Minimum = min
				branch.Maximum = max
			}
		}
		return
	}
	if schema.Type == jsTypeInteger || schema.Type == jsTypeNumber {
		schema.Minimum = min
		schema.Maximum = max
	}
}
