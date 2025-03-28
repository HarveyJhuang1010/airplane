// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/booking": {
            "post": {
                "description": "user booking ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "user booking ticket",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bookingpkg.AddBookingCond"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bookingpkg.AddBookingResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/booking/:id": {
            "get": {
                "description": "user get booking ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "user get booking ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bookingpkg.BookingResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "user cancel booking ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "user cancel booking ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "patch": {
                "description": "user edit booking ticket. ex: change seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "user edit booking ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bookingpkg.EditBookingCond"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/bookingpkg.BookingResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/flight": {
            "get": {
                "description": "get flight list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flight"
                ],
                "summary": "get flight list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "到達機場",
                        "name": "arrivalAirport",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "出發機場",
                        "name": "departureAirport",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "出發時間結束(YYYY-MM-DDTHH:MM:SSZ)",
                        "name": "departureTimeEndAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "出發時間起始(YYYY-MM-DDTHH:MM:SSZ)",
                        "name": "departureTimeStartAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "頁碼 page index",
                        "name": "pi",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "筆數 page size",
                        "name": "ps",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardListResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/flight.FlightResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/payment/notify": {
            "post": {
                "description": "3rd party payment gateway will call this API to notify payment result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Notify Payment Result",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.NotifyPaymentCond"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/apis.StandardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/apis.StandardResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "$ref": "#/definitions/apis.StandardError"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apis.Meta": {
            "type": "object",
            "properties": {
                "requestID": {
                    "type": "string",
                    "example": "38df1107-94b9-498e-b24f-a0e82b77031b"
                }
            }
        },
        "apis.Pagination": {
            "type": "object",
            "properties": {
                "pi": {
                    "description": "頁碼 page index",
                    "type": "integer"
                },
                "ps": {
                    "description": "筆數 page size",
                    "type": "integer"
                },
                "total_page": {
                    "description": "總頁數 total pages",
                    "type": "integer"
                },
                "total_row": {
                    "description": "總筆數 total items",
                    "type": "integer"
                }
            }
        },
        "apis.StandardError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "apis.StandardListResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "PM 定義的 code",
                    "type": "string"
                },
                "data": {
                    "description": "資料內容"
                },
                "error": {
                    "description": "Error 資訊"
                },
                "meta": {
                    "description": "Response 的 Meta 資訊",
                    "allOf": [
                        {
                            "$ref": "#/definitions/apis.Meta"
                        }
                    ]
                },
                "pagination": {
                    "$ref": "#/definitions/apis.Pagination"
                }
            }
        },
        "apis.StandardResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "PM 定義的 code",
                    "type": "string"
                },
                "data": {
                    "description": "資料內容"
                },
                "error": {
                    "description": "Error 資訊"
                },
                "meta": {
                    "description": "Response 的 Meta 資訊",
                    "allOf": [
                        {
                            "$ref": "#/definitions/apis.Meta"
                        }
                    ]
                }
            }
        },
        "bookingpkg.AddBookingCond": {
            "type": "object",
            "properties": {
                "cabinClassID": {
                    "type": "integer"
                },
                "countryCode": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "flightID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "seatID": {
                    "type": "integer"
                }
            }
        },
        "bookingpkg.AddBookingResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "bookingpkg.BookingResponse": {
            "type": "object",
            "properties": {
                "airlineCode": {
                    "type": "string"
                },
                "arrivalAirport": {
                    "type": "string"
                },
                "arrivalTime": {
                    "type": "string"
                },
                "baggageAllowance": {
                    "type": "string",
                    "example": "0"
                },
                "classCode": {
                    "type": "string",
                    "enum": [
                        "economy_standard",
                        "economy_flex",
                        "business_basic",
                        "business_standard"
                    ]
                },
                "departureAirport": {
                    "type": "string"
                },
                "departureTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expiredAt": {
                    "type": "string"
                },
                "flightNumber": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "phoneCountryCode": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "refundable": {
                    "type": "boolean"
                },
                "seatNumber": {
                    "type": "string"
                },
                "seatSelection": {
                    "type": "boolean"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "pending",
                        "confirming",
                        "confirmed",
                        "failed",
                        "cancelled"
                    ]
                }
            }
        },
        "bookingpkg.EditBookingCond": {
            "type": "object",
            "properties": {
                "cabinClassID": {
                    "type": "integer"
                },
                "seatID": {
                    "type": "integer"
                }
            }
        },
        "flight.CabinClassResponse": {
            "type": "object",
            "properties": {
                "baggageAllowance": {
                    "type": "string",
                    "example": "0"
                },
                "classCode": {
                    "type": "string",
                    "enum": [
                        "economy_standard",
                        "economy_flex",
                        "business_basic",
                        "business_standard"
                    ]
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "maxSeats": {
                    "type": "string",
                    "example": "0"
                },
                "price": {
                    "type": "number"
                },
                "refundable": {
                    "type": "boolean"
                },
                "remainSeats": {
                    "type": "string",
                    "example": "0"
                },
                "seatSelection": {
                    "type": "boolean"
                }
            }
        },
        "flight.FlightResponse": {
            "type": "object",
            "properties": {
                "airlineCode": {
                    "description": "航空公司代碼",
                    "type": "string"
                },
                "arrivalAirport": {
                    "description": "到達機場",
                    "type": "string"
                },
                "arrivalTime": {
                    "description": "到達時間",
                    "type": "string"
                },
                "cabinClasses": {
                    "description": "艙等",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/flight.CabinClassResponse"
                    }
                },
                "departureAirport": {
                    "description": "出發機場",
                    "type": "string"
                },
                "departureTime": {
                    "description": "出發時間",
                    "type": "string"
                },
                "flightNumber": {
                    "description": "航班號碼",
                    "type": "string"
                },
                "id": {
                    "description": "Flight ID",
                    "type": "string",
                    "example": "0"
                },
                "sellableSeats": {
                    "description": "可售座位數",
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "航班狀態",
                    "type": "string",
                    "enum": [
                        "scheduled",
                        "boarding",
                        "departed",
                        "arrived",
                        "cancelled"
                    ]
                }
            }
        },
        "payment.NotifyPaymentCond": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Payment ID",
                    "type": "string",
                    "example": "0"
                },
                "method": {
                    "description": "Payment Method",
                    "type": "string",
                    "enum": [
                        "credit_card",
                        "debit_card",
                        "bank_transfer"
                    ]
                },
                "paidAt": {
                    "description": "Paid time",
                    "type": "string"
                },
                "provider": {
                    "description": "Payment Provider",
                    "type": "string",
                    "enum": [
                        "stripe",
                        "paypal",
                        "line_pay",
                        "apple_pay",
                        "google_pay"
                    ]
                },
                "status": {
                    "description": "Payment Status",
                    "type": "string",
                    "enum": [
                        "pending",
                        "success",
                        "failed",
                        "cancelled"
                    ]
                },
                "transactionID": {
                    "description": "3rd Party Transaction ID",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Airplane - Flight Booking System",
	Description:      "This is a RESTFUL API documentation of Flight Booking System.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
