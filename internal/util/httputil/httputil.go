package httputil

import (
	"github.com/gin-gonic/gin"
)

type emptyResponce struct {
	Code int `json:"code"`
}

type errorResponce struct {
	Code    int    `json:"code"`
	Message string `json:"message" example:"error description"`
}

func Empty(ctx *gin.Context, code int) {
	ctx.JSON(code, emptyResponce{
		Code: code,
	})
}

func Error(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, errorResponce{
		Code:    code,
		Message: err.Error(),
	})
}

func JSON(ctx *gin.Context, code int, obj any) {
	ctx.JSON(code, obj)
}
