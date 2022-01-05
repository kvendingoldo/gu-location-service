package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kvendingoldo/gu-location-service/config"
	"github.com/kvendingoldo/gu-location-service/internal/server/rest"
	"github.com/kvendingoldo/gu-location-service/model"
	"github.com/kvendingoldo/gu-location-service/pkg/logger"
	"log"
)

func startHTTPServer() {
	router := gin.New()
	router.Use(logger.GinLogger(config.Config.Logger), gin.Recovery())
	rest.ApplicationRouter(router)

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
