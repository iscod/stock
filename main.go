package main

import (
	"github.com/iscod/goStock/base"
	"github.com/iscod/goStock/chart"
	"github.com/iscod/goStock/model"
	"github.com/iscod/goStock/price"
	"gorm.io/gorm"
	"time"
)
import "fmt"

const host string = "https://xueqiu.com"
const stockhost string = "https://stock.xueqiu.com"

func main() {
	db, err := model.InitDb()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	for true {
		var symbols []*model.Symbol
		err = db.Where("status > ?", 0).Find(&symbols).Error

		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		if time.Now().Hour() >= 15 {
			for _, symbol := range symbols {
				price.Run(symbol.Symbol, db)
			}
		}

		for _, symbol := range symbols {
			//now := time.Now()
			run(symbol.Symbol, db)
			chart.Run(symbol.Symbol, db)
		}
		time.Sleep(time.Minute * 10)
	}
}

func run(symbol string, db *gorm.DB) {
	//获取评论
	comments, err := base.GetComment(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	//保存评论
	for _, comment := range comments {
		if comment.Id != 0 {
			fmt.Printf("New Comment Username : %s,title: %s, Time: %s\n", comment.User.ScreenName, comment.Title, time.Unix(comment.CreatedAt, 0).Format("2006-01-02 15:04:05"))
		}
		if comment.UserId < 0 {
			comment.UserId = 0
		}
		comment.Symbol = symbol
		comment.CreatedAt = comment.CreatedAt / 1000
		err := db.Save(&comment).Error
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
