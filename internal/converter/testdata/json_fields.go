package testdata

const JSONFields = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/JSONFields",
    "$defs": {
        "JSONFields": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "identifier": {
                    "type": "integer"
                },
                "someThing": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "snakeNumb": {
                    "type": "string"
                },
                "otherNumb": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "otherNumb"
            ],
            "title": "JSON Fields"
        }
    }
}`

const JSONFieldsFail = `{"someThing": "onetwothree", "other_numb": 123}`

const JSONFieldsPass = `{"someThing": 12345, "otherNumb": 123}`
