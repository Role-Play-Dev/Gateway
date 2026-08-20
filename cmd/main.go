package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "role-play-dev/backend/gateway/docs"
	"role-play-dev/backend/gateway/internal/config"
	"role-play-dev/backend/gateway/internal/server"
)

//	@title			Role-Play-Dev Gateway service API
//	@version		1.0
//	@description	Backend Gateway service API of a tabletop RPG helper
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/v1

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	conf := config.NewConfig()
	serv := server.NewServer(conf)

	go func() {
		if err := serv.Start(); err != nil {
			log.Fatalf("server start error: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("server shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serv.Stop(ctx); err != nil {
		log.Println("server shutdown error:", err.Error())
	}

	log.Println("server shutdown")
}
