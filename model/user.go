package model

import (
	"database/sql/driver"
	"encoding/json"
)

type User struct {
	Id         int    `json:"id"`
	ScreenName string `json:"screen_name"` //用户名
	City       string `json:"city"`        //城市
	Province   string `json:"province"`
}

func (this *User) Scan(value interface{}) error {
	*this = User{}
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), this)
}

func (this User) Value() (driver.Value, error) {
	return json.Marshal(this)
}
