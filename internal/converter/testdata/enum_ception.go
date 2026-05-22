package testdata

const EnumCeption = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/$defs/Enumception",
    "$defs": {
        "Enumception": {
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
                "failureMode": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ],
                    "enum": [
                        "RECURSION_ERROR",
                        0,
                        "SYNTAX_ERROR",
                        1
                    ],
                    "title": "Failure Modes",
                    "description": "FailureModes enum"
                },
                "payload": {
                    "$ref": "#/$defs/samples.PayloadMessage",
                    "additionalProperties": true
                },
                "payloads": {
                    "items": {
                        "$ref": "#/$defs/samples.PayloadMessage"
                    },
                    "type": "array"
                },
                "importedEnum": {
                    "oneOf": [
                        {
                            "const": "VALUE_0",
                            "description": "Zero"
                        },
                        {
                            "const": 0,
                            "description": "Zero"
                        },
                        {
                            "const": "VALUE_1",
                            "description": "One"
                        },
                        {
                            "const": 1,
                            "description": "One"
                        },
                        {
                            "const": "VALUE_2",
                            "description": "Two"
                        },
                        {
                            "const": 2,
                            "description": "Two"
                        },
                        {
                            "const": "VALUE_3",
                            "description": "Three"
                        },
                        {
                            "const": 3,
                            "description": "Three"
                        }
                    ],
                    "enum": [
                        "VALUE_0",
                        0,
                        "VALUE_1",
                        1,
                        "VALUE_2",
                        2,
                        "VALUE_3",
                        3
                    ],
                    "title": "Imported Enum",
                    "description": "This is an enum"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Enumception"
        },
        "samples.PayloadMessage": {
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
            "title": "Payload Message"
        }
    }
}`

const EnumCeptionFail = `{"payloads": [ {"topology": "MAP"} ]}`

const EnumCeptionPass = `{"payloads": [ {"topology": "ARRAY_OF_MESSAGE"} ]}`
