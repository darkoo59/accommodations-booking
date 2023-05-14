// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: user/user-service.proto

package user

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

type UserRole int32

const (
	UserRole_Host  UserRole = 0
	UserRole_Guest UserRole = 1
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "Host",
		1: "Guest",
	}
	UserRole_value = map[string]int32{
		"Host":  0,
		"Guest": 1,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_user_user_service_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_user_user_service_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{0}
}

type Empty_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty_Request) Reset() {
	*x = Empty_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty_Request) ProtoMessage() {}

func (x *Empty_Request) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty_Request.ProtoReflect.Descriptor instead.
func (*Empty_Request) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{0}
}

type US_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *US_Request) Reset() {
	*x = US_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *US_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*US_Request) ProtoMessage() {}

func (x *US_Request) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use US_Request.ProtoReflect.Descriptor instead.
func (*US_Request) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *US_Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type US_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *US_Response) Reset() {
	*x = US_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *US_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*US_Response) ProtoMessage() {}

func (x *US_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use US_Response.ProtoReflect.Descriptor instead.
func (*US_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *US_Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RegisterUser_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        UserRole `protobuf:"varint,1,opt,name=Role,proto3,enum=user.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password    string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	FirstName   string   `protobuf:"bytes,4,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,5,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,7,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *RegisterUser_Request) Reset() {
	*x = RegisterUser_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser_Request) ProtoMessage() {}

func (x *RegisterUser_Request) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser_Request.ProtoReflect.Descriptor instead.
func (*RegisterUser_Request) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUser_Request) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *RegisterUser_Request) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUser_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterUser_Request) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterUser_Request) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterUser_Request) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUser_Request) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

type RegisterUser_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                           `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Message string                         `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Errors  []*RegisterUser_Response_Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *RegisterUser_Response) Reset() {
	*x = RegisterUser_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser_Response) ProtoMessage() {}

func (x *RegisterUser_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser_Response.ProtoReflect.Descriptor instead.
func (*RegisterUser_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterUser_Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RegisterUser_Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterUser_Response) GetErrors() []*RegisterUser_Response_Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetAllUsers_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GetAllUsers_Response_User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllUsers_Response) Reset() {
	*x = GetAllUsers_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsers_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsers_Response) ProtoMessage() {}

func (x *GetAllUsers_Response) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsers_Response.ProtoReflect.Descriptor instead.
func (*GetAllUsers_Response) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllUsers_Response) GetUsers() []*GetAllUsers_Response_User {
	if x != nil {
		return x.Users
	}
	return nil
}

type RegisterUser_Response_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyName string `protobuf:"bytes,1,opt,name=PropertyName,proto3" json:"PropertyName,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
}

func (x *RegisterUser_Response_Error) Reset() {
	*x = RegisterUser_Response_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUser_Response_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUser_Response_Error) ProtoMessage() {}

func (x *RegisterUser_Response_Error) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUser_Response_Error.ProtoReflect.Descriptor instead.
func (*RegisterUser_Response_Error) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *RegisterUser_Response_Error) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *RegisterUser_Response_Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetAllUsers_Response_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role        UserRole `protobuf:"varint,1,opt,name=Role,proto3,enum=user.UserRole" json:"Role,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	FirstName   string   `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string   `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	LivingPlace string   `protobuf:"bytes,6,opt,name=LivingPlace,proto3" json:"LivingPlace,omitempty"`
}

func (x *GetAllUsers_Response_User) Reset() {
	*x = GetAllUsers_Response_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsers_Response_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsers_Response_User) ProtoMessage() {}

func (x *GetAllUsers_Response_User) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsers_Response_User.ProtoReflect.Descriptor instead.
func (*GetAllUsers_Response_User) Descriptor() ([]byte, []int) {
	return file_user_user_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetAllUsers_Response_User) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_Host
}

func (x *GetAllUsers_Response_User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAllUsers_Response_User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetAllUsers_Response_User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetAllUsers_Response_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAllUsers_Response_User) GetLivingPlace() string {
	if x != nil {
		return x.LivingPlace
	}
	return ""
}

var File_user_user_service_proto protoreflect.FileDescriptor

var file_user_user_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20,
	0x0a, 0x0a, 0x55, 0x53, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x27, 0x0a, 0x0b, 0x55, 0x53, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x14, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x22, 0xd7, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x4f, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0xb8, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x2a, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x10, 0x01, 0x32, 0xf4, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x53, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x53, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x12, 0x47, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_user_service_proto_rawDescOnce sync.Once
	file_user_user_service_proto_rawDescData = file_user_user_service_proto_rawDesc
)

func file_user_user_service_proto_rawDescGZIP() []byte {
	file_user_user_service_proto_rawDescOnce.Do(func() {
		file_user_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_service_proto_rawDescData)
	})
	return file_user_user_service_proto_rawDescData
}

var file_user_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_user_service_proto_goTypes = []interface{}{
	(UserRole)(0),                       // 0: user.UserRole
	(*Empty_Request)(nil),               // 1: user.Empty_Request
	(*US_Request)(nil),                  // 2: user.US_Request
	(*US_Response)(nil),                 // 3: user.US_Response
	(*RegisterUser_Request)(nil),        // 4: user.RegisterUser_Request
	(*RegisterUser_Response)(nil),       // 5: user.RegisterUser_Response
	(*GetAllUsers_Response)(nil),        // 6: user.GetAllUsers_Response
	(*RegisterUser_Response_Error)(nil), // 7: user.RegisterUser_Response.Error
	(*GetAllUsers_Response_User)(nil),   // 8: user.GetAllUsers_Response.User
}
var file_user_user_service_proto_depIdxs = []int32{
	0, // 0: user.RegisterUser_Request.Role:type_name -> user.UserRole
	7, // 1: user.RegisterUser_Response.errors:type_name -> user.RegisterUser_Response.Error
	8, // 2: user.GetAllUsers_Response.users:type_name -> user.GetAllUsers_Response.User
	0, // 3: user.GetAllUsers_Response.User.Role:type_name -> user.UserRole
	2, // 4: user.UserService.GetUser:input_type -> user.US_Request
	4, // 5: user.UserService.RegisterUser:input_type -> user.RegisterUser_Request
	1, // 6: user.UserService.GetAllUsers:input_type -> user.Empty_Request
	3, // 7: user.UserService.GetUser:output_type -> user.US_Response
	5, // 8: user.UserService.RegisterUser:output_type -> user.RegisterUser_Response
	6, // 9: user.UserService.GetAllUsers:output_type -> user.GetAllUsers_Response
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_user_service_proto_init() }
func file_user_user_service_proto_init() {
	if File_user_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty_Request); i {
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
		file_user_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*US_Request); i {
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
		file_user_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*US_Response); i {
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
		file_user_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser_Request); i {
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
		file_user_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser_Response); i {
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
		file_user_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsers_Response); i {
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
		file_user_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUser_Response_Error); i {
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
		file_user_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsers_Response_User); i {
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
			RawDescriptor: file_user_user_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_service_proto_goTypes,
		DependencyIndexes: file_user_user_service_proto_depIdxs,
		EnumInfos:         file_user_user_service_proto_enumTypes,
		MessageInfos:      file_user_user_service_proto_msgTypes,
	}.Build()
	File_user_user_service_proto = out.File
	file_user_user_service_proto_rawDesc = nil
	file_user_user_service_proto_goTypes = nil
	file_user_user_service_proto_depIdxs = nil
}
