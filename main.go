package main

import (
	"github.com/iscod/stock/chart"
	"github.com/iscod/stock/comment"
	"github.com/iscod/stock/detail"
	"github.com/iscod/stock/model"
	"github.com/iscod/stock/price"
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
			comment.Run(symbol, db)
		}
		for _, symbol := range symbols {
			_ = chart.Run(symbol.Symbol, db)
		}
		time.Sleep(time.Minute * 10)
	}
}
