package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	guLogger "github.com/kvendingoldo/gu-common/pkg/logger"
	"github.com/kvendingoldo/gu-location-service/config"
	v1 "github.com/kvendingoldo/gu-location-service/internal/apis/rest/v1"
	"github.com/kvendingoldo/gu-location-service/internal/models"
	proto "github.com/kvendingoldo/gu-location-service/proto_gen/api"
	"google.golang.org/grpc"
	"log"
)

func startHTTPServer() {
	router := gin.New()
	router.Use(guLogger.GinLogger(config.Config.Logger), gin.Recovery())
	v1.NewRouter(router)

	if err := router.Run(fmt.Sprintf(":%v", config.Config.RestPort)); err != nil {
		log.Fatalf("Could not start HTTP server: %v", err)
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
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	_ = proto.NewUserServiceClient(conn)

	startHTTPServer()
}
