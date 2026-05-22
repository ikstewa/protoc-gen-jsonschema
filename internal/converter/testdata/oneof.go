package testdata

const OneOf = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OneOf",
    "$defs": {
        "OneOf": {
            "allOf": [
                {
                    "oneOf": [
                        {
                            "not": {
                                "anyOf": [
                                    {
                                        "required": [
                                            "bar"
                                        ]
                                    },
                                    {
                                        "required": [
                                            "baz"
                                        ]
                                    }
                                ]
                            }
                        },
                        {
                            "required": [
                                "bar"
                            ]
                        },
                        {
                            "required": [
                                "baz"
                            ]
                        }
                    ]
                }
            ],
            "properties": {
                "bar": {
                    "$ref": "#/$defs/samples.OneOf.Bar"
                },
                "baz": {
                    "$ref": "#/$defs/samples.OneOf.Baz"
                },
                "something": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "something"
            ],
            "title": "One Of"
        },
        "samples.OneOf.Bar": {
            "properties": {
                "foo": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "foo"
            ],
            "title": "Bar"
        },
        "samples.OneOf.Baz": {
            "properties": {
                "foo": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "foo"
            ],
            "title": "Baz"
        }
    }
}`

const OneOfFail = `{
	"something": true,
	"bar": {"foo": 1},
	"baz": {"foo": "one"}
}`

const OneOfPass = `{
	"something": true,
	"bar": {"foo": 1}
}`
