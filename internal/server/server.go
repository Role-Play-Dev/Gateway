package server

import (
	"context"
	"log"
	"net/http"

	_ "role-play-dev/backend/gateway/docs"
	"role-play-dev/backend/gateway/internal/config"
	"role-play-dev/backend/gateway/internal/handler/auth"
	"role-play-dev/backend/gateway/internal/handler/user"
	"role-play-dev/backend/gateway/internal/middleware"

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
			s := auth.NewService()
			rg := api.Group("/auth")

			rg.POST("/register", auth.Register(s))
			rg.POST("/verify", auth.Verify(s))
			rg.POST("/login", auth.Login(s))
			rg.GET("/refresh", auth.Refresh(s))
			rg.GET("/logout", auth.Logout(s))
		}

		{
			api.Use(middleware.Auth())

			{
				{ // User
					s := user.NewService()
					rg := api.Group("/user")

					rg.GET("/", user.Get(s))
				}
			}
		}
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
