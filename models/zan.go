/*
@Time : 2019-09-14 23:31
@Author : biglight
@File : zan
@Software: GoLand
*/
package models

type Zan struct {
	Id int
	Article *Article `orm:"rel(one)"`
	Client *Client `orm:"rel(one)"`
}

func (m *Zan) TableName() string {
	return TableName("zan")
}
