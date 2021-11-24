package main

import (
	"github.com/iscod/stock/base"
	"github.com/iscod/stock/chart"
	"github.com/iscod/stock/detail"
	"github.com/iscod/stock/model"
	"github.com/iscod/stock/price"
	"gorm.io/gorm"
	"time"
)
import "fmt"

const host string = "https://xueqiu.com"
const stockhost string = "https://stock.xueqiu.com"

func main() {
	for true {
		db, err := model.InitDb()
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		var symbols []*model.Symbol
		err = db.Where("status > ?", 0).Find(&symbols).Error

		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		if time.Now().Hour() <= 16 && time.Now().Hour() >= 9 {
			for _, symbol := range symbols {
				price.Run(symbol.Symbol, db)

				err = detail.Run(symbol.Symbol, db)
				if err != nil {
					fmt.Printf("GetDetail Error %s", err)
				}
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
			fmt.Printf("New Comment Username : %s,title: %s, Time: %s\n", comment.User.ScreenName, comment.Title, time.Unix(comment.CreatedAt/1000, 0).Format("2006-01-02 15:04:05"))
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
