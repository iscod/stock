package model

type CommentMessage struct {
	Symbol string
	List   []Comment
}

type Comment struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`    //用户ID
	Title     string `json:"title"`      //标题
	Source    string `json:"source"`     //来源
	Text      string `json:"text"`       //正文
	CreatedAt int64  `json:"created_at"` //创建时间
	Type      string `json:"type"`
	User      User   `json:"user"` //User struct
}
