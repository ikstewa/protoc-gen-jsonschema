package testdata

const OptionMinLength = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionMinLength",
    "$defs": {
        "OptionMinLength": {
            "properties": {
                "query": {
                    "type": "string",
                    "minLength": 2
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
            "title": "Option Min Length"
        }
    }
}`

const OptionMinLengthFail = `{
    "query": "a",
	"page_number": 4
}`

const OptionMinLengthPass = `{
	"query": "what?",
	"page_number": 4
}`
