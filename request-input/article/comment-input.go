package article

type CommentInput struct {
	UserID  int    `json:"-"`
	Comment string `json:"comment" binding:"required"`
}
