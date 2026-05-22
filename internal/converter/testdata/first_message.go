package testdata

const FirstMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/FirstMessage",
    "$defs": {
        "FirstMessage": {
            "properties": {
                "name1": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "timestamp1": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "id1": {
                    "type": "integer"
                },
                "rating1": {
                    "type": "number"
                },
                "complete1": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "First Message"
        }
    }
}`

const FirstMessageFail = `{"complete1": "hello"}`

const FirstMessagePass = `{"complete1": true}`
