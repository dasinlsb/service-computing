package model

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"strconv"
	"time"
)

func CrawArticleList() (articles []Article, err error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"),
	)
	c.OnHTML("noscript", func(element *colly.HTMLElement) {
		element.DOM.SetHtml(element.DOM.Text())
	})
	c.OnHTML("div[itemprop]", func(element *colly.HTMLElement) {
		element.DOM.Each(func(i int, s *goquery.Selection) {
			floorStr := s.Find("span[class=posts]").First().Text()
			floors, err := strconv.Atoi(floorStr[1:len(floorStr)-1])
			if err != nil {
				floors = -1
			}
			article := Article{
				//Id:       id,
				Title:    s.Find("span[itemprop=name]").First().Text(),
				Floors:   floors,
				//Author:   User{},
				Content:  "",
				PageURL:  s.Find("a").First().AttrOr("href", ""),
				Tag:      s.Find("span[class=category]").ChildrenFiltered("a").Text(),
				CrawTime: time.Now(),
			}
			articles = append(articles, article)
		})
	})
	c.OnResponse(func(response *colly.Response) {

	})
	if err := c.Visit("https://emacs-china.org"); err != nil {
		return []Article{}, err
	}
	return
}

func FetchAllData() error {
	if Db.HasTable(&Article{}) {
		Db.DropTable(&Article{})
	}
	if Db.HasTable(&User{}) {
		Db.DropTable(&User{})
	}
	Db.CreateTable(&Article{})
	Db.CreateTable(&User{})
	// check latest crawling time
	// decide if re-crawling is required
	{
		var latestArticle Article
		Db.Order("craw_time desc").First(&latestArticle)
		// abort if at least 1 article record exists and last crawling time is within 60 minutes ago
		if latestArticle.Title != "" && (time.Now().Sub(latestArticle.CrawTime)).Minutes() < 30 {
			log.Printf("detected exsiting data, last crawled at: %v\n", latestArticle.CrawTime)
			return nil
		}
	}
	log.Println("updating data...")
	articles, err := CrawArticleList()
	if err != nil {
		return err
	}
	//userSet := make(map[string]int)
	//userCnt := 0
	for _, article := range articles {
		Db.Create(&article)
		//_ = userSet
		//_ = userCnt
		//user := article.Author
		//if _, ok := userSet[user.Name]; !ok  {
		//	userCnt += 1
		//	user.Id = userCnt
		//	userSet[user.Name] = userCnt
		//	Db.Create(&user)
		//}
	}
	log.Println("updating data successfully!")
	return nil
}
