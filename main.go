package main

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/models"
	"github.com/kvendingoldo/gu-user-service/routes"
	"google.golang.org/grpc"
	"log"
	"net"
)

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	s := grpc.NewServer()
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
}

func startHTTPServer() {
	router := gin.Default()

	routes.ApplicationV1Router(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}

func init() {
	cfg.Config = cfg.GetConfig()
	cfg.Config.DB.AutoMigrate(&models.User{})

}

func main() {
	startHTTPServer()

	//go startGRPCServer()
	//
	//sig := make(chan os.Signal)
	//signal.Notify(sig, os.Interrupt, os.Kill)
	//<-sig
}
