// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: v1/room/room.proto

package room

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

type PianoType int32

const (
	PianoType_PIANO_TYPE_UNSPECIFIED PianoType = 0
	PianoType_GRAND                  PianoType = 1
	PianoType_UPRIGHT                PianoType = 2
	PianoType_NONE                   PianoType = 3
)

// Enum value maps for PianoType.
var (
	PianoType_name = map[int32]string{
		0: "PIANO_TYPE_UNSPECIFIED",
		1: "GRAND",
		2: "UPRIGHT",
		3: "NONE",
	}
	PianoType_value = map[string]int32{
		"PIANO_TYPE_UNSPECIFIED": 0,
		"GRAND":                  1,
		"UPRIGHT":                2,
		"NONE":                   3,
	}
)

func (x PianoType) Enum() *PianoType {
	p := new(PianoType)
	*p = x
	return p
}

func (x PianoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PianoType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_room_room_proto_enumTypes[0].Descriptor()
}

func (PianoType) Type() protoreflect.EnumType {
	return &file_v1_room_room_proto_enumTypes[0]
}

func (x PianoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PianoType.Descriptor instead.
func (PianoType) EnumDescriptor() ([]byte, []int) {
	return file_v1_room_room_proto_rawDescGZIP(), []int{0}
}

type CampusType int32

const (
	CampusType_CAMPUS_UNSPECIFIED CampusType = 0
	CampusType_NAKAMEGURO         CampusType = 1
	CampusType_IKEBUKURO          CampusType = 2
)

// Enum value maps for CampusType.
var (
	CampusType_name = map[int32]string{
		0: "CAMPUS_UNSPECIFIED",
		1: "NAKAMEGURO",
		2: "IKEBUKURO",
	}
	CampusType_value = map[string]int32{
		"CAMPUS_UNSPECIFIED": 0,
		"NAKAMEGURO":         1,
		"IKEBUKURO":          2,
	}
)

func (x CampusType) Enum() *CampusType {
	p := new(CampusType)
	*p = x
	return p
}

func (x CampusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampusType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_room_room_proto_enumTypes[1].Descriptor()
}

func (CampusType) Type() protoreflect.EnumType {
	return &file_v1_room_room_proto_enumTypes[1]
}

func (x CampusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampusType.Descriptor instead.
func (CampusType) EnumDescriptor() ([]byte, []int) {
	return file_v1_room_room_proto_rawDescGZIP(), []int{1}
}

type Room struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PianoType     PianoType              `protobuf:"varint,3,opt,name=piano_type,json=pianoType,proto3,enum=room.v1.PianoType" json:"piano_type,omitempty"`
	PianoNumber   int32                  `protobuf:"varint,4,opt,name=piano_number,json=pianoNumber,proto3" json:"piano_number,omitempty"`
	IsClassroom   bool                   `protobuf:"varint,5,opt,name=is_classroom,json=isClassroom,proto3" json:"is_classroom,omitempty"`
	IsBasement    bool                   `protobuf:"varint,6,opt,name=is_basement,json=isBasement,proto3" json:"is_basement,omitempty"`
	CampusType    CampusType             `protobuf:"varint,7,opt,name=campus_type,json=campusType,proto3,enum=room.v1.CampusType" json:"campus_type,omitempty"`
	Floor         int32                  `protobuf:"varint,8,opt,name=floor,proto3" json:"floor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_v1_room_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_v1_room_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_v1_room_room_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetPianoType() PianoType {
	if x != nil {
		return x.PianoType
	}
	return PianoType_PIANO_TYPE_UNSPECIFIED
}

func (x *Room) GetPianoNumber() int32 {
	if x != nil {
		return x.PianoNumber
	}
	return 0
}

func (x *Room) GetIsClassroom() bool {
	if x != nil {
		return x.IsClassroom
	}
	return false
}

func (x *Room) GetIsBasement() bool {
	if x != nil {
		return x.IsBasement
	}
	return false
}

func (x *Room) GetCampusType() CampusType {
	if x != nil {
		return x.CampusType
	}
	return CampusType_CAMPUS_UNSPECIFIED
}

func (x *Room) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

type GetRoomsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomsRequest) Reset() {
	*x = GetRoomsRequest{}
	mi := &file_v1_room_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsRequest) ProtoMessage() {}

func (x *GetRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_room_room_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsRequest.ProtoReflect.Descriptor instead.
func (*GetRoomsRequest) Descriptor() ([]byte, []int) {
	return file_v1_room_room_proto_rawDescGZIP(), []int{1}
}

type GetRoomsReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rooms         []*Room                `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRoomsReply) Reset() {
	*x = GetRoomsReply{}
	mi := &file_v1_room_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoomsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsReply) ProtoMessage() {}

func (x *GetRoomsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_room_room_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsReply.ProtoReflect.Descriptor instead.
func (*GetRoomsReply) Descriptor() ([]byte, []int) {
	return file_v1_room_room_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomsReply) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_v1_room_room_proto protoreflect.FileDescriptor

const file_v1_room_room_proto_rawDesc = "" +
	"\n" +
	"\x12v1/room/room.proto\x12\aroom.v1\"\x90\x02\n" +
	"\x04Room\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x121\n" +
	"\n" +
	"piano_type\x18\x03 \x01(\x0e2\x12.room.v1.PianoTypeR\tpianoType\x12!\n" +
	"\fpiano_number\x18\x04 \x01(\x05R\vpianoNumber\x12!\n" +
	"\fis_classroom\x18\x05 \x01(\bR\visClassroom\x12\x1f\n" +
	"\vis_basement\x18\x06 \x01(\bR\n" +
	"isBasement\x124\n" +
	"\vcampus_type\x18\a \x01(\x0e2\x13.room.v1.CampusTypeR\n" +
	"campusType\x12\x14\n" +
	"\x05floor\x18\b \x01(\x05R\x05floor\"\x11\n" +
	"\x0fGetRoomsRequest\"4\n" +
	"\rGetRoomsReply\x12#\n" +
	"\x05rooms\x18\x01 \x03(\v2\r.room.v1.RoomR\x05rooms*I\n" +
	"\tPianoType\x12\x1a\n" +
	"\x16PIANO_TYPE_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05GRAND\x10\x01\x12\v\n" +
	"\aUPRIGHT\x10\x02\x12\b\n" +
	"\x04NONE\x10\x03*C\n" +
	"\n" +
	"CampusType\x12\x16\n" +
	"\x12CAMPUS_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"NAKAMEGURO\x10\x01\x12\r\n" +
	"\tIKEBUKURO\x10\x022K\n" +
	"\vRoomService\x12<\n" +
	"\bGetRooms\x12\x18.room.v1.GetRoomsRequest\x1a\x16.room.v1.GetRoomsReplyB?Z=github.com/ekkx/tcmrsv-web/server/internal/shared/api/v1/roomb\x06proto3"

var (
	file_v1_room_room_proto_rawDescOnce sync.Once
	file_v1_room_room_proto_rawDescData []byte
)

func file_v1_room_room_proto_rawDescGZIP() []byte {
	file_v1_room_room_proto_rawDescOnce.Do(func() {
		file_v1_room_room_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_room_room_proto_rawDesc), len(file_v1_room_room_proto_rawDesc)))
	})
	return file_v1_room_room_proto_rawDescData
}

