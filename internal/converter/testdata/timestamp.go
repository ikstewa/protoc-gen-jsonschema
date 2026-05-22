package testdata

const Timestamp = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/Timestamp",
    "$defs": {
        "Timestamp": {
            "properties": {
                "timestamp": {
                    "type": "string",
                    "format": "date-time"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Timestamp"
        }
    }
}`

const TimestampFail = `{"timestamp": "twelve oclock"}`

const TimestampPass = `{"timestamp": "1970-01-01T00:00:00Z"}`
