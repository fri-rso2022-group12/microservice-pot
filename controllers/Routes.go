package controllers

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "googlemaps.github.io/maps"
    "log"
    "net/http"
    "os"
)

type PathBody struct {
    From string `json:"from" binding:"required" form:"from"`
    To   string `json:"to" binding:"required" form:"to"`
}

// GetPath godoc
// @Summary      Get shortest path
// @Description  Describe shortest path from one location to another
// @Tags         path
// @Accept       json
// @Produce      json
// @Param        from  query  string  true  "Starting location"
// @Param        to  query  string  true  "End location"
// @Router       /path [get]
func RouteGet(c *gin.Context) {

    var body PathBody

    err := c.ShouldBind(&body)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest,
            gin.H{
                "error":   "validation_err",
                "message": err.Error(),
            })
        return
    }

    mapClient, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }
    r := &maps.DirectionsRequest{
        Origin:      body.From,
        Destination: body.To,
    }
    route, _, err := mapClient.Directions(context.Background(), r)
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }

    c.JSON(http.StatusOK, gin.H{
        "route": route,
    })
    fmt.Println(route)
}
