package testdata

const NestedMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/NestedMessage",
    "$defs": {
        "NestedMessage": {
            "properties": {
                "payload": {
                    "$ref": "#/$defs/samples.PayloadMessage"
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Message"
        },
        "samples.PayloadMessage": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "topology": {
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
            "title": "Payload Message"
        }
    }
}`

const NestedMessageFail = `{
	"payload": {
		"topology": "ROUND"
	}
}`

const NestedMessagePass = `{
	"payload": {
		"topology": "FLAT"
	}
}`
