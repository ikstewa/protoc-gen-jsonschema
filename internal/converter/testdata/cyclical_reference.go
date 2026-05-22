package testdata

const CyclicalReferenceMessageM = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/M",
    "$defs": {
        "M": {
            "properties": {
                "foo": {
                    "$ref": "#/$defs/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "M"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/$defs/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bar"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/$defs/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Baz"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "bar": {
                    "items": {
                        "$ref": "#/$defs/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Foo"
        }
    }
}`

const CyclicalReferenceMessageFoo = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Foo",
    "$defs": {
        "Foo": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "bar": {
                    "items": {
                        "$ref": "#/$defs/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Foo"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/$defs/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bar"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/$defs/Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Baz"
        }
    }
}`

const CyclicalReferenceMessageBar = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Bar",
    "$defs": {
        "Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/$defs/samples.Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bar"
        },
        "samples.Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/$defs/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Baz"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "bar": {
                    "items": {
                        "$ref": "#/$defs/Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Foo"
        }
    }
}`

const CyclicalReferenceMessageBaz = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Baz",
    "$defs": {
        "Baz": {
            "properties": {
                "enabled": {
                    "type": "boolean"
                },
                "foo": {
                    "$ref": "#/$defs/samples.Foo",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Baz"
        },
        "samples.Bar": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "baz": {
                    "$ref": "#/$defs/Baz",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Bar"
        },
        "samples.Foo": {
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "bar": {
                    "items": {
                        "$ref": "#/$defs/samples.Bar"
                    },
                    "type": "array"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Foo"
        }
    }
}`
