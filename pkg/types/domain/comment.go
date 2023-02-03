package domain

type Comment struct {
	From        string
	DataSource  string
	CommentText string
}

func NewComment(from string, dataSource string, commentText string) *Comment {
	return &Comment{From: from, DataSource: dataSource, CommentText: commentText}
}
