package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
//
//	@Summary		Login by email and password
//	@Description	Returns access token in body and sets refresh token to cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/login [post]
//	@Param			request	body		LoginRequest	true	"Login request"
//	@Success		200		{object}	LoginResponce
func Register(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.register(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

// Login godoc
//
//	@Summary		Login by email and password
//	@Description	Returns access token in body and sets refresh token to cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/login [post]
//	@Param			request	body		LoginRequest	true	"Login request"
//	@Success		200		{object}	LoginResponce
func Login(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.login(ctx); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	}
}

// Logout godoc
//
//	@Summary		Logout from account
//	@Description	Clears refresh token from cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/logout [get]
//	@Success		200	{object}	LoginResponce
func Logout(s Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := s.logout(ctx); err != nil {
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
