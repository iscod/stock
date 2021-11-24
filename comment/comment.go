package comment

import (
	"fmt"
	"github.com/iscod/stock/base"
	"github.com/iscod/stock/model"
	"gorm.io/gorm"
	"time"
)

func Run(symbol *model.Symbol, db *gorm.DB) {
	//获取评论
	comments, err := base.GetComment(symbol.Symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	//保存评论
	for _, comment := range comments {
		if comment.Id != 0 {
			fmt.Printf(" %s 新评论: 用户名: %s, 标题: %s, Time: %s\n", symbol.Name, comment.User.ScreenName, comment.Title, time.Unix(comment.CreatedAt/1000, 0).Format("2006-01-02 15:04:05"))
		}
		if comment.UserId < 0 {
			comment.UserId = 0
		}
		comment.Symbol = symbol.Symbol
		comment.CreatedAt = comment.CreatedAt / 1000
		err := db.Save(&comment).Error
		if err != nil {
			fmt.Printf("%s", err)
		}
	}
}
