/*
@Time : 2019-08-26 10:56
@Author : biglight
@File : client
@Software: GoLand
*/
package models

type FeedBack struct {
	Id int
	Email string
	Name string
	Content string
	Reply string
	Created string
}

func (m *FeedBack) TableName() string {
	return TableName("feedback")
}