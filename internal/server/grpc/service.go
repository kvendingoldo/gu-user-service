package grpc

import (
	"context"
	"github.com/kvendingoldo/gu-user-service/config"
	"github.com/kvendingoldo/gu-user-service/model"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// TODO

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

	// TODO: return real ID
	return &v1.CreateResponse{Api: "v1", Id: int64(user.ID)}, nil
}

func (UserServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// TODO: panic: reflect: reflect.Value.SetInt using unaddressable value
	user := model.User{}
	err := config.Config.DB.Where("id = ?", req.Id).First(user).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	// TODO
	resUser := &v1.User{
		Id:          int64(user.ID),
		Name:        user.Name,
		Coordinates: user.Coordinates,
	}

	return &v1.ReadResponse{Api: "v1", User: resUser}, nil
}
func (UserServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// TODO
	//user := model.User{
	//	ID: int(req.User.Id),
	//}

	//err := config.Config.DB.Model(&user).
	//	Select("name", "coordinates").
	//	Updates(userMap).Error
	//
	//err = config.Config.DB.Save(user).Error
	//if err != nil {
	//	byteErr, _ := json.Marshal(err)
	//	var newError modelErrors.GormErr
	//	err = json.Unmarshal(byteErr, &newError)
	//	if err != nil {
	//		// TODO
	//		//return
	//	}
	//	switch newError.Number {
	//	case 1062:
	//		err = modelErrors.NewAppErrorWithType(modelErrors.ResourceAlreadyExists)
	//		// TODO
	//		//return
	//	default:
	//		err = modelErrors.NewAppErrorWithType(modelErrors.UnknownError)
	//	}
	//}
	//
	//err = config.Config.DB.Where("id = ?", req.User.Id).First(&user).Error
	//if err != nil {
	//	// TODO
	//}

	return &v1.UpdateResponse{Api: "v1", Updated: int64(req.User.Id)}, nil

}

func (UserServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	userID := int(req.Id)
	err := model.DeleteUser(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.DeleteResponse{Api: "v1", Deleted: req.Id}, nil
}
