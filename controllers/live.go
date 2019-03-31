package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"has-it-respawned/clients/steampowered"
)

// SteamTest fetches the list of games for a user
func SteamTest(c *gin.Context) {
	sc := c.MustGet("steamClient").(steampowered.Client)

	id, found := c.GetQuery("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id"})
		return
	}

	response, err := sc.GetGames(id)
	if err != nil {
		msg := fmt.Sprintf("Failed to fetch data: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		return
	}
	c.JSON(http.StatusOK, response)
}
