package server

import (
	"context"
	"log"
	"net/http"

	_ "role-play-dev/backend/gateway/docs"
	"role-play-dev/backend/gateway/internal/config"
	"role-play-dev/backend/gateway/internal/handler/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

type server struct {
	serv *http.Server
}

func NewServer(conf config.Config) Server {
	r := newRouter(conf)

	{
		api := r.Group("/api/v1")

		{
			authRG := api.Group("/auth")

			s := auth.NewService()
			h := auth.NewHandler(s, conf)

			authRG.POST("/creds/reg/link/send", h.CredentialsRegisterLinkSend())
			authRG.POST("/creds/reg/link/confirm", h.CredentialsRegisterLinkConfirm())
			authRG.POST("/creds/login", h.CredentialsLogin())
			authRG.GET("/token/refresh", h.TokenRefresh())
			authRG.GET("/logout", h.Logout())
		}

		// api.Use(middleware.Auth())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &server{
		serv: &http.Server{
			Addr:    conf.ServerAddress,
			Handler: r.Handler(),
		},
	}
}

func (s *server) Start() error {
	log.Println("server listen at", s.serv.Addr)
	if err := s.serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}

func newRouter(conf config.Config) *gin.Engine {
	if conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	return router
}
