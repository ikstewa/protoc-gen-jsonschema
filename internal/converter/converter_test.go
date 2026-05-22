package converter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"

	"github.com/ikstewa/protoc-gen-jsonschema/internal/converter/testdata"
)

const (
	sampleProtoDirectory = "testdata/proto"
)

type sampleProto struct {
	Flags                 ConverterFlags
	ExpectedFileNames     []string
	ExpectedJSONSchema    []string
	FilesToGenerate       []string
	ObjectsToValidateFail []string
	ObjectsToValidatePass []string
	ProtoFileName         string
	TargetedMessages      []string
}

func TestGenerateJsonSchema(t *testing.T) {

	// Configure the list of sample protos to test, and their expected JSON-Schemas:
	sampleProtos := configureSampleProtos()

	// Convert the protos, compare the results against the expected JSON-Schemas:
	for name, sampleProto := range sampleProtos {
		t.Run(name, func(t *testing.T) {
			testConvertSampleProto(t, sampleProto)
		})
	}
}

func testConvertSampleProto(t *testing.T, sampleProto sampleProto) {
	t.Helper()

	// Make a Logrus logger:
	logger := logrus.New()
	logger.SetLevel(logrus.ErrorLevel)
	logger.SetOutput(os.Stderr)

	// Use the logger to make a Converter:
	protoConverter := New(logger)
	protoConverter.Flags = sampleProto.Flags

	// Open the sample proto file:
	sampleProtoFileName := fmt.Sprintf("%v/%v", sampleProtoDirectory, sampleProto.ProtoFileName)
	fileDescriptorSet := mustReadProtoFiles(t, sampleProtoDirectory, sampleProto.ProtoFileName)

	// Prepare a request:
	codeGeneratorRequest := plugin.CodeGeneratorRequest{
		FileToGenerate: sampleProto.FilesToGenerate,
		ProtoFile:      fileDescriptorSet.GetFile(),
	}

	// Test the TargetedMessages feature:
	if len(sampleProto.TargetedMessages) > 0 {
		arg := fmt.Sprintf("messages=[%s]", strings.Join(sampleProto.TargetedMessages, messageDelimiter))
		codeGeneratorRequest.Parameter = &arg
	}

	// Perform the conversion:
	response, err := protoConverter.convert(&codeGeneratorRequest)
	assert.NoError(t, err, "Unable to convert sample proto file (%v)", sampleProtoFileName)
	assert.Equal(t, len(sampleProto.ExpectedJSONSchema), len(response.File), "Incorrect number of JSON-Schema files returned for sample proto file (%v)", sampleProtoFileName)
	if len(sampleProto.ExpectedJSONSchema) != len(response.File) {
		t.Fail()
	} else {
		for responseFileIndex, responseFile := range response.File {

			expected := strings.TrimSpace(sampleProto.ExpectedJSONSchema[responseFileIndex])
			actual := responseFile.GetContent()

			// UPDATE_GOLDENS=1 rewrites the matching testdata const in place
			// instead of asserting. See updateGoldenForExpected for details.
			if os.Getenv("UPDATE_GOLDENS") == "1" {
				if err := updateGoldenForExpected(t, expected, actual); err != nil {
					t.Fatalf("UPDATE_GOLDENS: %v (proto=%v idx=%d)", err, sampleProtoFileName, responseFileIndex)
				}
				continue
			}

			// Ensure that the generated schema matches the expected (canned) one:
			assert.Equal(t, expected, actual, "Incorrect JSON-Schema returned for sample proto file (%v)", sampleProtoFileName)

			// Validate the generated filenames:
			if len(sampleProto.ExpectedFileNames) > 0 {
				assert.Equal(t, sampleProto.ExpectedFileNames[responseFileIndex], responseFile.GetName())
			}

			// Validate any intended-to-fail data against the new schema:
			if len(sampleProto.ObjectsToValidateFail) >= responseFileIndex+1 {
				valid, err := validateSchema(*responseFile.Content, sampleProto.ObjectsToValidateFail[responseFileIndex])
				assert.NoError(t, err)
				assert.False(t, valid, "Expected canned data to fail validation)")
			}

			// Validate any intended-to-pass data against the new schema:
			if len(sampleProto.ObjectsToValidatePass) >= responseFileIndex+1 {
				valid, err := validateSchema(*responseFile.Content, sampleProto.ObjectsToValidatePass[responseFileIndex])
				assert.NoError(t, err, "Error validating canned data with generated schema")
				assert.True(t, valid, "Expected canned data validate)")
			}
		}
	}

	// Return now if we have no files:
	if len(response.File) == 0 {
		return
	}

	// Check for the correct prefix:
	if protoConverter.Flags.PrefixSchemaFilesWithPackage {
		assert.Contains(t, response.File[0].GetName(), "samples")
	} else {
		assert.NotContains(t, response.File[0].GetName(), "samples")
	}
}

