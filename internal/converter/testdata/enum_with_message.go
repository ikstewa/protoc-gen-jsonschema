package testdata

const EnumWithMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/WithFooBarBaz",
    "$defs": {
        "WithFooBarBaz": {
            "properties": {
                "enumField": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ],
                    "enum": [
                        "Foo",
                        0,
                        "Bar",
                        1,
                        "Baz",
                        2
                    ],
                    "title": "Foo Bar Baz"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "With Foo Bar Baz"
        }
    }
}`

const EnumWithMessageFail = `{"enumField": 4}`

const EnumWithMessagePass = `{"enumField": 2}`
