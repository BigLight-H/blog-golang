package models

type Browse struct {
	Id int
	Article  *Article  `orm:"rel(one)"`
	Client   *Client   `orm:"rel(one)"`
}

func (m *Browse) TableName() string {
	return TableName("look")
}