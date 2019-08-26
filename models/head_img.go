/*
@Time : 2019-08-23 21:29
@Author : biglight
@File : head_img
@Software: GoLand
*/
package models

type HeadImg struct {
	Id  int
	ImgUrl string
}

func (m *HeadImg) TableName() string {
	return TableName("head_img")
}
