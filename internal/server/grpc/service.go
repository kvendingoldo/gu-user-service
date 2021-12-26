package grpc

import (
	"context"
	"encoding/json"
	"github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/model"
	modelErrors "github.com/kvendingoldo/gu-user-service/model/errors"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	v1.UserServiceServer
}

func (UserServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	user := model.User{
		Name:        req.User.Name,
		Coordinates: req.User.Coordinates,
	}

	err := model.CreateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &v1.CreateResponse{Id: user.ID}, nil
}

func (UserServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	var user model.User
	err := config.Config.DB.Where("id = ?", req.Id).First(&user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	resUser := &v1.User{
		Id:          user.ID,
		Name:        user.Name,
		Coordinates: user.Coordinates,
	}

	return &v1.ReadResponse{User: resUser}, nil
}
func (UserServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	user := model.User{
		ID:          req.User.Id,
		Name:        req.User.Name,
		Coordinates: req.User.Coordinates,
	}

	var err error
	if req.User.Id == 0 {
		err = config.Config.DB.Model(&user).
			Where("name = ?", req.User.Name).
			Updates(map[string]interface{}{"coordinates": req.User.Coordinates}).
			Error
	} else {
		err = config.Config.DB.Model(&user).
			Updates(&user).
			Error
	}

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError modelErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			// TODO
			//return
		}
		switch newError.Number {
		case 1062:
			err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
			// TODO
			//return
		default:
			err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
		}
	}

	var updatedUser model.User

	if req.User.Id == 0 {
		err = config.Config.DB.
			Where("name = ?", req.User.Name).
			First(&updatedUser).
			Error
	} else {
		err = config.Config.DB.
			Where("id = ?", req.User.Id).
			First(&updatedUser).
			Error
	}

	if err != nil {
		// TODO
	}

	return &v1.UpdateResponse{Updated: updatedUser.ID}, nil
}

func (UserServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	userID := req.Id
	err := model.DeleteUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.DeleteResponse{Deleted: userID}, nil
}
