// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: v1/authorization/authorization.proto

package authorization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Authorization struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	mi := &file_v1_authorization_authorization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_v1_authorization_authorization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_v1_authorization_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *Authorization) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Authorization) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AuthorizeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	mi := &file_v1_authorization_authorization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_authorization_authorization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_v1_authorization_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthorizeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthorizeReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authorization *Authorization         `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorizeReply) Reset() {
	*x = AuthorizeReply{}
	mi := &file_v1_authorization_authorization_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeReply) ProtoMessage() {}

func (x *AuthorizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_authorization_authorization_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeReply.ProtoReflect.Descriptor instead.
func (*AuthorizeReply) Descriptor() ([]byte, []int) {
	return file_v1_authorization_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizeReply) GetAuthorization() *Authorization {
	if x != nil {
		return x.Authorization
	}
	return nil
}

type ReauthorizeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReauthorizeRequest) Reset() {
	*x = ReauthorizeRequest{}
	mi := &file_v1_authorization_authorization_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReauthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReauthorizeRequest) ProtoMessage() {}

func (x *ReauthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_authorization_authorization_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReauthorizeRequest.ProtoReflect.Descriptor instead.
func (*ReauthorizeRequest) Descriptor() ([]byte, []int) {
	return file_v1_authorization_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *ReauthorizeRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ReauthorizeReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authorization *Authorization         `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReauthorizeReply) Reset() {
	*x = ReauthorizeReply{}
	mi := &file_v1_authorization_authorization_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReauthorizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReauthorizeReply) ProtoMessage() {}

func (x *ReauthorizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_authorization_authorization_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReauthorizeReply.ProtoReflect.Descriptor instead.
func (*ReauthorizeReply) Descriptor() ([]byte, []int) {
	return file_v1_authorization_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *ReauthorizeReply) GetAuthorization() *Authorization {
	if x != nil {
		return x.Authorization
	}
	return nil
}

var File_v1_authorization_authorization_proto protoreflect.FileDescriptor

const file_v1_authorization_authorization_proto_rawDesc = "" +
	"\n" +
	"$v1/authorization/authorization.proto\x12\x10authorization.v1\"W\n" +
	"\rAuthorization\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\"G\n" +
	"\x10AuthorizeRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"W\n" +
	"\x0eAuthorizeReply\x12E\n" +
	"\rauthorization\x18\x01 \x01(\v2\x1f.authorization.v1.AuthorizationR\rauthorization\"9\n" +
	"\x12ReauthorizeRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"Y\n" +
	"\x10ReauthorizeReply\x12E\n" +
	"\rauthorization\x18\x01 \x01(\v2\x1f.authorization.v1.AuthorizationR\rauthorization2\xc2\x01\n" +
	"\x14AuthorizationService\x12Q\n" +
	"\tAuthorize\x12\".authorization.v1.AuthorizeRequest\x1a .authorization.v1.AuthorizeReply\x12W\n" +
	"\vReauthorize\x12$.authorization.v1.ReauthorizeRequest\x1a\".authorization.v1.ReauthorizeReplyBHZFgithub.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/authorizationb\x06proto3"

var (
	file_v1_authorization_authorization_proto_rawDescOnce sync.Once
	file_v1_authorization_authorization_proto_rawDescData []byte
)

func file_v1_authorization_authorization_proto_rawDescGZIP() []byte {
	file_v1_authorization_authorization_proto_rawDescOnce.Do(func() {
		file_v1_authorization_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_authorization_authorization_proto_rawDesc), len(file_v1_authorization_authorization_proto_rawDesc)))
	})
	return file_v1_authorization_authorization_proto_rawDescData
}

var file_v1_authorization_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_authorization_authorization_proto_goTypes = []any{
	(*Authorization)(nil),      // 0: authorization.v1.Authorization
	(*AuthorizeRequest)(nil),   // 1: authorization.v1.AuthorizeRequest
	(*AuthorizeReply)(nil),     // 2: authorization.v1.AuthorizeReply
	(*ReauthorizeRequest)(nil), // 3: authorization.v1.ReauthorizeRequest
	(*ReauthorizeReply)(nil),   // 4: authorization.v1.ReauthorizeReply
}
var file_v1_authorization_authorization_proto_depIdxs = []int32{
	0, // 0: authorization.v1.AuthorizeReply.authorization:type_name -> authorization.v1.Authorization
	0, // 1: authorization.v1.ReauthorizeReply.authorization:type_name -> authorization.v1.Authorization
	1, // 2: authorization.v1.AuthorizationService.Authorize:input_type -> authorization.v1.AuthorizeRequest
	3, // 3: authorization.v1.AuthorizationService.Reauthorize:input_type -> authorization.v1.ReauthorizeRequest
	2, // 4: authorization.v1.AuthorizationService.Authorize:output_type -> authorization.v1.AuthorizeReply
	4, // 5: authorization.v1.AuthorizationService.Reauthorize:output_type -> authorization.v1.ReauthorizeReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_authorization_authorization_proto_init() }
func file_v1_authorization_authorization_proto_init() {
	if File_v1_authorization_authorization_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_authorization_authorization_proto_rawDesc), len(file_v1_authorization_authorization_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_authorization_authorization_proto_goTypes,
		DependencyIndexes: file_v1_authorization_authorization_proto_depIdxs,
		MessageInfos:      file_v1_authorization_authorization_proto_msgTypes,
	}.Build()
	File_v1_authorization_authorization_proto = out.File
	file_v1_authorization_authorization_proto_goTypes = nil
	file_v1_authorization_authorization_proto_depIdxs = nil
}
