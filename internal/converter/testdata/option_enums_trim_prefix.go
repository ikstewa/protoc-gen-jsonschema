package testdata

const OptionEnumsTrimPrefix = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "type": "string",
    "enum": [
        "UNSPECIFIED",
        "HTTP",
        "HTTPS"
    ],
    "title": "Scheme"
}`

const OptionEnumsTrimPrefixPass = `"HTTP"`

const OptionEnumsTrimPrefixFail = `4`
