package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.get(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}
