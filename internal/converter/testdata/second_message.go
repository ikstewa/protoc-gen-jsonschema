package testdata

const SecondMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/SecondMessage",
    "$defs": {
        "SecondMessage": {
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
            "title": "Second Message"
        }
    }
}`

const SecondMessageFail = `{"complete2": "hello"}`

const SecondMessagePass = `{"complete2": true}`
