/*
@Time : 2019-09-14 23:23
@Author : biglight
@File : collection
@Software: GoLand
*/
package models

type Collection struct {
	Id int
	Article *Article `orm:"rel(one)"`
	Client *Client `orm:"rel(one)"`
}

func (m *Collection) TableName() string {
	return TableName("collection")
}
