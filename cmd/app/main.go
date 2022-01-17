package main

import (
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-user-service/api"
	"github.com/kvendingoldo/gu-user-service/config"
	v1Grpc "github.com/kvendingoldo/gu-user-service/internal/apis/grpc/v1"
	"github.com/kvendingoldo/gu-user-service/internal/models"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
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

func NewGinPetServer(petStore *api.PetStore, port int) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic chi router
	r := gin.Default()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	r = api.RegisterHandlers(r, petStore)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
	return s
}

func startHTTPServer() {
	petStore := api.NewPetStore()
	s := NewGinPetServer(petStore, 8082)
	log.Fatal(s.ListenAndServe())

	//router := gin.New()
	//router.RedirectTrailingSlash = false
	//router.Use(guLogger.GinLogger(config.Config.Logger), gin.Recovery())
	//
	//v1Rest.NewRouter(router)
	//v2Rest.NewRouter(router)
	//
	//if err := router.Run(fmt.Sprintf(":%v", config.Config.RestPort)); err != nil {
	//	log.Fatalf("could not start http server: %v", err)
	//}
}

func init() {
	err := config.Setup()
	if err != nil {
		fmt.Println(err)
	}
	models.Setup()
}

func main() {
	//startGRPCServer()
	startHTTPServer()
}
