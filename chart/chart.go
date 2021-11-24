package chart

import (
	"fmt"
	"github.com/iscod/stock/model"
	"gorm.io/gorm"
	"time"
)

func Run(symbol string, db *gorm.DB) error {
	var quotes []*model.Quote
	err := db.Where(model.Quote{Symbol: symbol}).Where("exec_at >= ?", time.Now().AddDate(0, 0, -1).Format("2006-01-02")).Find(&quotes).Error

	if err != nil {
		return err
	}

	for _, v := range quotes {
		startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v.ExecAt+" 00:00:00", time.Local)
		endTime := startTime.AddDate(0, 0, 1)
		c, err := model.CountComment(symbol, startTime, endTime, -1, db)
		if err != nil {
			fmt.Printf("Comment err : %s", err)
			continue
		}
		v.CommentCount = c
		c3, err := model.CountComment(symbol, startTime, endTime, 3, db)
		v.CommentCount3 = c3

		fmt.Printf("%s %v\t评论数量: %d, 机构: %d\n", v.Name, v.ExecAt, c, c3)

		err = db.Updates(v).Error
		if err != nil {
			fmt.Printf("Error : %s", err)
			continue
		}
	}
	return nil
}
