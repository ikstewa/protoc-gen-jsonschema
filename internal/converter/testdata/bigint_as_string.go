package testdata

const BigIntAsString = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/BigIntAsString",
    "$defs": {
        "BigIntAsString": {
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
                            "type": "integer"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "title": "Big Int As String"
        }
    }
}`

const BigIntAsStringFail = `{"big_number": "1827634182736443333"}`

const BigIntAsStringPass = `{"big_number": 1827634182736443333}`
