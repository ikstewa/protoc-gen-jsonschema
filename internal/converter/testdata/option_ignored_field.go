package testdata

const OptionIgnoredField = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionIgnoredField",
    "$defs": {
        "OptionIgnoredField": {
            "properties": {
                "visible1": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "visible2": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Option Ignored Field"
        }
    }
}`

const OptionIgnoredFieldFail = `{"visible1": 12345}`

const OptionIgnoredFieldPass = `{"visible2": "hello"}`
