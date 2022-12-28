package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"microservice-pot/controllers"
	docs "microservice-pot/docs"
	"microservice-pot/initializers"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	healthcheck.New(r, healthcheck.DefaultConfig(), []checks.Check{})

	r.GET("/path", controllers.RouteGet)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
