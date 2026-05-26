package testdata

const NumericBounds = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/NumericBounds",
    "$defs": {
        "NumericBounds": {
            "properties": {
                "int32_val": {
                    "type": "integer",
                    "maximum": 2147483647,
                    "minimum": -2147483648
                },
                "sint32_val": {
                    "type": "integer",
                    "maximum": 2147483647,
                    "minimum": -2147483648
                },
                "sfixed32_val": {
                    "type": "integer",
                    "maximum": 2147483647,
                    "minimum": -2147483648
                },
                "uint32_val": {
                    "type": "integer",
                    "maximum": 4294967295,
                    "minimum": 0
                },
                "fixed32_val": {
                    "type": "integer",
                    "maximum": 4294967295,
                    "minimum": 0
                },
                "int64_val": {
                    "type": "string"
                },
                "sint64_val": {
                    "type": "string"
                },
                "sfixed64_val": {
                    "type": "string"
                },
                "uint64_val": {
                    "type": "string"
                },
                "fixed64_val": {
                    "type": "string"
                },
                "float_val": {
                    "type": "number",
                    "maximum": 3.4028234663852886e+38,
                    "minimum": -3.4028234663852886e+38
                },
                "double_val": {
                    "type": "number",
                    "maximum": 1.7976931348623157e+308,
                    "minimum": -1.7976931348623157e+308
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Numeric Bounds"
        }
    }
}`
