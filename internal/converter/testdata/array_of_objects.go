package testdata

const ArrayOfObjects = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/ArrayOfObjects",
    "$defs": {
        "ArrayOfObjects": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "payload": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "items": {
                                "$ref": "#/$defs/samples.ArrayOfObjects.RepeatedPayload"
                            },
                            "type": "array"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "title": "Array Of Objects"
        },
        "samples.ArrayOfObjects.RepeatedPayload": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "name": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "timestamp": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "id": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                },
                "rating": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "number"
                        }
                    ]
                },
                "complete": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "boolean"
                        }
                    ]
                },
                "topology": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        },
                        {
                            "type": "null"
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
            "title": "Repeated Payload"
        }
    }
}`

const ArrayOfObjectsFail = `{
    "description": "something",
    "payload": [
        {
            "topology": "cruft"
        }
    ]
}`

const ArrayOfObjectsPass = `{
    "description": "something",
    "payload": [
        {
            "topology": "ARRAY_OF_OBJECT"
        }
    ]
}`
