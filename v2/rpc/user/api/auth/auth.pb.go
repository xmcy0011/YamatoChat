// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rpc/user/api/auth/auth.proto

package auth

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest_LoginType int32

const (
	AuthRequest_loginTypeMobile AuthRequest_LoginType = 0
)

// Enum value maps for AuthRequest_LoginType.
var (
	AuthRequest_LoginType_name = map[int32]string{
		0: "loginTypeMobile",
	}
	AuthRequest_LoginType_value = map[string]int32{
		"loginTypeMobile": 0,
	}
)

func (x AuthRequest_LoginType) Enum() *AuthRequest_LoginType {
	p := new(AuthRequest_LoginType)
	*p = x
	return p
}

func (x AuthRequest_LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthRequest_LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_user_api_auth_auth_proto_enumTypes[0].Descriptor()
}

func (AuthRequest_LoginType) Type() protoreflect.EnumType {
	return &file_rpc_user_api_auth_auth_proto_enumTypes[0]
}

func (x AuthRequest_LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthRequest_LoginType.Descriptor instead.
func (AuthRequest_LoginType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{2, 0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId   string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	AppVersion string `protobuf:"bytes,2,opt,name=appVersion,proto3" json:"appVersion,omitempty"`
	OsVersion  string `protobuf:"bytes,3,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_api_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_api_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *RegisterRequest) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *RegisterRequest) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorReason `protobuf:"varint,1,opt,name=code,proto3,enum=auth.ErrorReason" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_api_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_api_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetCode() ErrorReason {
	if x != nil {
		return x.Code
	}
	return ErrorReason_GEETER_UNSPECIFIED
}

func (x *RegisterReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginType AuthRequest_LoginType   `protobuf:"varint,1,opt,name=loginType,proto3,enum=auth.AuthRequest_LoginType" json:"loginType,omitempty"`
	ByMobile  *AuthRequest_MobileAuth `protobuf:"bytes,2,opt,name=byMobile,proto3" json:"byMobile,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_api_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_api_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRequest) GetLoginType() AuthRequest_LoginType {
	if x != nil {
		return x.LoginType
	}
	return AuthRequest_loginTypeMobile
}

func (x *AuthRequest) GetByMobile() *AuthRequest_MobileAuth {
	if x != nil {
		return x.ByMobile
	}
	return nil
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ErrorReason `protobuf:"varint,1,opt,name=code,proto3,enum=auth.ErrorReason" json:"code,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_api_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_api_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthReply) GetCode() ErrorReason {
	if x != nil {
		return x.Code
	}
	return ErrorReason_GEETER_UNSPECIFIED
}

func (x *AuthReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AuthRequest_MobileAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Code  string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *AuthRequest_MobileAuth) Reset() {
	*x = AuthRequest_MobileAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_api_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest_MobileAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest_MobileAuth) ProtoMessage() {}

func (x *AuthRequest_MobileAuth) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_api_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest_MobileAuth.ProtoReflect.Descriptor instead.
func (*AuthRequest_MobileAuth) Descriptor() ([]byte, []int) {
	return file_rpc_user_api_auth_auth_proto_rawDescGZIP(), []int{2, 0}
}

func (x *AuthRequest_MobileAuth) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AuthRequest_MobileAuth) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_rpc_user_api_auth_auth_proto protoreflect.FileDescriptor

var file_rpc_user_api_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x62, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x08, 0x62, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x1a, 0x36, 0x0a, 0x0a,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x13, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x10, 0x00, 0x22, 0x44, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa4, 0x01, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0x42, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x14, 0x5a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_user_api_auth_auth_proto_rawDescOnce sync.Once
	file_rpc_user_api_auth_auth_proto_rawDescData = file_rpc_user_api_auth_auth_proto_rawDesc
)

func file_rpc_user_api_auth_auth_proto_rawDescGZIP() []byte {
	file_rpc_user_api_auth_auth_proto_rawDescOnce.Do(func() {
		file_rpc_user_api_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_user_api_auth_auth_proto_rawDescData)
	})
	return file_rpc_user_api_auth_auth_proto_rawDescData
}

var file_rpc_user_api_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_user_api_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_user_api_auth_auth_proto_goTypes = []interface{}{
	(AuthRequest_LoginType)(0),     // 0: auth.AuthRequest.LoginType
	(*RegisterRequest)(nil),        // 1: auth.RegisterRequest
	(*RegisterReply)(nil),          // 2: auth.RegisterReply
	(*AuthRequest)(nil),            // 3: auth.AuthRequest
	(*AuthReply)(nil),              // 4: auth.AuthReply
	(*AuthRequest_MobileAuth)(nil), // 5: auth.AuthRequest.MobileAuth
	(ErrorReason)(0),               // 6: auth.ErrorReason
}
var file_rpc_user_api_auth_auth_proto_depIdxs = []int32{
	6, // 0: auth.RegisterReply.code:type_name -> auth.ErrorReason
	0, // 1: auth.AuthRequest.loginType:type_name -> auth.AuthRequest.LoginType
	5, // 2: auth.AuthRequest.byMobile:type_name -> auth.AuthRequest.MobileAuth
	6, // 3: auth.AuthReply.code:type_name -> auth.ErrorReason
	1, // 4: auth.Auth.Register:input_type -> auth.RegisterRequest
	3, // 5: auth.Auth.Auth:input_type -> auth.AuthRequest
	2, // 6: auth.Auth.Register:output_type -> auth.RegisterReply
	4, // 7: auth.Auth.Auth:output_type -> auth.AuthReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_user_api_auth_auth_proto_init() }
func file_rpc_user_api_auth_auth_proto_init() {
	if File_rpc_user_api_auth_auth_proto != nil {
		return
	}
	file_rpc_user_api_auth_error_reason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_user_api_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_rpc_user_api_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_rpc_user_api_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_rpc_user_api_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_rpc_user_api_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest_MobileAuth); i {
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
			RawDescriptor: file_rpc_user_api_auth_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_user_api_auth_auth_proto_goTypes,
		DependencyIndexes: file_rpc_user_api_auth_auth_proto_depIdxs,
		EnumInfos:         file_rpc_user_api_auth_auth_proto_enumTypes,
		MessageInfos:      file_rpc_user_api_auth_auth_proto_msgTypes,
	}.Build()
	File_rpc_user_api_auth_auth_proto = out.File
	file_rpc_user_api_auth_auth_proto_rawDesc = nil
	file_rpc_user_api_auth_auth_proto_goTypes = nil
	file_rpc_user_api_auth_auth_proto_depIdxs = nil
}
