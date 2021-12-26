package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cfg "github.com/kvendingoldo/gu-location-service/config"
	"github.com/kvendingoldo/gu-location-service/internal/server/rest"
	"github.com/kvendingoldo/gu-location-service/model"
	"log"
)

func startHTTPServer() {
	ginRouter := gin.Default()

	rest.ApplicationRouter(ginRouter)

	if err := ginRouter.Run(fmt.Sprintf(":%v", cfg.Config.RestPort)); err != nil {
		log.Fatalf("could not start http server: %v", err)
	}
}

func init() {
	if err := cfg.Config.DB.AutoMigrate(&model.Location{}); err != nil {
		return
	}
}

func main() {
	startHTTPServer()
}
