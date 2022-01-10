package v1

import (
	"context"
	"fmt"

	"github.com/kvendingoldo/gu-user-service/internal/models"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	v1.UserServiceServer
}

func (UserServiceServer) GetAll(ctx context.Context, req *emptypb.Empty) (*v1.GetAllResponse, error) {
	var users []models.User

	err := models.GetAllUsers(&users)
	if err != nil {
		fmt.Println(err)
		// TODO
	}

	var response []*v1.User
	for _, user := range users {
		fmt.Println(user)
		//response = append(response, user.GetGRPCModel())
	}

	return &v1.GetAllResponse{Users: response}, nil
}

func (UserServiceServer) GetByID(ctx context.Context, req *v1.GetByIdRequest) (*v1.GetByIdResponse, error) {
	var user models.User

	err := models.GetUserByID(&user, req.Id)
	if err != nil {

		return nil, err
	}
	gRPCResult := user.GetGRPCModel()

	return &v1.GetByIdResponse{User: &gRPCResult}, nil
}

func (UserServiceServer) New(ctx context.Context, req *v1.NewRequest) (*v1.NewResponse, error) {
	user := models.User{
		Name:      req.User.Name,
		Latitude:  req.User.Latitude,
		Longitude: req.User.Longitude,
	}

	err := models.CreateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &v1.NewResponse{Id: user.ID}, nil
}

func (UserServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	//user := models.User{
	//	ID:        req.User.Id,
	//	Name:      req.User.Name,
	//	Latitude:  req.User.Latitude,
	//	Longitude: req.User.Longitude,
	//}
	//
	//user, err := models.UpdateUser(req.User.Id, requestMap)
	//if err != nil {
	//	// TODO
	//	fmt.Println(err)
	//}

	//
	//var err error
	//if req.User.Id == 0 {
	//	err = config.Config.DB.Model(&user).
	//		Where("name = ?", req.User.Name).
	//		Updates(map[string]interface{}{
	//			"latitude":  req.User.Latitude,
	//			"longitude": req.User.Longitude,
	//		}).
	//		Error
	//} else {
	//	err = config.Config.DB.Model(&user).
	//		Updates(&user).
	//		Error
	//}
	//
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
	//var updatedUser models.User
	//
	//if req.User.Id == 0 {
	//	err = config.Config.DB.
	//		Where("name = ?", req.User.Name).
	//		First(&updatedUser).
	//		Error
	//} else {
	//	err = config.Config.DB.
	//		Where("id = ?", req.User.Id).
	//		First(&updatedUser).
	//		Error
	//}
	//
	//if err != nil {
	//	// TODO
	//}

	return &v1.UpdateResponse{Id: 0}, nil
}

func (UserServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	uid := req.Id
	err := models.DeleteUser(uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.DeleteResponse{Id: uid}, nil
}
