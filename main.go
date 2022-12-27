package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"microservice-pot/controllers"
	docs "microservice-pot/docs"
	"microservice-pot/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	healthcheck.New(r, healthcheck.DefaultConfig(), []checks.Check{})

	r.GET("/path", controllers.RouteGet)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
