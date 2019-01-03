// XueQiu comment message get and save

package xueqiu

import "io"
import "net/http"
import "encoding/json"

var commentcuri = "/statuses/search.json?count=10&comment=0&hl=0&source=all&sort=&page=1&q=&symbol="

type CommentMessage struct {
	//commentmessage 雪球评论接口返回结构体
	Symbol string
	List   []Comment
}

type User struct {
	//User 用户结构体
	Id          int
	Screen_name string //用户名
	City        string //城市
	Province    string
}

type Comment struct {
	//comment 雪球评论结构体
	Id         int
	User_id    int    //用户ID
	Title      string //标题
	Source     string //来源
	Text       string //正文
	Created_at int64  //创建时间
	User       User   //User struct
}

func Getcomment(symbol string) ([]Comment, error) {
	curl := host + commentcuri + symbol
	cookie, err := GetCookie()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", curl, nil)
	for i := 0; i < len(cookie); i++ {
		req.AddCookie(cookie[i])
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)

	if err != nil {
		return nil, err
	}

	var comments []Comment

	for {
		var m CommentMessage
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		for _, value := range m.List {
			comments = append(comments, value)
		}
	}

	return comments, err
}
