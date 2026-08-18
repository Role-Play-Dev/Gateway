package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/role-play-dev/backend/gateway/internal/config"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
}

type server struct {
	conf config.Config
	rout *gin.Engine
	serv *http.Server
}

func NewServer(conf config.Config) Server {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	return &server{
		conf: conf,
		rout: router,
		serv: &http.Server{
			Addr:    conf.ServerAddress,
			Handler: router.Handler(),
		},
	}
}

func (s *server) Start() error {
	log.Println("server listen at", s.serv.Addr)
	if err := s.serv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}
