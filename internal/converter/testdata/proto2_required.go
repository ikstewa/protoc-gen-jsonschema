package testdata

const Proto2Required = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Proto2Required",
    "$defs": {
        "Proto2Required": {
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
                "query"
            ],
            "title": "Proto 2 Required"
        }
    }
}`

const Proto2RequiredFail = `{
	"page_number": 4
}`

const Proto2RequiredPass = `{
	"query": "what?",
	"page_number": 4
}`
