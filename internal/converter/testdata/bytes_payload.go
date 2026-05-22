package testdata

const BytesPayload = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/BytesPayload",
    "$defs": {
        "BytesPayload": {
            "properties": {
                "description": {
                    "type": "string"
                },
                "payload": {
                    "type": "string",
                    "format": "binary",
                    "contentEncoding": "base64"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bytes Payload"
        }
    }
}`

const BytesPayloadFail = `{"payload": 12345}`
