package main

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/models"
	v1 "github.com/kvendingoldo/gu-user-service/pkg/api/proto/v1"
	"github.com/kvendingoldo/gu-user-service/pkg/service"
	"github.com/kvendingoldo/gu-user-service/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	server := grpc.NewServer()
	svc := &service.UserServiceServer{}
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
	router := gin.Default()

	routes.ApplicationV1Router(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}

func init() {
	if err := cfg.Config.DB.AutoMigrate(&models.User{}); err != nil {
		return
	}

}

func main() {
	//go
	startGRPCServer()
	startHTTPServer()

	//sig := make(chan os.Signal)
	//signal.Notify(sig, os.Interrupt, os.Kill)
	//<-sig
}
