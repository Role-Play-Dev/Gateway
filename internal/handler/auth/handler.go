package auth

import (
	"net/http"
	"role-play-dev/backend/gateway/internal/config"
	"role-play-dev/backend/gateway/internal/util/httputil"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serv Service
	conf config.Config
}

func NewHandler(serv Service, conf config.Config) *handler {
	return &handler{
		serv: serv,
		conf: conf,
	}
}

// Register godoc
//
//	@Summary		Register email to send verification link
//	@Description	Sends verification link to provided email address
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/register [post]
//	@Param			request	body		RegisterRequest	true	"Register email request"
//	@Success		200		{object}	RegisterRequest
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody RegisterRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.register(ctx, reqBody.Email); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
	}
}

// Verify godoc
//
//	@Summary		Verifies email by token and applies user data
//	@Description	Checks if provided token equals to stored one and saves user data
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/verify [post]
//	@Param			request	body		VerifyRequest	true	"Verify request"
//	@Success		200		{object}	httputil.emptyResponce
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody VerifyRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.verify(ctx, reqBody.Token, reqBody.Username, reqBody.Password); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
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
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody LoginRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		accToken, refToken, err := h.serv.login(ctx, reqBody.Username, reqBody.Password)
		if err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		ctx.SetCookie(
			h.conf.Token.Cookie.Name,
			refToken,
			h.conf.Token.Cookie.MaxAge,
			h.conf.Token.Cookie.Path,
			h.conf.ClientAddress,
			h.conf.Token.Cookie.Secure,
			h.conf.Token.Cookie.HttpOnly,
		)

		httputil.JSON(ctx, http.StatusOK, LoginResponce{
			AccessToken: accToken,
		})
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
//	@Success		200	{object}	httputil.emptyResponce
//	@Failure		500	{object}	httputil.errorResponce
func (h *handler) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := h.serv.logout(ctx); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		ctx.SetCookie(
			h.conf.Token.Cookie.Name,
			"",
			h.conf.Token.Cookie.MaxAge,
			h.conf.Token.Cookie.Path,
			h.conf.ClientAddress,
			h.conf.Token.Cookie.Secure,
			h.conf.Token.Cookie.HttpOnly,
		)

		httputil.Empty(ctx, http.StatusOK)
	}
}

// Refresh godoc
//
//	@Summary		Refresh access token
//	@Description	Refreshs access token using refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/refresh [get]
//	@Success		200	{object}	RefreshResponce
//	@Failure		400	{object}	httputil.errorResponce
//	@Failure		500	{object}	httputil.errorResponce
func (h *handler) Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		accToken, err := h.serv.refresh(ctx, refToken)
		if err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.JSON(ctx, http.StatusOK, RefreshResponce{
			AccessToken: accToken,
		})
	}
}
