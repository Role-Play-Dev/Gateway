package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.login(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

func Logout(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.logout(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

func Register(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.register(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

func Refresh(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.refresh(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}
