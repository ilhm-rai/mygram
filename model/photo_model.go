package model

type SavePhotoRequest struct {
	ID       uint   `json:"-"`
	UserId   uint   `json:"-"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required,url"`
}

type PhotoResponse struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}
