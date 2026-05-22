package testdata

const PayloadMessage2 = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/PayloadMessage2",
    "$defs": {
        "PayloadMessage2": {
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
                },
                "array": {
                    "items": {
                        "type": "integer"
                    },
                    "type": "array"
                },
                "opt_int": {
                    "type": "integer"
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
            "title": "Payload Message 2",
            "description": "PayloadMessage2 contains some common types  PayloadMessage2 is used throughout the test suite and can have multi-line comments"
        }
    }
}`

const PayloadMessage2Fail = `{
}`

const PayloadMessage2Pass = `{
    "name": "test",
    "timestamp": "1970-01-01T00:00:00Z",
    "id": 1,
    "rating": 100,
    "complete": true,
    "topology": "FLAT"
}`
