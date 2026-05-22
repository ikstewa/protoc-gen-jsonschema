package testdata

const SelfReference = `{
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
                        "$ref": "#/$defs/Foo"
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

const SelfReferenceFail = `{
	"bar": [
		{
			"name": false
		}
	]
}`

const SelfReferencePass = `{
	"bar": [
		{
			"name": "referenced-bar",
			"bar": [
				{
					"name": "barception"
				}
			]
		}
	]
}`
