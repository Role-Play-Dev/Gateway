package auth

import (
	"fmt"
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
//	@Summary		Send registration link to email
//	@Description	Saves registration session and sends registration link to client with session ID in token in query to provided email address
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
		var req CredentialsRegisterLinkSendRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.CredentialsRegisterLinkSend(ctx, req); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
	}
}

// CredentialsRegisterLinkConfirm godoc
//
//	@Summary		Confirm email and fill credentials
//	@Description	Accepts token in query and user credential in body, verifies token and creates account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			auth_token	query		string									true	"Auth token"
//	@Param			request		body		CredentialsRegisterLinkConfirmRequest	true	"CredentialsRegisterLinkConfirm request"
//	@Success		200			{object}	httputil.emptyResponce
//	@Failure		400			{object}	httputil.errorResponce
//	@Failure		500			{object}	httputil.errorResponce
//	@Router			/auth/creds/reg/link/confirm [post]
func (h *handler) CredentialsRegisterLinkConfirm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ex := ctx.GetQuery("auth_token")
		if !ex {
			httputil.Error(ctx, http.StatusBadRequest, fmt.Errorf(
				"Auth token is not correctly provided",
			))
		}

		var req CredentialsRegisterLinkConfirmRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		if err := h.serv.CredentialsRegisterLinkConfirm(ctx, token, req); err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.Empty(ctx, http.StatusOK)
	}
}

// CredentialsLogin godoc
//
//	@Summary		Login by user credentials
//	@Description	Accepts email and password, returns access token in body and sets refresh token to cookie
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
		var req CredentialsLoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			httputil.Error(ctx, http.StatusBadRequest, err)
		}

		resp, refToken, err := h.serv.CredentialsLogin(ctx, req)
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
//	@Summary		Refresh access token
//	@Description	Validates refresh token then generates and sends new access token
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

		resp, err := h.serv.TokenRefresh(ctx, refToken)
		if err != nil {
			httputil.Error(ctx, http.StatusInternalServerError, err)
		}

		httputil.JSON(ctx, http.StatusOK, resp)
	}
}
