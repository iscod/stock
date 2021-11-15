package model

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var user = "root"
var passwd = "root"

func InitDb() (*gorm.DB, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/stock?charset=utf8")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(16)
	db.SetMaxOpenConns(32)
	db.SetConnMaxLifetime(time.Second * 60) //db链接超时时间
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return gormDB, err
}
