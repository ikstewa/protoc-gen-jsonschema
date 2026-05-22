package testdata

const SecondEnum = `{
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
    "title": "Second Enum"
}`

const SecondEnumFail = `"VALUE_3"`

const SecondEnumPass = `"VALUE_7"`
