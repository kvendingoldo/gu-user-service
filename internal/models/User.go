package models

import (
	"encoding/json"
	cValidation "github.com/kvendingoldo/gu-common/pkg/validation"

	appErrors "github.com/kvendingoldo/gu-common/pkg/errors"
	"github.com/kvendingoldo/gu-user-service/config"

	v1 "github.com/kvendingoldo/gu-user-service/gen/go/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64     `json:"id" example:"23" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`

	Name      string  `json:"name" example:"Steven" gorm:"unique"`
	Latitude  float64 `json:"lat" example:"39.12355"`
	Longitude float64 `json:"lon" example:"27.64538"`
}

// TableName represents name of SQL table, used by GORM
func (u *User) TableName() string {
	return "users"
}

func (u *User) GetGRPCModel() v1.User {
	return v1.User{
		Id:        u.ID,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		Name:      u.Name,
		Latitude:  u.Latitude,
		Longitude: u.Longitude,
	}
}

func (u *User) From(gRPCModel *v1.User) {
	u.ID = gRPCModel.Id
	u.CreatedAt = gRPCModel.CreatedAt.AsTime()
	u.UpdatedAt = gRPCModel.UpdatedAt.AsTime()
	u.Name = gRPCModel.Name
	u.Latitude = gRPCModel.Latitude
	u.Longitude = gRPCModel.Longitude
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	err = cValidation.ValidateUsername(user.Name)
	if err != nil {
		return err
	}
	err = cValidation.ValidateCoordinates(user.Latitude, user.Longitude)
	if err != nil {
		return err
	}

	err = config.Config.DB.Create(user).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError appErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return err
		}

		switch newError.Number {
		case 1062:
			err = appErrors.NewAppErrorWithType(appErrors.ResourceAlreadyExists)
			return
		default:
			err = appErrors.NewAppErrorWithType(appErrors.UnknownError)
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
func GetUserByID(user *User, id int64) (err error) {
	if id == 0 {
		err = appErrors.NewAppErrorWithType(appErrors.NotFound)
		return
	}

	err = config.Config.DB.Where("id = ?", id).First(user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = appErrors.NewAppErrorWithType(appErrors.NotFound)
		default:
			err = appErrors.NewAppErrorWithType(appErrors.UnknownError)
		}
	}

	return
}

// UpdateUser ... Update user
func UpdateUser(id int64, userMap map[string]interface{}) (user User, err error) {
	//err = cValidation.ValidateUID(user.ID)
	//if err != nil {
	//	return err
	//}
	//err = cValidation.ValidateUsername(user.Name)
	//if err != nil {
	//	return err
	//}
	//err = cValidation.ValidateCoordinates(user.Latitude, user.Longitude)
	//if err != nil {
	//	return err
	//}

	// TODO: check exist

	user.ID = id
	err = config.Config.DB.Model(&user).
		Select("name", "coordinates").
		Updates(userMap).Error

	err = config.Config.DB.Save(user).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError appErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = appErrors.NewAppErrorWithType(appErrors.ResourceAlreadyExists)
			return
		default:
			err = appErrors.NewAppErrorWithType(appErrors.UnknownError)
		}
	}

	err = config.Config.DB.Where("id = ?", id).First(&user).Error
	return
}

// DeleteUser ... Delete user
func DeleteUser(id int64) (err error) {
	tx := config.Config.DB.Delete(&User{}, id)
	if tx.Error != nil {
		//err = appErrors.NewAppErrorWithType(appErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = appErrors.NewAppErrorWithType(appErrors.NotFound)
	}

	return
}
