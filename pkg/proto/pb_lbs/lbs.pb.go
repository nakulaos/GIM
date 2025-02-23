// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.1
// source: pb_lbs/lbs.proto

package pb_lbs

import (
	pb_enum "GIM/pkg/proto/pb_enum"
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

type ReportLngLatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       int64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`              // uid
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"` // 经度
	Latitude  float64 `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`   // 纬度
}

func (x *ReportLngLatReq) Reset() {
	*x = ReportLngLatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_lbs_lbs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLngLatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLngLatReq) ProtoMessage() {}

func (x *ReportLngLatReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_lbs_lbs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLngLatReq.ProtoReflect.Descriptor instead.
func (*ReportLngLatReq) Descriptor() ([]byte, []int) {
	return file_pb_lbs_lbs_proto_rawDescGZIP(), []int{0}
}

func (x *ReportLngLatReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReportLngLatReq) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ReportLngLatReq) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type ReportLngLatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReportLngLatResp) Reset() {
	*x = ReportLngLatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_lbs_lbs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportLngLatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLngLatResp) ProtoMessage() {}

func (x *ReportLngLatResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_lbs_lbs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLngLatResp.ProtoReflect.Descriptor instead.
func (*ReportLngLatResp) Descriptor() ([]byte, []int) {
	return file_pb_lbs_lbs_proto_rawDescGZIP(), []int{1}
}

func (x *ReportLngLatResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ReportLngLatResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PeopleNearbyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         int64          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                           // uid
	Longitude   float64        `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`              // 经度
	Latitude    float64        `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`                // 纬度
	Radius      int64          `protobuf:"varint,4,opt,name=radius,proto3" json:"radius,omitempty"`                     // 半径
	Gender      pb_enum.GENDER `protobuf:"varint,5,opt,name=gender,proto3,enum=pb_enum.GENDER" json:"gender,omitempty"` // 性别
	MinAge      int32          `protobuf:"varint,6,opt,name=min_age,json=minAge,proto3" json:"min_age,omitempty"`       // 最小年龄
	MaxAge      int32          `protobuf:"varint,7,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`       // 最大年龄
	Limit       int64          `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`                       // 数量限制
	Skip        int64          `protobuf:"varint,9,opt,name=Skip,proto3" json:"Skip,omitempty"`
	MinDistance float64        `protobuf:"fixed64,10,opt,name=min_distance,json=minDistance,proto3" json:"min_distance,omitempty"`
}

func (x *PeopleNearbyReq) Reset() {
	*x = PeopleNearbyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_lbs_lbs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleNearbyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleNearbyReq) ProtoMessage() {}

func (x *PeopleNearbyReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_lbs_lbs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleNearbyReq.ProtoReflect.Descriptor instead.
func (*PeopleNearbyReq) Descriptor() ([]byte, []int) {
	return file_pb_lbs_lbs_proto_rawDescGZIP(), []int{2}
}

func (x *PeopleNearbyReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PeopleNearbyReq) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PeopleNearbyReq) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PeopleNearbyReq) GetRadius() int64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *PeopleNearbyReq) GetGender() pb_enum.GENDER {
	if x != nil {
		return x.Gender
	}
	return pb_enum.GENDER(0)
}

func (x *PeopleNearbyReq) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *PeopleNearbyReq) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *PeopleNearbyReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PeopleNearbyReq) GetSkip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *PeopleNearbyReq) GetMinDistance() float64 {
	if x != nil {
		return x.MinDistance
	}
	return 0
}

type PeopleNearbyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	List []*UserLocation `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PeopleNearbyResp) Reset() {
	*x = PeopleNearbyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_lbs_lbs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleNearbyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleNearbyResp) ProtoMessage() {}

func (x *PeopleNearbyResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_lbs_lbs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleNearbyResp.ProtoReflect.Descriptor instead.
func (*PeopleNearbyResp) Descriptor() ([]byte, []int) {
	return file_pb_lbs_lbs_proto_rawDescGZIP(), []int{3}
}

func (x *PeopleNearbyResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PeopleNearbyResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PeopleNearbyResp) GetList() []*UserLocation {
	if x != nil {
		return x.List
	}
	return nil
}

type UserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         int64          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                           // uid
	Gender      pb_enum.GENDER `protobuf:"varint,2,opt,name=gender,proto3,enum=pb_enum.GENDER" json:"gender,omitempty"` // 性别
	Age         int32          `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`                           // 生日
	Nickname    string         `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`                  // 昵称
	Avatar      string         `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`                      // 头像
	Distance    int64          `protobuf:"varint,6,opt,name=distance,proto3" json:"distance,omitempty"`                 // 距离
	Coordinates []float64      `protobuf:"fixed64,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`   // 经纬度
}

