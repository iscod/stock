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

		for _, symbol := range symbols {
			now := time.Now()
			go run(symbol.Symbol, db)
			if now.Hour() == 16{
				go func() {
					chart.Run(symbol.Symbol, db)
				}()

			}
			if time.Now().Hour() == 1 {
				go func() {
					price.Run(symbol.Symbol, db)
				}()
			}
		}
		time.Sleep(time.Minute * 10)
	}
}

func run(symbol string, db *gorm.DB) {
	fmt.Printf("名称: %s\n", symbol)

	//获取价格信息
	quote, err := base.GetQuote(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("当前价格: %f, 开盘价: %f, 均价: %f, 当天最低: %f, 当天最高: %f\n", quote.Current, quote.Open, quote.AvgPrice, quote.Low, quote.High)

	//获取评论
	comments, err := base.GetComment(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	//保存评论
	for _, comment := range comments {
		if comment.Id != 0 {
			fmt.Printf("New Comment Username : %s,title: %s, Time: ,  %\n", comment.User.ScreenName, comment.Title)
		}
		comment.Symbol = symbol
		comment.CreatedAt = comment.CreatedAt / 1000
		err := db.Save(&comment).Error
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
