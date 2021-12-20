package grpc

import (
	"context"
	"github.com/kvendingoldo/gu-user-service/model"
	v1 "github.com/kvendingoldo/gu-user-service/proto_gen/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	v1.UserServiceServer
}

func (UserServiceServer) Create(context.Context, *v1.CreateRequest) (*v1.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (UserServiceServer) Read(context.Context, *v1.ReadRequest) (*v1.ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UserServiceServer) Update(context.Context, *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UserServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	medicineID := int(req.Id)
	err := model.DeleteUser(medicineID)
	if err != nil {
		// TODO
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &v1.DeleteResponse{Api: "v1", Deleted: req.Id}, nil
}
