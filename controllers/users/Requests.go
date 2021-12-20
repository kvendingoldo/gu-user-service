package users

type NewUserRequest struct {
	Name        string `json:"name" example:"Bill" binding:"required"`
	Coordinates string `json:"coordinates" example:"39.12355, 27.64538"  binding:"required"`
}