func configureSampleProtos() map[string]sampleProto {
	return map[string]sampleProto{
		"AllRequired": {
			Flags:                 ConverterFlags{AllFieldsRequired: true},
			ExpectedJSONSchema:    []string{testdata.PayloadMessage2},
			FilesToGenerate:       []string{"PayloadMessage2.proto"},
			ProtoFileName:         "PayloadMessage2.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessage2Fail},
			ObjectsToValidatePass: []string{testdata.PayloadMessage2Pass},
		},
		"ArrayOfEnums": {
			ExpectedJSONSchema:    []string{testdata.ArrayOfEnums},
			FilesToGenerate:       []string{"ArrayOfEnums.proto"},
			ProtoFileName:         "ArrayOfEnums.proto",
			ObjectsToValidateFail: []string{testdata.ArrayOfEnumsFail},
			ObjectsToValidatePass: []string{testdata.ArrayOfEnumsPass},
		},
		"ArrayOfMessages": {
			ExpectedJSONSchema:    []string{testdata.PayloadMessage, testdata.ArrayOfMessages},
			FilesToGenerate:       []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
			ProtoFileName:         "ArrayOfMessages.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessageFail, testdata.ArrayOfMessagesFail},
			ObjectsToValidatePass: []string{testdata.PayloadMessagePass, testdata.ArrayOfMessagesPass},
		},
		"TypeNamesWithNoPackage": {
			Flags:                 ConverterFlags{TypeNamesWithNoPackage: true},
			ExpectedJSONSchema:    []string{testdata.PayloadMessage, testdata.TypeNamesWithNoPackage},
			FilesToGenerate:       []string{"ArrayOfMessages.proto", "PayloadMessage.proto"},
			ProtoFileName:         "ArrayOfMessages.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessageFail, testdata.TypeNamesWithNoPackageFail},
			ObjectsToValidatePass: []string{testdata.PayloadMessagePass, testdata.TypeNamesWithNoPackagePass},
		},
		"ArrayOfObjects": {
			Flags:                 ConverterFlags{AllowNullValues: true},
			ExpectedJSONSchema:    []string{testdata.ArrayOfObjects},
			FilesToGenerate:       []string{"ArrayOfObjects.proto"},
			ProtoFileName:         "ArrayOfObjects.proto",
			ObjectsToValidateFail: []string{testdata.ArrayOfObjectsFail},
			ObjectsToValidatePass: []string{testdata.ArrayOfObjectsPass},
		},
		"ArrayOfPrimitives": {
			Flags:                 ConverterFlags{AllowNullValues: true},
			ExpectedJSONSchema:    []string{testdata.ArrayOfPrimitives},
			FilesToGenerate:       []string{"ArrayOfPrimitives.proto"},
			ProtoFileName:         "ArrayOfPrimitives.proto",
			ObjectsToValidateFail: []string{testdata.ArrayOfPrimitivesFail},
			ObjectsToValidatePass: []string{testdata.ArrayOfPrimitivesPass},
		},
		"ArrayOfPrimitivesDouble": {
			Flags: ConverterFlags{
				AllowNullValues:           true,
				UseProtoAndJSONFieldNames: true,
			},
			ExpectedJSONSchema:    []string{testdata.ArrayOfPrimitivesDouble},
			FilesToGenerate:       []string{"ArrayOfPrimitives.proto"},
			ProtoFileName:         "ArrayOfPrimitives.proto",
			ObjectsToValidateFail: []string{testdata.ArrayOfPrimitivesDoubleFail},
			ObjectsToValidatePass: []string{testdata.ArrayOfPrimitivesDoublePass},
		},
		"BigIntAsString": {
			Flags: ConverterFlags{
				AllowNullValues:          true,
				DisallowBigIntsAsStrings: true,
			},
			ExpectedJSONSchema:    []string{testdata.BigIntAsString},
			FilesToGenerate:       []string{"BigIntAsString.proto"},
			ProtoFileName:         "BigIntAsString.proto",
			ObjectsToValidateFail: []string{testdata.BigIntAsStringFail},
			ObjectsToValidatePass: []string{testdata.BigIntAsStringPass},
		},
		"BytesPayload": {
			ExpectedJSONSchema:    []string{testdata.BytesPayload},
			FilesToGenerate:       []string{"BytesPayload.proto"},
			ProtoFileName:         "BytesPayload.proto",
			ObjectsToValidateFail: []string{testdata.BytesPayloadFail},
		},
		"Comments": {
			ExpectedJSONSchema:    []string{testdata.MessageWithComments},
			FilesToGenerate:       []string{"MessageWithComments.proto"},
			ProtoFileName:         "MessageWithComments.proto",
			ObjectsToValidateFail: []string{testdata.MessageWithCommentsFail},
		},
		"CyclicalReference": {
			ExpectedJSONSchema: []string{testdata.CyclicalReferenceMessageM, testdata.CyclicalReferenceMessageFoo, testdata.CyclicalReferenceMessageBar, testdata.CyclicalReferenceMessageBaz},
			FilesToGenerate:    []string{"CyclicalReference.proto"},
			ProtoFileName:      "CyclicalReference.proto",
		},
		"EnumNestedReference": {
			ExpectedJSONSchema:    []string{testdata.EnumNestedReference},
			FilesToGenerate:       []string{"EnumNestedReference.proto"},
			ProtoFileName:         "EnumNestedReference.proto",
			ObjectsToValidateFail: []string{testdata.EnumNestedReferenceFail},
			ObjectsToValidatePass: []string{testdata.EnumNestedReferencePass},
		},
		"EnumWithMessage": {
			ExpectedJSONSchema:    []string{testdata.EnumWithMessage},
			FilesToGenerate:       []string{"EnumWithMessage.proto"},
			ProtoFileName:         "EnumWithMessage.proto",
			ObjectsToValidateFail: []string{testdata.EnumWithMessageFail},
			ObjectsToValidatePass: []string{testdata.EnumWithMessagePass},
		},
		"EnumCeption": {
			ExpectedJSONSchema:    []string{testdata.PayloadMessage, testdata.ImportedEnum, testdata.EnumCeption},
			FilesToGenerate:       []string{"Enumception.proto", "PayloadMessage.proto", "ImportedEnum.proto"},
			ProtoFileName:         "Enumception.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessageFail, testdata.ImportedEnumFail, testdata.EnumCeptionFail},
			ObjectsToValidatePass: []string{testdata.PayloadMessagePass, testdata.ImportedEnumPass, testdata.EnumCeptionPass},
		},
		"GoogleValue": {
			Flags:                 ConverterFlags{DisallowAdditionalProperties: true},
			ExpectedJSONSchema:    []string{testdata.GoogleValue},
			FilesToGenerate:       []string{"GoogleValue.proto"},
			ProtoFileName:         "GoogleValue.proto",
			ObjectsToValidateFail: []string{testdata.GoogleValueFail},
			ObjectsToValidatePass: []string{testdata.GoogleValuePass},
		},
		"GoogleInt64Value": {
			ExpectedJSONSchema:    []string{testdata.GoogleInt64Value},
			FilesToGenerate:       []string{"GoogleInt64Value.proto"},
			ProtoFileName:         "GoogleInt64Value.proto",
			ObjectsToValidateFail: []string{testdata.GoogleInt64ValueFail},
			ObjectsToValidatePass: []string{testdata.GoogleInt64ValuePass},
		},
		"GoogleInt64ValueAllowNull": {
			Flags:                 ConverterFlags{AllowNullValues: true},
			ExpectedJSONSchema:    []string{testdata.GoogleInt64ValueAllowNull},
			FilesToGenerate:       []string{"GoogleInt64ValueAllowNull.proto"},
			ProtoFileName:         "GoogleInt64ValueAllowNull.proto",
			ObjectsToValidateFail: []string{testdata.GoogleInt64ValueAllowNullFail},
			ObjectsToValidatePass: []string{testdata.GoogleInt64ValueAllowNullPass},
		},
		"GoogleInt64ValueDisallowString": {
			Flags:                 ConverterFlags{DisallowBigIntsAsStrings: true},
			ExpectedJSONSchema:    []string{testdata.GoogleInt64ValueDisallowString},
			FilesToGenerate:       []string{"GoogleInt64ValueDisallowString.proto"},
			ProtoFileName:         "GoogleInt64ValueDisallowString.proto",
			ObjectsToValidateFail: []string{testdata.GoogleInt64ValueDisallowStringFail},
			ObjectsToValidatePass: []string{testdata.GoogleInt64ValueDisallowStringPass},
		},
		"GoogleInt64ValueDisallowStringAllowNull": {
			Flags: ConverterFlags{
				DisallowBigIntsAsStrings: true,
				AllowNullValues:          true,
			},
			ExpectedJSONSchema:    []string{testdata.GoogleInt64ValueDisallowStringAllowNull},
			FilesToGenerate:       []string{"GoogleInt64ValueDisallowStringAllowNull.proto"},
			ProtoFileName:         "GoogleInt64ValueDisallowStringAllowNull.proto",
			ObjectsToValidateFail: []string{testdata.GoogleInt64ValueDisallowStringAllowNullFail},
			ObjectsToValidatePass: []string{testdata.GoogleInt64ValueDisallowStringAllowNullPass},
		},
		"ImportedEnum": {
			ExpectedJSONSchema:    []string{testdata.ImportedEnum},
			FilesToGenerate:       []string{"ImportedEnum.proto"},
			ProtoFileName:         "ImportedEnum.proto",
			ObjectsToValidateFail: []string{testdata.ImportedEnumFail},
			ObjectsToValidatePass: []string{testdata.ImportedEnumPass},
		},
		"JSONFields": {
			Flags:                 ConverterFlags{UseJSONFieldnamesOnly: true},
			ExpectedJSONSchema:    []string{testdata.JSONFields},
			FilesToGenerate:       []string{"JSONFields.proto"},
			ProtoFileName:         "JSONFields.proto",
			ObjectsToValidateFail: []string{testdata.JSONFieldsFail},
			ObjectsToValidatePass: []string{testdata.JSONFieldsPass},
		},
		"Maps": {
			ExpectedJSONSchema:    []string{testdata.Maps},
			FilesToGenerate:       []string{"Maps.proto"},
			ProtoFileName:         "Maps.proto",
			ObjectsToValidateFail: []string{testdata.MapsFail},
			ObjectsToValidatePass: []string{testdata.MapsPass},
		},
		"NestedMessage": {
			ExpectedJSONSchema:    []string{testdata.PayloadMessage, testdata.NestedMessage},
			FilesToGenerate:       []string{"NestedMessage.proto", "PayloadMessage.proto"},
			ProtoFileName:         "NestedMessage.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessageFail, testdata.NestedMessageFail},
			ObjectsToValidatePass: []string{testdata.PayloadMessagePass, testdata.NestedMessagePass},
		},
		"NestedObject": {
			ExpectedJSONSchema:    []string{testdata.NestedObject},
			FilesToGenerate:       []string{"NestedObject.proto"},
			ProtoFileName:         "NestedObject.proto",
			ObjectsToValidateFail: []string{testdata.NestedObjectFail},
			ObjectsToValidatePass: []string{testdata.NestedObjectPass},
		},
		"NoPackage": {
			ExpectedJSONSchema: []string{},
			FilesToGenerate:    []string{},
			ProtoFileName:      "NoPackage.proto",
		},
		"NumericFormat": {
			Flags:                 ConverterFlags{IncludeNumericFormat: true},
			ExpectedJSONSchema:    []string{testdata.NumericFormat},
			FilesToGenerate:       []string{"NumericFormat.proto"},
			ProtoFileName:         "NumericFormat.proto",
		},
		"NumericFormatBigIntAsString": {
			Flags:                 ConverterFlags{IncludeNumericFormat: true, DisallowBigIntsAsStrings: true},
			ExpectedJSONSchema:    []string{testdata.NumericFormatBigIntAsString},
			FilesToGenerate:       []string{"NumericFormat.proto"},
			ProtoFileName:         "NumericFormat.proto",
		},
		"OneOf": {
			Flags:                 ConverterFlags{AllFieldsRequired: true, EnforceOneOf: true},
			ExpectedJSONSchema:    []string{testdata.OneOf},
			FilesToGenerate:       []string{"OneOf.proto"},
			ProtoFileName:         "OneOf.proto",
			ObjectsToValidateFail: []string{testdata.OneOfFail},
			ObjectsToValidatePass: []string{testdata.OneOfPass},
		},
		"OptionAllowNullValues": {
			ExpectedJSONSchema:    []string{testdata.OptionAllowNullValues},
			FilesToGenerate:       []string{"OptionAllowNullValues.proto"},
			ProtoFileName:         "OptionAllowNullValues.proto",
			ObjectsToValidateFail: []string{testdata.OptionAllowNullValuesFail},
			ObjectsToValidatePass: []string{testdata.OptionAllowNullValuesPass},
		},
		"OptionDisallowAdditionalProperties": {
			ExpectedJSONSchema:    []string{testdata.OptionDisallowAdditionalProperties},
			FilesToGenerate:       []string{"OptionDisallowAdditionalProperties.proto"},
			ProtoFileName:         "OptionDisallowAdditionalProperties.proto",
			ObjectsToValidateFail: []string{testdata.OptionDisallowAdditionalPropertiesFail},
			ObjectsToValidatePass: []string{testdata.OptionDisallowAdditionalPropertiesPass},
		},
		"OptionEnumsAsConstants": {
			ExpectedJSONSchema:    []string{testdata.OptionEnumsAsConstants},
			FilesToGenerate:       []string{"OptionEnumsAsConstants.proto"},
			ProtoFileName:         "OptionEnumsAsConstants.proto",
			ObjectsToValidateFail: []string{testdata.OptionEnumsAsConstantsFail},
			ObjectsToValidatePass: []string{testdata.OptionEnumsAsConstantsPass},
		},
		"OptionEnumsAsStringsOnly": {
			Flags: ConverterFlags{
				EnumsAsStringsOnly: true,
			},
			ExpectedJSONSchema:    []string{testdata.OptionEnumsAsStringsOnly},
			FilesToGenerate:       []string{"OptionEnumsAsStringsOnly.proto"},
			ProtoFileName:         "OptionEnumsAsStringsOnly.proto",
			ObjectsToValidateFail: []string{testdata.OptionEnumsAsStringsOnlyFail},
			ObjectsToValidatePass: []string{testdata.OptionEnumsAsStringsOnlyPass},
		},
		"OptionEnumsTrimPrefix": {
			Flags: ConverterFlags{
				EnumsTrimPrefix: true,
			},
			ExpectedJSONSchema:    []string{testdata.OptionEnumsTrimPrefix},
			FilesToGenerate:       []string{"OptionEnumsTrimPrefix.proto"},
			ProtoFileName:         "OptionEnumsTrimPrefix.proto",
			ObjectsToValidateFail: []string{testdata.OptionEnumsTrimPrefixFail},
			ObjectsToValidatePass: []string{testdata.OptionEnumsTrimPrefixPass},
		},
		"OptionFileExtension": {
			ExpectedJSONSchema: []string{testdata.OptionFileExtension},
			ExpectedFileNames:  []string{"OptionFileExtension.jsonschema"},
			FilesToGenerate:    []string{"OptionFileExtension.proto"},
			ProtoFileName:      "OptionFileExtension.proto",
		},
		"OptionIgnoredEnum": {
			ExpectedJSONSchema: []string{testdata.UnignoredEnum},
			FilesToGenerate:    []string{"OptionIgnoredEnum.proto"},
			ProtoFileName:      "OptionIgnoredEnum.proto",
		},
		"OptionIgnoredFile": {
			ExpectedJSONSchema: []string{},
			FilesToGenerate:    []string{"OptionIgnoredFile.proto"},
			ProtoFileName:      "OptionIgnoredFile.proto",
		},
		"OptionIgnoredField": {
			ExpectedJSONSchema:    []string{testdata.OptionIgnoredField},
			FilesToGenerate:       []string{"OptionIgnoredField.proto"},
			ProtoFileName:         "OptionIgnoredField.proto",
			ObjectsToValidateFail: []string{testdata.OptionIgnoredFieldFail},
			ObjectsToValidatePass: []string{testdata.OptionIgnoredFieldPass},
		},
		"OptionIgnoredMessage": {
			ExpectedJSONSchema: []string{testdata.UnignoredMessage},
			FilesToGenerate:    []string{"OptionIgnoredMessage.proto"},
			ProtoFileName:      "OptionIgnoredMessage.proto",
		},
		"OptionRequiredField": {
			ExpectedJSONSchema:    []string{testdata.OptionRequiredField},
			FilesToGenerate:       []string{"OptionRequiredField.proto"},
			ProtoFileName:         "OptionRequiredField.proto",
			ObjectsToValidateFail: []string{testdata.OptionRequiredFieldFail},
			ObjectsToValidatePass: []string{testdata.OptionRequiredFieldPass},
		},
		"OptionMinLength": {
			ExpectedJSONSchema:    []string{testdata.OptionMinLength},
			FilesToGenerate:       []string{"OptionMinLength.proto"},
			ProtoFileName:         "OptionMinLength.proto",
			ObjectsToValidateFail: []string{testdata.OptionMinLengthFail},
			ObjectsToValidatePass: []string{testdata.OptionMinLengthPass},
		},
		"OptionMaxLength": {
			ExpectedJSONSchema:    []string{testdata.OptionMaxLength},
			FilesToGenerate:       []string{"OptionMaxLength.proto"},
			ProtoFileName:         "OptionMaxLength.proto",
			ObjectsToValidateFail: []string{testdata.OptionMaxLengthFail},
			ObjectsToValidatePass: []string{testdata.OptionMaxLengthPass},
		},
		"OptionPattern": {
			ExpectedJSONSchema:    []string{testdata.OptionPattern},
			FilesToGenerate:       []string{"OptionPattern.proto"},
			ProtoFileName:         "OptionPattern.proto",
			ObjectsToValidateFail: []string{testdata.OptionPatternFail},
			ObjectsToValidatePass: []string{testdata.OptionPatternPass},
		},
		"OptionRequiredMessage": {
			ExpectedJSONSchema:    []string{testdata.OptionRequiredMessage},
			FilesToGenerate:       []string{"OptionRequiredMessage.proto"},
			ProtoFileName:         "OptionRequiredMessage.proto",
			ObjectsToValidateFail: []string{testdata.OptionRequiredMessageFail},
			ObjectsToValidatePass: []string{testdata.OptionRequiredMessagePass},
		},
		"PackagePrefix": {
			Flags:                 ConverterFlags{PrefixSchemaFilesWithPackage: true},
			ExpectedJSONSchema:    []string{testdata.Timestamp},
			FilesToGenerate:       []string{"Timestamp.proto"},
			ProtoFileName:         "Timestamp.proto",
			ObjectsToValidateFail: []string{testdata.TimestampFail},
			ObjectsToValidatePass: []string{testdata.TimestampPass},
		},
		"PayloadMessage": {
			ExpectedJSONSchema:    []string{testdata.PayloadMessage},
			FilesToGenerate:       []string{"PayloadMessage.proto"},
			ProtoFileName:         "PayloadMessage.proto",
			ObjectsToValidateFail: []string{testdata.PayloadMessageFail},
			ObjectsToValidatePass: []string{testdata.PayloadMessagePass},
		},
		"Proto2NestedMessage": {
			ExpectedJSONSchema:    []string{testdata.Proto2PayloadMessage, testdata.Proto2NestedMessage},
			FilesToGenerate:       []string{"Proto2PayloadMessage.proto", "Proto2NestedMessage.proto"},
			ProtoFileName:         "Proto2NestedMessage.proto",
			ObjectsToValidateFail: []string{testdata.Proto2PayloadMessageFail, testdata.Proto2NestedMessageFail},
			ObjectsToValidatePass: []string{testdata.Proto2PayloadMessagePass, testdata.Proto2NestedMessagePass},
		},
		"Proto2NestedObject": {
			Flags:                 ConverterFlags{AllFieldsRequired: true},
			ExpectedJSONSchema:    []string{testdata.Proto2NestedObject},
			FilesToGenerate:       []string{"Proto2NestedObject.proto"},
			ProtoFileName:         "Proto2NestedObject.proto",
			ObjectsToValidateFail: []string{testdata.Proto2NestedObjectFail},
			ObjectsToValidatePass: []string{testdata.Proto2NestedObjectPass},
		},
		"Proto2Required": {
			ExpectedJSONSchema:    []string{testdata.Proto2Required},
			FilesToGenerate:       []string{"Proto2Required.proto"},
			ProtoFileName:         "Proto2Required.proto",
			ObjectsToValidateFail: []string{testdata.Proto2RequiredFail},
			ObjectsToValidatePass: []string{testdata.Proto2RequiredPass},
		},
		"SelfReference": {
			ExpectedJSONSchema:    []string{testdata.SelfReference},
			FilesToGenerate:       []string{"SelfReference.proto"},
			ProtoFileName:         "SelfReference.proto",
			ObjectsToValidateFail: []string{testdata.SelfReferenceFail},
			ObjectsToValidatePass: []string{testdata.SelfReferencePass},
		},
		"SeveralEnums": {
			ExpectedJSONSchema:    []string{testdata.FirstEnum, testdata.SecondEnum},
			FilesToGenerate:       []string{"SeveralEnums.proto"},
			ProtoFileName:         "SeveralEnums.proto",
			ObjectsToValidateFail: []string{testdata.FirstEnumFail, testdata.SecondEnumFail},
			ObjectsToValidatePass: []string{testdata.FirstEnumPass, testdata.SecondEnumPass},
		},
		"SeveralMessages": {
			ExpectedJSONSchema:    []string{testdata.FirstMessage, testdata.SecondMessage},
			FilesToGenerate:       []string{"SeveralMessages.proto"},
			ProtoFileName:         "SeveralMessages.proto",
			ObjectsToValidateFail: []string{testdata.FirstMessageFail, testdata.SecondMessageFail},
			ObjectsToValidatePass: []string{testdata.FirstMessagePass, testdata.SecondMessagePass},
		},
		"TargetedMessages": {
			TargetedMessages:   []string{"MessageKind10", "MessageKind11", "MessageKind12"},
			ExpectedJSONSchema: []string{testdata.MessageKind10, testdata.MessageKind11, testdata.MessageKind12},
			FilesToGenerate:    []string{"TwelveMessages.proto"},
			ProtoFileName:      "TwelveMessages.proto",
		},
		"Timestamp": {
			ExpectedJSONSchema:    []string{testdata.Timestamp},
			FilesToGenerate:       []string{"Timestamp.proto"},
			ProtoFileName:         "Timestamp.proto",
			ObjectsToValidateFail: []string{testdata.TimestampFail},
			ObjectsToValidatePass: []string{testdata.TimestampPass},
		},
		"ValidationOptions": {
			ExpectedJSONSchema:    []string{testdata.ValidationOptions},
			FilesToGenerate:       []string{"ValidationOptions.proto"},
			ProtoFileName:         "ValidationOptions.proto",
			ObjectsToValidateFail: []string{testdata.ValidationOptionsFail},
			ObjectsToValidatePass: []string{testdata.ValidationOptionsPass},
		},
		"WellKnown": {
			ExpectedJSONSchema:    []string{testdata.WellKnown},
			FilesToGenerate:       []string{"WellKnown.proto"},
			ProtoFileName:         "WellKnown.proto",
			ObjectsToValidateFail: []string{testdata.WellKnownFail},
			ObjectsToValidatePass: []string{testdata.WellKnownPass},
		},
	}
}

