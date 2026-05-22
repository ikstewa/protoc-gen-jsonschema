package testdata

const GoogleInt64ValueAllowNull = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/GoogleInt64ValueAllowNull",
    "$defs": {
        "GoogleInt64ValueAllowNull": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "big_number": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ],
                    "title": "Int 64 Value",
                    "description": "Wrapper message for ` + "`" + `int64` + "`" + `. The JSON representation for ` + "`" + `Int64Value` + "`" + ` is JSON string. Not recommended for use in new APIs, but still useful for legacy APIs and has no plan to be removed."
                }
            },
            "additionalProperties": true,
            "title": "Google Int 64 Value Allow Null"
        }
    }
}`

const GoogleInt64ValueAllowNullFail = `{"big_number": 12345}`

const GoogleInt64ValueAllowNullPass = `{"big_number": null}`
