package testdata

const MessageKind12 = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/MessageKind12",
    "$defs": {
        "MessageKind12": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "f": {
                    "$ref": "#/$defs/samples.MessageKind11"
                },
                "kind5": {
                    "$ref": "#/$defs/samples.MessageKind5"
                },
                "kind6": {
                    "$ref": "#/$defs/samples.MessageKind6"
                },
                "kind7": {
                    "$ref": "#/$defs/samples.MessageKind7"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 12"
        },
        "samples.MessageKind1": {
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
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 1"
        },
        "samples.MessageKind11": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "ones": {
                    "items": {
                        "$ref": "#/$defs/samples.MessageKind1"
                    },
                    "type": "array"
                },
                "kind2": {
                    "$ref": "#/$defs/samples.MessageKind2"
                },
                "kind3": {
                    "$ref": "#/$defs/samples.MessageKind3"
                },
                "kind4": {
                    "$ref": "#/$defs/samples.MessageKind4"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 11"
        },
        "samples.MessageKind2": {
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
                "someProp": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 3"
        },
        "samples.MessageKind4": {
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
                "special": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 4"
        },
        "samples.MessageKind5": {
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
                "foo": {
                    "type": "number"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 5"
        },
        "samples.MessageKind6": {
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
                "bar": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 6"
        },
        "samples.MessageKind7": {
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
                "baz": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 7"
        }
    }
}`
