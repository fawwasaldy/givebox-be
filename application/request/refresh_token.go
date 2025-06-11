package request

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	UserID       string `json:"user_id" binding:"required"`
}
