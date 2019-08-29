package models

type Comment struct {
	Id int
	Article *Article `orm:"rel(fk)"`
	Client *Client `orm:"rel(fk)"`
	Content string
	Created string
}

func (m *Comment) TableName() string {
	return TableName("comment")
}