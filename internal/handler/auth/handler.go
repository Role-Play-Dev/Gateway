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

// CredentialsRegisterLinkSend godoc
//
//	@Summary		Sends registration link to email
//	@Description	Sends registration link with token in query to provided email address
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/creds/reg/link/send [post]
//	@Param			request	body		CredentialsRegisterLinkSendRequest	true	"Send link request"
//	@Success		200		{object}	httputil.emptyResponce
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) CredentialsRegisterLinkSend() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody CredentialsRegisterLinkSendRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.credentialsRegisterLinkSend(ctx, reqBody); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
	}
}

// CredentialsRegisterLinkConfirm godoc
//
//	@Summary		Verify email and register user
//	@Description	Verifies token from link and register user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/creds/reg/link/confirm [post]
//	@Param			request	body		CredentialsRegisterLinkConfirmRequest	true	"CredentialsRegisterLinkConfirm request"
//	@Success		200		{object}	httputil.emptyResponce
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) CredentialsRegisterLinkConfirm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody CredentialsRegisterLinkConfirmRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.credentialsRegisterLinkConfirm(ctx, reqBody); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
	}
}

// CredentialsLogin godoc
//
//	@Summary		CredentialsLogin into by email and password
//	@Description	Returns access token in body and sets refresh token to cookie
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/creds/login [post]
//	@Param			request	body		CredentialsLoginRequest	true	"CredentialsLogin request"
//	@Success		200		{object}	LoginResponce
//	@Failure		400		{object}	httputil.errorResponce
//	@Failure		500		{object}	httputil.errorResponce
func (h *handler) CredentialsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody CredentialsLoginRequest
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		resp, refToken, err := h.serv.credentialsLogin(ctx, reqBody)
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

		httputil.JSON(ctx, http.StatusOK, resp)
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

// TokenRefresh godoc
//
//	@Summary		TokenRefresh access token
//	@Description	Refreshs access token using refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Router			/auth/token/refresh [get]
//	@Success		200	{object}	TokenRefreshResponce
//	@Failure		400	{object}	httputil.errorResponce
//	@Failure		500	{object}	httputil.errorResponce
func (h *handler) TokenRefresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refToken, err := ctx.Cookie("refresh_token")
		if err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		resp, err := h.serv.tokenRefresh(ctx, refToken)
		if err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.JSON(ctx, http.StatusOK, resp)
	}
}
