package main

import (
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
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

    r.GET("/path", controllers.RouteGet)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    r.Run()
}