// Load the specified .proto files into a FileDescriptorSet. Any errors in loading/parsing will
// immediately fail the test.
func mustReadProtoFiles(t *testing.T, includePath string, filenames ...string) *descriptor.FileDescriptorSet {
	protocBinary, err := exec.LookPath("protoc")
	if err != nil {
		t.Fatalf("Can't find 'protoc' binary in $PATH: %s", err.Error())
	}

	// Use protoc to output descriptor info for the specified .proto files.
	var args []string
	args = append(args, "--descriptor_set_out=/dev/stdout")
	args = append(args, "--include_source_info")
	args = append(args, "--include_imports")
	args = append(args, "-I../../")
	args = append(args, "--proto_path="+includePath)
	args = append(args, filenames...)
	cmd := exec.Command(protocBinary, args...)
	stdoutBuf := bytes.Buffer{}
	stderrBuf := bytes.Buffer{}
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err = cmd.Run()
	if err != nil {
		t.Fatalf("failed to load descriptor set (%s): %s: %s",
			strings.Join(cmd.Args, " "), err.Error(), stderrBuf.String())
	}
	fds := &descriptor.FileDescriptorSet{}
	err = proto.Unmarshal(stdoutBuf.Bytes(), fds)
	if err != nil {
		t.Fatalf("failed to parse protoc output as FileDescriptorSet: %s", err.Error())
	}
	return fds
}

