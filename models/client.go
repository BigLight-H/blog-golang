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
	Motto string
	HeadImg *HeadImg  `orm:"rel(one)"`  //头像一对一
}

func (m *Client) TableName() string {
	return TableName("client")
}
