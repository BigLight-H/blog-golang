package models

type Comment struct {
	Id int
	//Article
	//Client
	Content string
	Created string
}

func (m *Comment) TableName() string {
	return TableName("comment")
}