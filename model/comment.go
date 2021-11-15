package model

import (
	"gorm.io/gorm"
	"time"
)

type CommentMessage struct {
	Symbol string
	List   []Comment
}

type Comment struct {
	Id     int    `json:"id"`
	Symbol string `json:"-"`
	UserId int    `json:"user_id"` //用户ID
	Title  string `json:"title"`   //标题
	Source string `json:"source"`  //来源
	Text   string `json:"text"`    //正文
	//Type      string    `json:"type"`
	ViewCount int       `json:"view_count"`
	User      User      `json:"user"`       //User struct
	CreatedAt int64     `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"-"`
}

func CountComment(symbol string, startTime time.Time, endTime time.Time, db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Comment{}).Where("symbol = ? and created_at >= ? and created_at <= ?", symbol, startTime.Unix(), endTime.Unix()).Count(&count).Error
	return count, err
}