func (x *UserLocation) Reset() {
	*x = UserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_lbs_lbs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocation) ProtoMessage() {}

func (x *UserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_lbs_lbs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocation.ProtoReflect.Descriptor instead.
func (*UserLocation) Descriptor() ([]byte, []int) {
	return file_pb_lbs_lbs_proto_rawDescGZIP(), []int{4}
}

func (x *UserLocation) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserLocation) GetGender() pb_enum.GENDER {
	if x != nil {
		return x.Gender
	}
	return pb_enum.GENDER(0)
}

func (x *UserLocation) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserLocation) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserLocation) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserLocation) GetDistance() int64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *UserLocation) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_pb_lbs_lbs_proto protoreflect.FileDescriptor

var file_pb_lbs_lbs_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2f, 0x6c, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d,
	0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x38, 0x0a,
	0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x9d, 0x02, 0x0a, 0x0f, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12,
	0x27, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x41, 0x67,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x6b, 0x69, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x53, 0x6b, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x44,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x62, 0x0a, 0x10, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xcd, 0x01, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x03,
	0x4c, 0x62, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6e, 0x67,
	0x4c, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6e, 0x67, 0x4c,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2e,
	0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x4e,
	0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1e, 0x5a, 0x1c, 0x6c, 0x61, 0x72,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x6c,
	0x62, 0x73, 0x3b, 0x70, 0x62, 0x5f, 0x6c, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_lbs_lbs_proto_rawDescOnce sync.Once
	file_pb_lbs_lbs_proto_rawDescData = file_pb_lbs_lbs_proto_rawDesc
)

func file_pb_lbs_lbs_proto_rawDescGZIP() []byte {
	file_pb_lbs_lbs_proto_rawDescOnce.Do(func() {
		file_pb_lbs_lbs_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_lbs_lbs_proto_rawDescData)
	})
	return file_pb_lbs_lbs_proto_rawDescData
}

var file_pb_lbs_lbs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_lbs_lbs_proto_goTypes = []interface{}{
	(*ReportLngLatReq)(nil),  // 0: pb_lbs.ReportLngLatReq
	(*ReportLngLatResp)(nil), // 1: pb_lbs.ReportLngLatResp
	(*PeopleNearbyReq)(nil),  // 2: pb_lbs.PeopleNearbyReq
	(*PeopleNearbyResp)(nil), // 3: pb_lbs.PeopleNearbyResp
	(*UserLocation)(nil),     // 4: pb_lbs.UserLocation
	(pb_enum.GENDER)(0),      // 5: pb_enum.GENDER
}
var file_pb_lbs_lbs_proto_depIdxs = []int32{
	5, // 0: pb_lbs.PeopleNearbyReq.gender:type_name -> pb_enum.GENDER
	4, // 1: pb_lbs.PeopleNearbyResp.list:type_name -> pb_lbs.UserLocation
	5, // 2: pb_lbs.UserLocation.gender:type_name -> pb_enum.GENDER
	0, // 3: pb_lbs.Lbs.ReportLngLat:input_type -> pb_lbs.ReportLngLatReq
	2, // 4: pb_lbs.Lbs.PeopleNearby:input_type -> pb_lbs.PeopleNearbyReq
	1, // 5: pb_lbs.Lbs.ReportLngLat:output_type -> pb_lbs.ReportLngLatResp
	3, // 6: pb_lbs.Lbs.PeopleNearby:output_type -> pb_lbs.PeopleNearbyResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_lbs_lbs_proto_init() }
func file_pb_lbs_lbs_proto_init() {
	if File_pb_lbs_lbs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_lbs_lbs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportLngLatReq); i {
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
		file_pb_lbs_lbs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportLngLatResp); i {
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
		file_pb_lbs_lbs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleNearbyReq); i {
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
		file_pb_lbs_lbs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleNearbyResp); i {
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
		file_pb_lbs_lbs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLocation); i {
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
			RawDescriptor: file_pb_lbs_lbs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_lbs_lbs_proto_goTypes,
		DependencyIndexes: file_pb_lbs_lbs_proto_depIdxs,
		MessageInfos:      file_pb_lbs_lbs_proto_msgTypes,
	}.Build()
	File_pb_lbs_lbs_proto = out.File
	file_pb_lbs_lbs_proto_rawDesc = nil
	file_pb_lbs_lbs_proto_goTypes = nil
	file_pb_lbs_lbs_proto_depIdxs = nil
}
