package models

type Comment struct {
	Id int
	Article *Article `orm:"rel(one)"`
	Client *Client `orm:"rel(one)"`
	Content string
	Created string
}

func (m *Comment) TableName() string {
	return TableName("comment")
}