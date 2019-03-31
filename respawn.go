package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"has-it-respawned/clients/steampowered"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test/steam", func(c *gin.Context) {
		id, found := c.GetQuery("id")
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id"})
			return
		}

		sc := steampowered.New()
		response, err := sc.GetGames(id)
		if err != nil {
			msg := fmt.Sprintf("Failed to fetch data: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
			return
		}

		c.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
