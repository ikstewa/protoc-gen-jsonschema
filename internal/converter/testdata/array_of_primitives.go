package testdata

const ArrayOfPrimitives = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/ArrayOfPrimitives",
    "$defs": {
        "ArrayOfPrimitives": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "luckyNumbers": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "luckyBigNumbers": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "null"
                            }
                        ]
                    }
                },
                "keyWords": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    }
                },
                "big_number": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "title": "Array Of Primitives"
        }
    }
}`

const ArrayOfPrimitivesFail = `{"luckyNumbers": ["false"]}`

const ArrayOfPrimitivesPass = `{"luckyNumbers": [1,2,3]}`

const ArrayOfPrimitivesDouble = `{
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$ref": "#/$defs/ArrayOfPrimitives",
    "$defs": {
        "ArrayOfPrimitives": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "properties": {
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "luckyNumbers": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    }
                },
                "luckyBigNumbers": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "string"
                            },
                            {
                                "type": "null"
                            }
                        ]
                    }
                },
                "keyWords": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ],
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    }
                },
                "big_number": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "null"
                        }
                    ]
                },
                "bigNumber": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "title": "Array Of Primitives"
        }
    }
}`

const ArrayOfPrimitivesDoubleFail = `{"bigNumber": false}`

const ArrayOfPrimitivesDoublePass = `{"bigNumber": "2"}`
