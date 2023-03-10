// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: magellan/v1/authentication_service_login.proto

package magellanv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticationServiceLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthenticationType AuthenticationType `protobuf:"varint,1,opt,name=authentication_type,json=authenticationType,proto3,enum=magellan.v1.AuthenticationType" json:"authentication_type,omitempty"`
}

func (x *AuthenticationServiceLoginRequest) Reset() {
	*x = AuthenticationServiceLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magellan_v1_authentication_service_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationServiceLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationServiceLoginRequest) ProtoMessage() {}

func (x *AuthenticationServiceLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_magellan_v1_authentication_service_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationServiceLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationServiceLoginRequest) Descriptor() ([]byte, []int) {
	return file_magellan_v1_authentication_service_login_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationServiceLoginRequest) GetAuthenticationType() AuthenticationType {
	if x != nil {
		return x.AuthenticationType
	}
	return AuthenticationType_AUTHENTICATION_TYPE_UNSPECIFIED
}

type AuthenticationServiceLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthenticationServiceLoginResponse) Reset() {
	*x = AuthenticationServiceLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_magellan_v1_authentication_service_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationServiceLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationServiceLoginResponse) ProtoMessage() {}

func (x *AuthenticationServiceLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_magellan_v1_authentication_service_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationServiceLoginResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationServiceLoginResponse) Descriptor() ([]byte, []int) {
	return file_magellan_v1_authentication_service_login_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationServiceLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_magellan_v1_authentication_service_login_proto protoreflect.FileDescriptor

var file_magellan_v1_authentication_service_login_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d,
	0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x21, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x13, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x55, 0x92, 0x41,
	0x52, 0x0a, 0x50, 0x2a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32,
	0x1c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0xd2, 0x01, 0x13,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x22, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x41, 0x92, 0x41, 0x3e, 0x0a, 0x3c, 0x2a, 0x1b, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x1d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x20, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa7, 0x02, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x1f, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x61, 0x76, 0x69,
	0x6e, 0x6a, 0x75, 0x6e, 0x65, 0x2f, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x4d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x17, 0x4d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x61,
	0x67, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x92, 0x41, 0x64, 0x12, 0x05, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x72, 0x5b, 0x0a, 0x31, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x61, 0x6e, 0x64,
	0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x61, 0x76, 0x69, 0x6e, 0x6a, 0x75, 0x6e, 0x65, 0x2f, 0x6d, 0x61, 0x67, 0x65, 0x6c, 0x6c, 0x61,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_magellan_v1_authentication_service_login_proto_rawDescOnce sync.Once
	file_magellan_v1_authentication_service_login_proto_rawDescData = file_magellan_v1_authentication_service_login_proto_rawDesc
)

func file_magellan_v1_authentication_service_login_proto_rawDescGZIP() []byte {
	file_magellan_v1_authentication_service_login_proto_rawDescOnce.Do(func() {
		file_magellan_v1_authentication_service_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_magellan_v1_authentication_service_login_proto_rawDescData)
	})
	return file_magellan_v1_authentication_service_login_proto_rawDescData
}

var file_magellan_v1_authentication_service_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_magellan_v1_authentication_service_login_proto_goTypes = []interface{}{
	(*AuthenticationServiceLoginRequest)(nil),  // 0: magellan.v1.AuthenticationServiceLoginRequest
	(*AuthenticationServiceLoginResponse)(nil), // 1: magellan.v1.AuthenticationServiceLoginResponse
	(AuthenticationType)(0),                    // 2: magellan.v1.AuthenticationType
}
var file_magellan_v1_authentication_service_login_proto_depIdxs = []int32{
	2, // 0: magellan.v1.AuthenticationServiceLoginRequest.authentication_type:type_name -> magellan.v1.AuthenticationType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_magellan_v1_authentication_service_login_proto_init() }
func file_magellan_v1_authentication_service_login_proto_init() {
	if File_magellan_v1_authentication_service_login_proto != nil {
		return
	}
	file_magellan_v1_authentication_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_magellan_v1_authentication_service_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationServiceLoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_magellan_v1_authentication_service_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationServiceLoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_magellan_v1_authentication_service_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_magellan_v1_authentication_service_login_proto_goTypes,
		DependencyIndexes: file_magellan_v1_authentication_service_login_proto_depIdxs,
		MessageInfos:      file_magellan_v1_authentication_service_login_proto_msgTypes,
	}.Build()
	File_magellan_v1_authentication_service_login_proto = out.File
	file_magellan_v1_authentication_service_login_proto_rawDesc = nil
	file_magellan_v1_authentication_service_login_proto_goTypes = nil
	file_magellan_v1_authentication_service_login_proto_depIdxs = nil
}
