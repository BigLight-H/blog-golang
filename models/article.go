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
	Author string
	Type *Type  `orm:"rel(one)"`
	Title string
	Description string
	Status int
	Review int
	ClickVolume int
	Cause string
	Content string
	Created string
	Comment []*Comment `orm:"reverse(many)"`
}

func (m *Article) TableName() string {
	return TableName("article")
}
