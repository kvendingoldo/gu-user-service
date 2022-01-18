package v1

import (
	"context"
	"fmt"

	"github.com/kvendingoldo/gu-user-service/internal/models"
	v1 "github.com/kvendingoldo/gu-user-service/pkg/api/kvendingoldo/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserServiceServer struct {
	v1.UserServiceServer
}

func (UserServiceServer) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	user := models.User{
		Name:      req.User.Name,
		Latitude:  req.User.Latitude,
		Longitude: req.User.Longitude,
	}

	fmt.Println("kek")

	err := models.CreateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return user.GetGRPCModel(), nil
}

func (UserServiceServer) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	var user models.User

	err := models.GetUser(&user, req.Name)
	if err != nil {
		return nil, err
	}

	return user.GetGRPCModel(), nil
}

func (UserServiceServer) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	var user models.User
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

	return user.GetGRPCModel(), nil
}

func (UserServiceServer) Delete(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	uid := req.Name
	err := models.DeleteUser(uid)
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (UserServiceServer) List(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	var users []models.User

	err := models.ListUsers(&users)
	if err != nil {
		fmt.Println(err)
		// TODO
	}

	var response []*v1.User
	for _, user := range users {
		response = append(response, user.GetGRPCModel())
	}

	return &v1.ListUsersResponse{Users: response}, nil
}
