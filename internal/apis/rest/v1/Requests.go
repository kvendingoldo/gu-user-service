package v1

type NewUserRequest struct {
	Name      string  `json:"name" example:"Bill" binding:"required"`
	Latitude  float64 `json:"lat" example:"39.12355"`
	Longitude float64 `json:"lon" example:"27.64538"`
}
