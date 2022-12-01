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
	From string `json:"from" binding:"required"`
	To   string `json:"to" binding:"required"`
}

// GetPath godoc
// @Summary      Get shortest path
// @Description  add by json user
// @Tags         path
// @Accept       json
// @Produce      json
// @Param        path  body      PathBody  true  "Get path"
// @Success      200      {object}  models.User
// @Router       /path [get]
func RouteGet(c *gin.Context) {

	var body PathBody

	err := c.ShouldBindJSON(&body)
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
