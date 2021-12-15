package models

import (
	"github.com/kvendingoldo/gu-user-service/config"
	"time"
)

type User struct {
	ID        int       `json:"id" example:"123" gorm:"primaryKey"`
	Name      string    `json:"name" example:"Steven" gorm:"unique"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	err = config.Config.DB.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}
