package testdata

const OptionDisallowAdditionalProperties = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionDisallowAdditionalProperties",
    "$defs": {
        "OptionDisallowAdditionalProperties": {
            "properties": {
                "name2": {
                    "type": "string"
                },
                "timestamp2": {
                    "type": "string"
                },
                "id2": {
                    "type": "integer"
                },
                "rating2": {
                    "type": "number"
                },
                "complete2": {
                    "type": "boolean"
                }
            },
            "additionalProperties": {
                "not": true
            },
            "type": "object",
            "title": "Option Disallow Additional Properties"
        }
    }
}`

const OptionDisallowAdditionalPropertiesFail = `{"something": 12345}`

const OptionDisallowAdditionalPropertiesPass = `{"rating2": 12345}`
