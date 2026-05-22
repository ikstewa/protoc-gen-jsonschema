package testdata

const WellKnown = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/WellKnown",
    "$defs": {
        "WellKnown": {
            "properties": {
                "string_value": {
                    "additionalProperties": true,
                    "type": "string"
                },
                "map_of_integers": {
                    "additionalProperties": {
                        "additionalProperties": true,
                        "type": "integer"
                    },
                    "type": "object"
                },
                "map_of_scalar_integers": {
                    "additionalProperties": {
                        "type": "integer"
                    },
                    "type": "object"
                },
                "list_of_integers": {
                    "items": {
                        "type": "integer",
                        "title": "Int 32 Value",
                        "description": "Wrapper message for ` + "`" + `int32` + "`" + `. The JSON representation for ` + "`" + `Int32Value` + "`" + ` is JSON number. Not recommended for use in new APIs, but still useful for legacy APIs and has no plan to be removed."
                    },
                    "type": "array"
                },
                "duration": {
                    "type": "string",
                    "pattern": "^([0-9]+\\.?[0-9]*|\\.[0-9]+)s$",
                    "format": "regex",
                    "description": "This is a duration:"
                },
                "struct": {
                    "additionalProperties": true,
                    "type": "object"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Well Known"
        }
    }
}`

const WellKnownFail = `{"duration": "9"}`

const WellKnownPass = `{"duration": "9s"}`
