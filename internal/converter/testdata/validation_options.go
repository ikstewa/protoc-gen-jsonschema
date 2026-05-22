package testdata

const ValidationOptions = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/ValidationOptions",
    "$defs": {
        "ValidationOptions": {
            "properties": {
                "stringWithLengthConstraints": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 5
                },
                "luckyNumbersWithArrayConstraints": {
                    "items": {
                        "type": "integer"
                    },
                    "type": "array",
                    "maxItems": 6,
                    "minItems": 2
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Validation Options"
        }
    }
}`

const ValidationOptionsFail = `{
	"stringWithLengthConstraints": "this string is way too long",
	"luckyNumbersWithArrayConstraints": [1]
}`

const ValidationOptionsPass = `{
	"stringWithLengthConstraints": "thisisok",
	"luckyNumbersWithArrayConstraints": [1,2,3,4]
}`
