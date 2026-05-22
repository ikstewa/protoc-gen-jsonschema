package testdata

const OptionMaxLength = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionMaxLength",
    "$defs": {
        "OptionMaxLength": {
            "properties": {
                "query": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 0
                },
                "result_per_page": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "query"
            ],
            "title": "Option Max Length"
        }
    }
}`

const OptionMaxLengthFail = `{
    "query": "abcdefghijklmnopqrstuvwxyz",
	"page_number": 4
}`

const OptionMaxLengthPass = `{
	"query": "abc",
	"page_number": 4
}`
