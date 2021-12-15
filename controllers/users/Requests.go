package users

type NewUserRequest struct {
	Name string `json:"name" example:"Paracetamol" gorm:"unique" binding:"required"`
}
