package testdata

const ImportedEnum = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "oneOf": [
        {
            "const": "VALUE_0",
            "description": "Zero"
        },
        {
            "const": 0,
            "description": "Zero"
        },
        {
            "const": "VALUE_1",
            "description": "One"
        },
        {
            "const": 1,
            "description": "One"
        },
        {
            "const": "VALUE_2",
            "description": "Two"
        },
        {
            "const": 2,
            "description": "Two"
        },
        {
            "const": "VALUE_3",
            "description": "Three"
        },
        {
            "const": 3,
            "description": "Three"
        }
    ],
    "enum": [
        "VALUE_0",
        0,
        "VALUE_1",
        1,
        "VALUE_2",
        2,
        "VALUE_3",
        3
    ],
    "title": "Imported Enum",
    "description": "This is an enum"
}`

const ImportedEnumFail = `"VALUE_5"`

const ImportedEnumPass = `"VALUE_3"`
