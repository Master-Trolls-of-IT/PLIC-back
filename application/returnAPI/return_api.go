package returnAPI

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReturnAPI struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var StatusMessage = map[int]string{
	http.StatusOK:                  "Request succeeded",
	http.StatusCreated:             "Created successfully",
	http.StatusAccepted:            "Request accepted",
	http.StatusNoContent:           "No content",
	http.StatusBadRequest:          "Bad request",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "Not found",
	http.StatusMethodNotAllowed:    "Method not allowed",
	http.StatusConflict:            "Conflict",
	http.StatusInternalServerError: "Internal server error",
	http.StatusNotImplemented:      "Not implemented",
}

func replaceNilObject(data interface{}) interface{} {
	if data != nil {
		return data
	} else {
		return gin.H{}
	}
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
		Data:    replaceNilObject(data),
	})
}
