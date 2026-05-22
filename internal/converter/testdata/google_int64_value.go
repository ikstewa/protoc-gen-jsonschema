package testdata

const GoogleInt64Value = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/GoogleInt64Value",
    "$defs": {
        "GoogleInt64Value": {
            "properties": {
                "big_number": {
                    "additionalProperties": true,
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Google Int 64 Value"
        }
    }
}`

const GoogleInt64ValueFail = `{"big_number": 12345}`

const GoogleInt64ValuePass = `{"big_number": "12345"}`
