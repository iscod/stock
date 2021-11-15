package model

type User struct {
	Id         int    `json:"id"`
	ScreenName string `json:"screen_name"` //用户名
	City       string `json:"city"`        //城市
	Province   string `json:"province"`
}
