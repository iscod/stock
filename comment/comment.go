package comment

import (
	"fmt"
	"github.com/iscod/stock/base"
	"gorm.io/gorm"
	"time"
)

func Run(symbol string, db *gorm.DB) {
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
