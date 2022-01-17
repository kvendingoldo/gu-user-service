package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	guLogger "github.com/kvendingoldo/gu-common/pkg/logger"
	"github.com/kvendingoldo/gu-user-service/config"
	v1Grpc "github.com/kvendingoldo/gu-user-service/internal/apis/grpc/v1"
	"github.com/kvendingoldo/gu-user-service/internal/models"
	v1 "github.com/kvendingoldo/gu-user-service/pkg/user_api/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"time"
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

func SwagDoc(c *gin.Context) {
	schemaPath := "static/api.swagger.json"
	fInfo, _ := os.Stat(schemaPath)
	data := map[string]string{
		"EnvName":    "ab",
		"AppName":    "cd",
		"JsonFile":   fmt.Sprintf("/%s", schemaPath),
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(time.RFC3339),
	}
	c.HTML(200, "swagger.tpl", data)
}

func startHTTPServer() {
	host := fmt.Sprintf(":%v", config.Config.RestPort)

	gwmux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			return metadata.Pairs("tracing", request.Header.Get("tracing"))
		}),
	)

	opt := []grpc.DialOption{grpc.WithInsecure()}

	err := v1.RegisterUserServiceHandlerFromEndpoint(context.Background(), gwmux, ":9092", opt)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	//router.RedirectTrailingSlash = false
	router.Use(guLogger.GinLogger(config.Config.Logger), gin.Recovery())

	//router.GET("/", func(c *gin.Context) {
	//	c.Redirect(http.StatusFound, "/swagger-ui/")
	//})

	router.LoadHTMLFiles("static/views/swagger.tpl")
	router.GET("/swagger-ui/", SwagDoc)
	router.Static("/static", "./static")

	router.GET("/api/*any", func(c *gin.Context) {
		c.Request.Header.Set("tracing", "ing")
	}, func(c *gin.Context) {
		gwmux.ServeHTTP(c.Writer, c.Request)
	})

	if err := router.Run(host); err != nil {
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
