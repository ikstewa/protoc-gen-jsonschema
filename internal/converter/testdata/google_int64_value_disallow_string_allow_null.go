package testdata

const GoogleInt64ValueDisallowStringAllowNull = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/GoogleInt64ValueDisallowStringAllowNull",
    "$defs": {
        "GoogleInt64ValueDisallowStringAllowNull": {
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
                            "type": "integer"
                        }
                    ],
                    "title": "Int 64 Value",
                    "description": "Wrapper message for ` + "`" + `int64` + "`" + `. The JSON representation for ` + "`" + `Int64Value` + "`" + ` is JSON string. Not recommended for use in new APIs, but still useful for legacy APIs and has no plan to be removed."
                }
            },
            "additionalProperties": true,
            "title": "Google Int 64 Value Disallow String Allow Null"
        }
    }
}`

const GoogleInt64ValueDisallowStringAllowNullFail = `{"big_number": "12345"}`

const GoogleInt64ValueDisallowStringAllowNullPass = `{"big_number": null}`
