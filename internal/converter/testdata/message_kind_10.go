package testdata

const MessageKind10 = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/MessageKind10",
    "$defs": {
        "MessageKind10": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "timestamp": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 10"
        }
    }
}`
