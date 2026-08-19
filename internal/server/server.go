package server

import (
	"context"
	"log"
	"net/http"

	"role-play-dev/backend/gateway/internal/config"
	"role-play-dev/backend/gateway/internal/handler/auth"
	"role-play-dev/backend/gateway/internal/handler/user"
	"role-play-dev/backend/gateway/internal/middleware"

	"github.com/gin-gonic/gin"
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

	{ // V1 group
		v1 := r.Group("/v1")

		{ // Auth group
			s := auth.NewService()

			rg := v1.Group("/auth")

			rg.GET("/login", auth.Login(s))
			rg.POST("/register", auth.Register(s))
			rg.GET("/refresh", auth.Refresh(s))
			rg.GET("/logout", auth.Logout(s))
		}

		{ // API group
			api := v1.Group("/api")
			api.Use(middleware.Auth())

			{
				s := user.NewService()

				rg := api.Group("/user")

				rg.GET("/", user.Get(s))
			}
		}
	}

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
