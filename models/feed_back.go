/*
@Time : 2019-08-26 10:56
@Author : biglight
@File : client
@Software: GoLand
*/
package models

import "time"

type FeedBack struct {
	Id int
	Email string
	Title string
	Content string
	Reply string
	Created time.Time
}

func (m *FeedBack) TableName() string {
	return TableName("feedback")
}