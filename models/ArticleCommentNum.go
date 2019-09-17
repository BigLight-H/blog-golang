package models

type ArticleCommentNum struct {
	ClientId int
	ArticleId int
	Number int
}

func (m *ArticleCommentNum) TableName() string {
	return TableName("article_comment_num")
}
