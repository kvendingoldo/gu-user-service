// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/kvendingoldo/user/v1/service.proto

package user_api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_kvendingoldo_user_v1_service_proto protoreflect.FileDescriptor

var file_api_kvendingoldo_user_v1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c,
	0x64, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x61,
	0x70, 0x69, 0x2f, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x99, 0x07, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x6f, 0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64,
	0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x55,
	0x92, 0x41, 0x37, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x9f, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x4e, 0x92, 0x41, 0x2f, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x08, 0x47, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x47,
	0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x20, 0x6f, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xdb, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c,
	0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x83, 0x01, 0x92, 0x41, 0x59, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x20, 0x6f, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x32, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x9d, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f,
	0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4e, 0x92, 0x41, 0x2f, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a,
	0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xba, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64,
	0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x92, 0x41, 0x41, 0x0a, 0x05, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x1a, 0x1c, 0x47, 0x65, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x23, 0x5a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x76, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x6f, 0x6c, 0x64, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_kvendingoldo_user_v1_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil), // 0: api.kvendingoldo.user.v1.CreateUserRequest
	(*GetUserRequest)(nil),    // 1: api.kvendingoldo.user.v1.GetUserRequest
	(*UpdateUserRequest)(nil), // 2: api.kvendingoldo.user.v1.UpdateUserRequest
	(*DeleteUserRequest)(nil), // 3: api.kvendingoldo.user.v1.DeleteUserRequest
	(*ListUsersRequest)(nil),  // 4: api.kvendingoldo.user.v1.ListUsersRequest
	(*User)(nil),              // 5: api.kvendingoldo.user.v1.User
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
	(*ListUsersResponse)(nil), // 7: api.kvendingoldo.user.v1.ListUsersResponse
}
var file_api_kvendingoldo_user_v1_service_proto_depIdxs = []int32{
	0, // 0: api.kvendingoldo.user.v1.UserService.Create:input_type -> api.kvendingoldo.user.v1.CreateUserRequest
	1, // 1: api.kvendingoldo.user.v1.UserService.Get:input_type -> api.kvendingoldo.user.v1.GetUserRequest
	2, // 2: api.kvendingoldo.user.v1.UserService.Update:input_type -> api.kvendingoldo.user.v1.UpdateUserRequest
	3, // 3: api.kvendingoldo.user.v1.UserService.Delete:input_type -> api.kvendingoldo.user.v1.DeleteUserRequest
	4, // 4: api.kvendingoldo.user.v1.UserService.List:input_type -> api.kvendingoldo.user.v1.ListUsersRequest
	5, // 5: api.kvendingoldo.user.v1.UserService.Create:output_type -> api.kvendingoldo.user.v1.User
	5, // 6: api.kvendingoldo.user.v1.UserService.Get:output_type -> api.kvendingoldo.user.v1.User
	5, // 7: api.kvendingoldo.user.v1.UserService.Update:output_type -> api.kvendingoldo.user.v1.User
	6, // 8: api.kvendingoldo.user.v1.UserService.Delete:output_type -> google.protobuf.Empty
	7, // 9: api.kvendingoldo.user.v1.UserService.List:output_type -> api.kvendingoldo.user.v1.ListUsersResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_kvendingoldo_user_v1_service_proto_init() }
func file_api_kvendingoldo_user_v1_service_proto_init() {
	if File_api_kvendingoldo_user_v1_service_proto != nil {
		return
	}
	file_api_kvendingoldo_user_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_kvendingoldo_user_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_kvendingoldo_user_v1_service_proto_goTypes,
		DependencyIndexes: file_api_kvendingoldo_user_v1_service_proto_depIdxs,
	}.Build()
	File_api_kvendingoldo_user_v1_service_proto = out.File
	file_api_kvendingoldo_user_v1_service_proto_rawDesc = nil
	file_api_kvendingoldo_user_v1_service_proto_goTypes = nil
	file_api_kvendingoldo_user_v1_service_proto_depIdxs = nil
}