// validateSchema validates a JSON document against a JSON Schema string.
// Uses github.com/santhosh-tekuri/jsonschema/v6 which supports JSON Schema
// Draft 2020-12 (the default emitted by invopop/jsonschema). The previous
// implementation used github.com/xeipuuv/gojsonschema, which is capped at
// Draft 7 and cannot parse 2020-12 keywords like $defs.
func validateSchema(jsonSchema, jsonData string) (bool, error) {
	// Parse the schema document into an `any` tree (the form AddResource expects):
	var schemaDoc any
	if err := json.Unmarshal([]byte(jsonSchema), &schemaDoc); err != nil {
		return false, fmt.Errorf("validateSchema: failed to parse schema JSON: %w", err)
	}

	compiler := jsonschema.NewCompiler()
	// gojsonschema enforced `format` by default; santhosh-tekuri follows the
	// 2020-12 spec where format is annotation-only unless explicitly asserted.
	// The test goldens depend on format-as-assertion (e.g. TimestampFail
	// expects "twelve oclock" to be rejected as not date-time), so enable it.
	compiler.AssertFormat()
	if err := compiler.AddResource("schema.json", schemaDoc); err != nil {
		return false, fmt.Errorf("validateSchema: AddResource failed: %w", err)
	}
	sch, err := compiler.Compile("schema.json")
	if err != nil {
		return false, fmt.Errorf("validateSchema: Compile failed: %w", err)
	}

	// Parse the instance document:
	var instance any
	if err := json.Unmarshal([]byte(jsonData), &instance); err != nil {
		return false, fmt.Errorf("validateSchema: failed to parse instance JSON: %w", err)
	}

	if err := sch.Validate(instance); err != nil {
		// Validation failed (schema is well-formed but instance doesn't match).
		// Match the gojsonschema contract: return (false, nil) for instance
		// non-conformance rather than (_, err).
		return false, nil
	}
	return true, nil
}

