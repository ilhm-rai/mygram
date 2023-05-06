package model

type SaveCommentRequest struct {
	ID      uint   `json:"-"`
	UserId  uint   `json:"-"`
	PhotoId uint   `json:"photo_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	UserId  uint   `json:"user_id"`
	PhotoId uint   `json:"photo_id"`
	Message string `json:"message"`
}
