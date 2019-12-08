package model

import (
	"github.com/dasinlsb/forum-mirror-backend/config"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	Id int
	Name string
	AvatarURL string
	PageURL string
}

type Article struct {
	Title string
	//Author User
	Floors int
	Content string
	PageURL string
	Tag string
	CrawTime time.Time
}

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open(config.DbName, config.DbArg)
	if err != nil {
		log.Fatalln("database connection error: ", err)
	}
	Db = db
}

func Close() {
	Db.Close()
}