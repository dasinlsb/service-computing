package model

import (
	"github.com/dasinlsb/forum-mirror-backend/config"
	"github.com/jinzhu/gorm"
)

func GetArticleWithinRange(l, r int) (articles []Article, err error) {
	db, err := gorm.Open(config.DbName, config.DbArg)
	if err != nil {
		return articles, err
	}
	defer db.Close()
	db.Order("craw_time ASC").Find(&articles)
	return
}