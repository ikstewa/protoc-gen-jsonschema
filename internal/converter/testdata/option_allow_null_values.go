package testdata

const OptionAllowNullValues = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionAllowNullValues",
    "$defs": {
        "OptionAllowNullValues": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "name2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "id2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                },
                "complete2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "boolean"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "title": "Option Allow Null Values"
        }
    }
}`

const OptionAllowNullValuesFail = `{"name2": 12345}`

const OptionAllowNullValuesPass = `{"name2": null}`
