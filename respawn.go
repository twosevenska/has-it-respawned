package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"has-it-respawned/clients/steampowered"
	"has-it-respawned/controllers"
)

// ContextParams holds the objects required
type ContextParams struct {
	SteamClient steampowered.Client
}

func main() {
	flag.Parse()

	args := flag.Args()
	k := os.Getenv("RESPAWN_STEAM_KEY")
	if k == "" {
		if len(args) != 1 {
			log.Fatal("Please provide an API key for Steam WEB API.")
		}
		k = args[0]
	}

	contextParams := ContextParams{
		SteamClient: steampowered.New(k),
	}

	r := gin.Default()
	r.Use(ContextObjects(&contextParams))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	testAPI := r.Group("/test")
	{
		testAPI.GET("/steam", controllers.SteamTest)
	}

	r.Run(":7000") // listen and serve on 0.0.0.0:7000
}

// ContextObjects attaches backend clients to the API context
func ContextObjects(contextParams *ContextParams) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("steamClient", contextParams.SteamClient)
		c.Next()
	}
}
