package returnAPI

import (
	"github.com/gin-gonic/gin"
)

type ReturnAPI struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
)

var StatusMessage = map[int]string{
	StatusOK:                  "Request succeeded",
	StatusCreated:             "Created successfully",
	StatusAccepted:            "Request accepted",
	StatusNoContent:           "No content",
	StatusBadRequest:          "Bad request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not found",
	StatusMethodNotAllowed:    "Method not allowed",
	StatusConflict:            "Conflict",
	StatusInternalServerError: "Internal server error",
	StatusNotImplemented:      "Not implemented",
}

func Error(context *gin.Context, statusCode int) {
	context.JSON(statusCode, ReturnAPI{
		Code:    statusCode,
		Message: StatusMessage[statusCode],
		Data:    nil,
	})
}

func Success(context *gin.Context, statusCode int, data interface{}) {
	context.JSON(statusCode, ReturnAPI{
		Code:    statusCode,
		Message: StatusMessage[statusCode],
		Data:    data,
	})
}
