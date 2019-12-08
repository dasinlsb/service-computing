package router

import (
	"github.com/dasinlsb/forum-mirror-backend/handler"
	"github.com/dasinlsb/forum-mirror-backend/middleware"
	"github.com/dasinlsb/forum-mirror-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Route(r *gin.Engine) {
	// need previous authentation
	api := r.Group("/api")
	authorized := api.Group("/auth")
	authorized.Use(middleware.Auth)
	authorized.Use(func(c *gin.Context) {
		if err := model.FetchAllData(); err != nil {
			log.Println("fetch data failed: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
	})
	{
		authorized.GET("/validate", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{})
		})
		authorized.GET("/article", handler.GetArticleList)
	}

	api.Group("/")
	{
		api.POST("/login", middleware.LoginHandler)
	}
}
