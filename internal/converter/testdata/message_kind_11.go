package testdata

const MessageKind11 = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/MessageKind11",
    "$defs": {
        "MessageKind11": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "ones": {
                    "items": {
                        "$ref": "#/$defs/samples.MessageKind1"
                    },
                    "type": "array"
                },
                "kind2": {
                    "$ref": "#/$defs/samples.MessageKind2",
                    "additionalProperties": true
                },
                "kind3": {
                    "$ref": "#/$defs/samples.MessageKind3",
                    "additionalProperties": true
                },
                "kind4": {
                    "$ref": "#/$defs/samples.MessageKind4",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 11"
        },
        "samples.MessageKind1": {
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
            "title": "Message Kind 1"
        },
        "samples.MessageKind2": {
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
                "isa": {
                    "type": "boolean"
                },
                "hasa": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 2"
        },
        "samples.MessageKind3": {
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
                "someProp": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 3"
        },
        "samples.MessageKind4": {
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
                "special": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 4"
        }
    }
}`
