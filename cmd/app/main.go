package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	guLogger "github.com/kvendingoldo/gu-common/pkg/logger"
	"github.com/kvendingoldo/gu-user-service/config"
	v1Grpc "github.com/kvendingoldo/gu-user-service/internal/apis/grpc/v1"
	v2 "github.com/kvendingoldo/gu-user-service/internal/apis/rest/v2"
	"github.com/kvendingoldo/gu-user-service/internal/models"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func startGRPCServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Config.GRPCPort))
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	server := grpc.NewServer()
	svc := &v1Grpc.UserServiceServer{}
	v1.RegisterUserServiceServer(server, svc)
	reflection.Register(server)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
}

func startHTTPServer() {
	router := gin.New()
	router.RedirectTrailingSlash = false
	router.Use(guLogger.GinLogger(config.Config.Logger), gin.Recovery())

	v2.NewRouter(router)

	if err := router.Run(fmt.Sprintf(":%v", config.Config.RestPort)); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}

func init() {
	err := config.Setup()
	if err != nil {
		fmt.Println(err)
	}
	models.Setup()
}

func main() {
	startGRPCServer()
	startHTTPServer()
}
