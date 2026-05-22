package testdata

const Proto2NestedObject = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Proto2NestedObject",
    "$defs": {
        "Proto2NestedObject": {
            "properties": {
                "payload": {
                    "$ref": "#/$defs/samples.Proto2NestedObject.NestedPayload",
                    "additionalProperties": {
                        "not": true
                    }
                },
                "description": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "payload",
                "description"
            ],
            "title": "Proto 2 Nested Object"
        },
        "samples.Proto2NestedObject.NestedPayload": {
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
            "required": [
                "name",
                "timestamp",
                "id",
                "rating",
                "complete",
                "topology"
            ],
            "title": "Nested Payload"
        }
    }
}`

const Proto2NestedObjectFail = `{
	"payload": {
		"topology": "FLAT"	
	}
}`

const Proto2NestedObjectPass = `{
	"description": "lots of attributes",
	"payload": {
		"name": "something",
		"timestamp": "1970-01-01T00:00:00Z",
		"id": 1,
		"rating": 100,
		"complete": true,
		"topology": "FLAT"
	}
}`
