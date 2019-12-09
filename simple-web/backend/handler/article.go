package handler

import (
	"github.com/dasinlsb/forum-mirror-backend/config"
	"github.com/dasinlsb/forum-mirror-backend/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	pageStr, ok:= c.GetQuery("page")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unrecognized query params",
		})
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Println("parse page number failed, string:", pageStr, "error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unrecognized query params",
		})
		return
	}
	if page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "page number should be greater than 0",
		})
		return
	}
	log.Println("will return article data in page", page)
	l := (page-1)*config.PostPerPage
	r := page*config.PostPerPage
	articles, err := model.GetArticleWithinRange(l, r)
	if err != nil {
		log.Println("get article data from database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	type Article struct {
		Title string `json:"title"`
		Content string `json:"content"`
		Floors int `json:"floors"`
		Url string `json:"url"`
	}
	var data []Article
	for _, article := range articles {
		data = append(data, Article{
			Title:   article.Title,
			Content: "",
			Floors: article.Floors,
			Url:article.PageURL,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func FakeGetArticleList(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "success",
		"data": []gin.H{
			{
				"title":   "fake-title-1",
				"content": "fake-content-1",
			},
			{
				"title":   "fake-title-2",
				"content": "fake-content-2",
			},
		},
	})
}