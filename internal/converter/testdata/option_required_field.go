package testdata

const OptionRequiredField = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionRequiredField",
    "$defs": {
        "OptionRequiredField": {
            "properties": {
                "query": {
                    "type": "string"
                },
                "page_number": {
                    "type": "integer"
                },
                "result_per_page": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "query",
                "page_number"
            ],
            "title": "Option Required Field"
        }
    }
}`

const OptionRequiredFieldFail = `{
	"page_number": 4
}`

const OptionRequiredFieldPass = `{
	"query": "what?",
	"page_number": 4
}`