// ---------------------------------------------------------------------------
// UPDATE_GOLDENS support
// ---------------------------------------------------------------------------
//
// When the env var UPDATE_GOLDENS=1 is set, the test harness rewrites the
// matching testdata const in place instead of asserting equality. Because
// the test feeds canned expected JSON via the `testdata.<Name>` constants
// (and ExpectedJSONSchema is just a []string), we recover the source const
// name by content-matching against an index built from every *.go file in
// testdata/. This sidesteps having to thread const names through the
// sampleProto table.

type goldenLocation struct {
	filePath  string
	constName string
}

var (
	goldenIndexOnce sync.Once
	goldenIndex     map[string]goldenLocation // key = strings.TrimSpace(constBody)
	goldenIndexErr  error
)

// loadGoldenIndex parses each testdata/*.go via Go's AST and indexes every
// top-level string const by its (trimmed) string value. The AST is necessary
// here because some golden consts include literal backticks (e.g. protoc-
// emitted descriptions referencing Go-style `int64`) which can't be encoded
// as raw-string literals — those declarations use concatenation form like
// `parta` + "`" + `partb`, which a naive regex would mis-parse.
func loadGoldenIndex() (map[string]goldenLocation, error) {
	goldenIndexOnce.Do(func() {
		idx := map[string]goldenLocation{}
		entries, err := os.ReadDir("testdata")
		if err != nil {
			goldenIndexErr = fmt.Errorf("reading testdata dir: %w", err)
			return
		}
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
				continue
			}
			path := filepath.Join("testdata", e.Name())
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, path, nil, 0)
			if err != nil {
				goldenIndexErr = fmt.Errorf("parsing %s: %w", path, err)
				return
			}
			for _, decl := range file.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.CONST {
					continue
				}
				for _, spec := range gen.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for i, name := range vs.Names {
						if i >= len(vs.Values) {
							continue
						}
						val, ok := evalConstString(vs.Values[i])
						if !ok {
							continue
						}
						key := strings.TrimSpace(val)
						idx[key] = goldenLocation{filePath: path, constName: name.Name}
					}
				}
			}
		}
		goldenIndex = idx
	})
	return goldenIndex, goldenIndexErr
}

