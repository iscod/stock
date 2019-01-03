// Package xueqiu main

package xueqiu

import "net/http"
import "encoding/json"
import "io"

func GetComments(Symbol string) ([]Comment, error) {
	url := "/statuses/search.json?count=10&comment=0&hl=0&source=all&sort=&page=1&q=&symbol="
	curl := host + url + Symbol
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
