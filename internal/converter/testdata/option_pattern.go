package testdata

const OptionPattern = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionPattern",
    "$defs": {
        "OptionPattern": {
            "properties": {
                "query": {
                    "type": "string",
                    "pattern": "^(\\([0-9]{3}\\))?[0-9]{3}-[0-9]{4}$"
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
            "title": "Option Pattern"
        }
    }
}`

const OptionPatternFail = `{
    "query": "a",
	"page_number": 4
}`

const OptionPatternPass = `{
	"query": "(888)555-1212",
	"page_number": 4
}`
