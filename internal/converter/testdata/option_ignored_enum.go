package testdata

const UnignoredEnum = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ],
    "enum": [
        "VALUE_4",
        0,
        "VALUE_5",
        1,
        "VALUE_6",
        2,
        "VALUE_7",
        3
    ],
    "title": "Unignored Enum"
}`
