// Package xueque url api helper
package xueqiucurl

import "net/http"
import "encoding/json"
import "io"
import "fmt"

const host string = "https://xueqiu.com"

func GetCookie() ([]*http.Cookie, error) {
	resp, err := http.Get("http://xueqiu.com/s/SZ000651")
	return resp.Cookies(), err
}

// func Status() (string, error) {
// 	resp, err := Init()
// 	return resp.Status, err
// }

func Get(url string) (*http.Response, error) {
	curl := host + url
	cookie, _ := GetCookie()
	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	} else {
		return resp, nil
	}
}

func GetBody() {

}

type Message struct {
	// json消息类型
	Symbol string
	List   []Comment
}

type User struct {
	Id int
	Screen_name string
	City string
	Province string
}

type Comment struct {
	//comment 雪球评论结构体
	//
	// 这是雪球相关Symbol stock 评论返回json结构机字段
	Id      int
	User_id int
	Title   string
	Source  string
	Text    string
	Created_at int64
	User User
}

func GetComments(Symbol string) ([]Comment, error) {
	url := "/statuses/search.json?count=10&comment=0&hl=0&source=all&sort=&page=1&q=&symbol="
	curl := url + Symbol
	rs, err := Get(curl)

	if err != nil {
		return nil, err
	}

	defer rs.Body.Close()

	dec := json.NewDecoder(rs.Body)

	var comments []Comment

	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		for _, value := range m.List {
			comments = append(comments, value)
		}
	}

	return comments, err
}
