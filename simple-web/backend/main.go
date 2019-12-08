package main

import (
	"github.com/dasinlsb/forum-mirror-backend/config"
	"github.com/dasinlsb/forum-mirror-backend/model"
	"github.com/dasinlsb/forum-mirror-backend/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		hp := strings.Split(os.Args[1], ":")
		if len(hp) != 2 {
			log.Fatalln("usage: go run main.go [(none)|host port|host:port]")
		}
		config.SetDbArg(hp[0], hp[1])
	} else if len(os.Args) == 3 {
		config.SetDbArg(os.Args[1], os.Args[2])
	}
	model.Connect()
	defer model.Close()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	r.Use(cors.Default())

	router.Route(r)
	log.Println("Server is listening on port 8080...")
	_ = r.Run(":8080")
}
