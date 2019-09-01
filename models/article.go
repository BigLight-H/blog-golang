/*
@Time : 2019-08-25 00:10
@Author : biglight
@File : article
@Software: GoLand
*/
package models

type Article struct {
	Id int
	Client *Client `orm:"rel(one)"`
	Type *Type  `orm:"rel(one)"`
	Title string
	Description string
	Tags string
	Status int
	Review int
	ClickVolume int
	CommentNum int
	Cause string
	Content string
	Created string
	Picture string
}

func (m *Article) TableName() string {
	return TableName("article")
}
