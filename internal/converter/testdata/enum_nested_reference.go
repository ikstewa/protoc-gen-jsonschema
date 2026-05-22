package testdata

const EnumNestedReference = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Msg",
    "$defs": {
        "Msg": {
            "properties": {
                "nestedEnumField": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ],
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Msg"
        }
    }
}`

const EnumNestedReferenceFail = `{"nestedEnumField": 8}`

const EnumNestedReferencePass = `{"nestedEnumField": "FLAT"}`
