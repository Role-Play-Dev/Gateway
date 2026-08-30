package middleware

import "github.com/gin-gonic/gin"

func Protected() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
