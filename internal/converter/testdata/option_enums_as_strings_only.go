package testdata

const OptionEnumsAsStringsOnly = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "type": "string",
    "enum": [
        "NOT_SPECIFIED",
        "USD",
        "GBP",
        "EUR"
    ],
    "title": "Currency"
}`

const OptionEnumsAsStringsOnlyPass = `"NOT_SPECIFIED"`
const OptionEnumsAsStringsOnlyFail = `2`
