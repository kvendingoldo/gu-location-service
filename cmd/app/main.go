package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	guLogger "github.com/kvendingoldo/gu-common/pkg/logger"
	"github.com/kvendingoldo/gu-location-service/config"
	v1 "github.com/kvendingoldo/gu-location-service/internal/controller/rest/v1"
	"github.com/kvendingoldo/gu-location-service/internal/model"
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
	model.Setup()
}

func main() {
	startHTTPServer()
}
