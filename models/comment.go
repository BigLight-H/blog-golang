package models

type Comment struct {
	Id int64
	Article *Article `orm:"rel(one)"`
	Client *Client `orm:"rel(one)"`
	Content string
	Created string
	Path string
}

func (m *Comment) TableName() string {
	return TableName("comment")
}