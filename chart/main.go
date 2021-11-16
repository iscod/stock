package main

import (
	"fmt"
	"github.com/iscod/goStock/model"
	"gorm.io/gorm"
	"time"
)

func main() {

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

	for _, symbol := range symbols {
		err = run(symbol.Symbol, db)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}
}

func run(symbol string, db *gorm.DB) error {
	fmt.Printf("名称: %s\n", symbol)
	var quotes []*model.Quote

	err := db.Where(model.Quote{Symbol: symbol, ExecAt: time.Now().AddDate(0, 0, -1).Format("2006-01-02")}).Find(&quotes).Error

	if err != nil {
		return err
	}

	for _, v := range quotes {
		startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v.ExecAt+" 00:00:00", time.Local)
		endTime := startTime.AddDate(0, 0, 1)
		c, err := model.CountComment(symbol, startTime, endTime, db)
		if err != nil {
			fmt.Printf("Comment err : %s", err)
			continue
		}
		v.CommentCount = c
		err = db.Updates(v).Error
		if err != nil {
			fmt.Printf("Error : %s", err)
			continue
		}
	}
	return nil
}
