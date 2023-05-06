package model

type SaveSocialMediaRequest struct {
	ID             uint   `json:"-"`
	UserId         uint   `json:"-"`
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

type SocialMediaResponse struct {
	ID             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}
