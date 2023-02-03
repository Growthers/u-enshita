package api

type CreateCommentRequestJSON struct {
	From        string `json:"from"`
	DataSource  string `json:"dataSource"`
	CommentText string `json:"commentText"`
}

type CreateCommentWebSocketJSON struct {
	From        string `json:"from"`
	DataSource  string `json:"dataSource"`
	CommentText string `json:"commentText"`
}
