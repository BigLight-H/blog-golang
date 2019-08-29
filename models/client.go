/*
@Time : 2019-08-25 00:15
@Author : biglight
@File : client
@Software: GoLand
*/
package models

type Client struct {
	Id int
	Username string
	Password string
	Email string
	Mobile string
	Sex int
	Age int
	HeadImg *HeadImg  `orm:"rel(one)"`  //头像一对一
	Comment []*Comment `orm:"reverse(many)"`
}

func (m *Client) TableName() string {
	return TableName("client")
}