// updateGoldenForExpected looks up the const whose current value matches
// `expected` and rewrites its body to `actual` in the corresponding
// testdata/*.go file. Uses AST byte offsets to find and replace just the
// targeted const's value expression — preserves sibling consts (e.g.
// <Name>Fail / <Name>Pass) in the same file and works even when the
// existing value expression is a concatenation of raw and double-quoted
// segments (as required for goldens that contain literal backticks).
func updateGoldenForExpected(t *testing.T, expected, actual string) error {
	t.Helper()

	idx, err := loadGoldenIndex()
	if err != nil {
		return err
	}
	loc, ok := idx[expected]
	if !ok {
		return fmt.Errorf("no testdata const matched the expected JSON (len=%d); add the const first or check for drift", len(expected))
	}

	src, err := os.ReadFile(loc.filePath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", loc.filePath, err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, loc.filePath, src, 0)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", loc.filePath, err)
	}

	var valueExpr ast.Expr
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		for _, spec := range gen.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for i, name := range vs.Names {
				if name.Name == loc.constName && i < len(vs.Values) {
					valueExpr = vs.Values[i]
				}
			}
		}
	}
	if valueExpr == nil {
		return fmt.Errorf("AST: const %s not found in %s", loc.constName, loc.filePath)
	}

	startOffset := fset.Position(valueExpr.Pos()).Offset
	endOffset := fset.Position(valueExpr.End()).Offset

	literal := goldenLiteral(actual)

	var out bytes.Buffer
	out.Write(src[:startOffset])
	out.WriteString(literal)
	out.Write(src[endOffset:])

	if bytes.Equal(out.Bytes(), src) {
		return nil
	}
	if err := os.WriteFile(loc.filePath, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", loc.filePath, err)
	}

	t.Logf("UPDATE_GOLDENS: rewrote const %s in %s", loc.constName, loc.filePath)
	return nil
}

