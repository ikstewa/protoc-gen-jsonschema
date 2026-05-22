package testdata

const UnignoredMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/UnignoredMessage",
    "$defs": {
        "UnignoredMessage": {
            "properties": {
                "name2": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "timestamp2": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
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
            "additionalProperties": true,
            "type": "object",
            "title": "Unignored Message"
        }
    }
}`
