package model

import (
	"encoding/json"
	"fmt"
	"github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/model/errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int       `json:"id" example:"123" gorm:"primaryKey"`
	Name        string    `json:"name" example:"Steven" gorm:"unique"`
	Coordinates string    `json:"coordinates" example:"39.12355, 27.64538"`
	CreatedAt   time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	err = config.Config.DB.Create(user).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return err
		}

		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
			return

		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
	}

	return
}

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	err = config.Config.DB.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id int) (err error) {
	err = config.Config.DB.Where("id = ?", id).First(user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			fmt.Println("todo")
			//err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
		default:
			fmt.Println("todo")
			//err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	return
}

// UpdateUser ... Update user
func UpdateUser(id int, userMap map[string]interface{}) (user User, err error) {
	user.ID = id
	err = config.Config.DB.Model(&user).
		Select("name", "coordinates").
		Updates(userMap).Error

	// err = config.DB.Save(medicine).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			//err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			return

		default:
			//err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	err = config.Config.DB.Where("id = ?", id).First(&user).Error
	return
}

// DeleteUser ... Delete user
func DeleteUser(id int) (err error) {
	tx := config.Config.DB.Delete(&User{}, id)
	if tx.Error != nil {
		//err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		//err = modelErrors.NewAppErrorWithType(modelErrors.NotFound)
	}

	return
}