// goldenLiteral encodes `s` as a Go string expression suitable for writing
// into a `const X = ...` declaration. Uses a raw-string literal for the
// common case; falls back to concatenation of raw-string and double-quoted
// segments when `s` contains a backtick (which can't appear inside a
// raw-string). The concatenation form looks like:
//
//	`parta` + "`" + `partb`
//
// preserving multi-line JSON readability for the common case.
func goldenLiteral(s string) string {
	if !strings.ContainsRune(s, '`') {
		return "`" + s + "`"
	}
	parts := strings.Split(s, "`")
	sep := "` + \"`\" + `"
	return "`" + strings.Join(parts, sep) + "`"
}

// evalConstString evaluates a Go AST expression and returns its string value
// if the expression is a string literal (raw or double-quoted) or a string
// concatenation of such literals (e.g. `a` + "`" + `b`). Returns (s, true)
// on success; (_, false) when the expression isn't a constant string.
func evalConstString(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind != token.STRING {
			return "", false
		}
		s, err := strconv.Unquote(e.Value)
		if err != nil {
			return "", false
		}
		return s, true
	case *ast.BinaryExpr:
		if e.Op != token.ADD {
			return "", false
		}
		lhs, lok := evalConstString(e.X)
		if !lok {
			return "", false
		}
		rhs, rok := evalConstString(e.Y)
		if !rok {
			return "", false
		}
		return lhs + rhs, true
	}
	return "", false
}
