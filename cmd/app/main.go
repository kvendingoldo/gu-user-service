package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cfg "github.com/kvendingoldo/gu-user-service/config"
	grpcSvc "github.com/kvendingoldo/gu-user-service/internal/server/grpc"
	"github.com/kvendingoldo/gu-user-service/internal/server/rest/router"
	"github.com/kvendingoldo/gu-user-service/model"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func startGRPCServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Config.GRPCPort))
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	server := grpc.NewServer()
	svc := &grpcSvc.UserServiceServer{}
	v1.RegisterUserServiceServer(server, svc)
	reflection.Register(server)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()

	//err = server.Serve(listener)
	//if err != nil {
	//	// todo
	//}
}

func startHTTPServer() {
	ginRouter := gin.Default()

	router.ApplicationV1Router(ginRouter)

	if err := ginRouter.Run(fmt.Sprintf(":%v", cfg.Config.RestPort)); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}

func init() {
	if err := cfg.Config.DB.AutoMigrate(&model.User{}); err != nil {
		return
	}
}

func main() {
	startGRPCServer()
	startHTTPServer()
}
