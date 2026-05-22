package testdata

const NestedObject = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/NestedObject",
    "$defs": {
        "NestedObject": {
            "properties": {
                "payload": {
                    "$ref": "#/$defs/samples.NestedObject.NestedPayload"
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Object"
        },
        "samples.NestedObject.NestedPayload": {
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
            "title": "Nested Payload"
        }
    }
}`

const NestedObjectFail = `{"payload": false}`

const NestedObjectPass = `{
	"payload": {
	  "topology": "NESTED_OBJECT"
	}
}`
