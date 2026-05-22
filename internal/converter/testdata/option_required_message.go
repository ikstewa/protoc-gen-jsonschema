package testdata

const OptionRequiredMessage = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/OptionRequiredMessage",
    "$defs": {
        "OptionRequiredMessage": {
            "properties": {
                "name2": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "timestamp2": {
                    "type": "string",
                    "maxLength": 0,
                    "minLength": 0
                },
                "id2": {
                    "type": "integer"
                },
                "rating2": {
                    "type": "number"
                },
                "complete2": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "required": [
                "name2",
                "timestamp2",
                "id2",
                "rating2",
                "complete2"
            ],
            "title": "Option Required Message"
        }
    }
}`

const OptionRequiredMessageFail = `{
	"name2": "some name",
	"id2": 1
}`

const OptionRequiredMessagePass = `{
	"name2": "some name",
	"timestamp2": "1970-01-01T00:00:00Z",
	"id2": 1,
	"rating2": 100,
	"complete2": true
}`
