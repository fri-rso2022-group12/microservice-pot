package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    healthcheck "github.com/tavsec/gin-healthcheck"
    "github.com/tavsec/gin-healthcheck/checks"
    ginprometheus "github.com/zsais/go-gin-prometheus"
    "microservice-pot/controllers"
    docs "microservice-pot/docs"
    "microservice-pot/initializers"
    "microservice-pot/middlewares"
    "time"
)

func init() {
    initializers.LoadEnvVariables()
    initializers.InitializeConsul()
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

    r.Use(middlewares.MaintenanceMode())
    p := ginprometheus.NewPrometheus("gin")
    p.MetricsPath = "/path-service/metrics"
    p.Use(r)

    config := healthcheck.DefaultConfig()
    config.HealthPath = "/path-service/healthz"
    healthcheck.New(r, healthcheck.DefaultConfig(), []checks.Check{})

    r.GET("/path-service/path", controllers.RouteGet)
    r.GET("/path-service/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    r.Run()
}
