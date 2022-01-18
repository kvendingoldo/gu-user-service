package models

import (
	"encoding/json"
	cValidation "github.com/kvendingoldo/gu-common/pkg/validation"

	appErrors "github.com/kvendingoldo/gu-common/pkg/errors"
	"github.com/kvendingoldo/gu-user-service/config"

	v1 "github.com/kvendingoldo/gu-user-service/pkg/api/kvendingoldo/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name      string  `json:"name" example:"Steven" gorm:"primaryKey"`
	Latitude  float64 `json:"lat" example:"39.12355"`
	Longitude float64 `json:"lon" example:"27.64538"`

	CreateTime time.Time `json:"create_time,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdateTime time.Time `json:"update_time,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}

// TableName represents name of SQL table, used by GORM
func (u *User) TableName() string {
	return "users"
}

func (u *User) GetGRPCModel() *v1.User {
	return &v1.User{
		CreateTime: timestamppb.New(u.CreateTime),
		UpdateTime: timestamppb.New(u.UpdateTime),
		Name:       u.Name,
		Latitude:   u.Latitude,
		Longitude:  u.Longitude,
	}
}

func (u *User) From(gRPCModel *v1.User) {
	u.CreateTime = gRPCModel.CreateTime.AsTime()
	u.UpdateTime = gRPCModel.UpdateTime.AsTime()
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

// ListUsers Fetch all users
func ListUsers(user *[]User) (err error) {
	err = config.Config.DB.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByName ... Fetch only one user by Id
func GetUser(user *User, name string) (err error) {
	err = config.Config.DB.Where("name = ?", name).First(user).Error

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
func DeleteUser(name string) (err error) {
	tx := config.Config.DB.Delete(&User{}, name)
	if tx.Error != nil {
		//err = appErrors.NewAppErrorWithType(appErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = appErrors.NewAppErrorWithType(appErrors.NotFound)
	}

	return
}