var file_v1_room_room_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_v1_room_room_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_room_room_proto_goTypes = []any{
	(PianoType)(0),          // 0: room.v1.PianoType
	(CampusType)(0),         // 1: room.v1.CampusType
	(*Room)(nil),            // 2: room.v1.Room
	(*GetRoomsRequest)(nil), // 3: room.v1.GetRoomsRequest
	(*GetRoomsReply)(nil),   // 4: room.v1.GetRoomsReply
}
var file_v1_room_room_proto_depIdxs = []int32{
	0, // 0: room.v1.Room.piano_type:type_name -> room.v1.PianoType
	1, // 1: room.v1.Room.campus_type:type_name -> room.v1.CampusType
	2, // 2: room.v1.GetRoomsReply.rooms:type_name -> room.v1.Room
	3, // 3: room.v1.RoomService.GetRooms:input_type -> room.v1.GetRoomsRequest
	4, // 4: room.v1.RoomService.GetRooms:output_type -> room.v1.GetRoomsReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_room_room_proto_init() }
func file_v1_room_room_proto_init() {
	if File_v1_room_room_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_room_room_proto_rawDesc), len(file_v1_room_room_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_room_room_proto_goTypes,
		DependencyIndexes: file_v1_room_room_proto_depIdxs,
		EnumInfos:         file_v1_room_room_proto_enumTypes,
		MessageInfos:      file_v1_room_room_proto_msgTypes,
	}.Build()
	File_v1_room_room_proto = out.File
	file_v1_room_room_proto_goTypes = nil
	file_v1_room_room_proto_depIdxs = nil
}
