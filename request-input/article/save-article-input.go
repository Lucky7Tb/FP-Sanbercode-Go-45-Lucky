package article

type SaveArticleInput struct {
	Content string `json:"content" binding:"required"`
}
