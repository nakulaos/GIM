// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb_red_env/red_env.proto

package pb_red_env

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb_enum "lark/pkg/proto/pb_enum"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RedEnvelopeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvId          int64                       `protobuf:"varint,1,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`                                                        // 红包ID
	EnvType        pb_enum.RED_ENVELOPE_TYPE   `protobuf:"varint,2,opt,name=env_type,json=envType,proto3,enum=pb_enum.RED_ENVELOPE_TYPE" json:"env_type,omitempty"`                   // 红包类型 1-均分红包 2-碰运气红包
	ReceiverType   pb_enum.RECEIVER_TYPE       `protobuf:"varint,3,opt,name=receiver_type,json=receiverType,proto3,enum=pb_enum.RECEIVER_TYPE" json:"receiver_type,omitempty"`        // 接收者类型 1-私聊对方 2-群聊所有人 3-群聊指定人
	TradeNo        string                      `protobuf:"bytes,4,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`                                                   // 交易编号
	ChatId         int64                       `protobuf:"varint,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`                                                     // 聊天ID
	SenderUid      int64                       `protobuf:"varint,6,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`                                            // 发红包用户ID
	Total          int64                       `protobuf:"varint,7,opt,name=total,proto3" json:"total,omitempty"`                                                                     // 红包总金额(分)
	Quantity       int32                       `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`                                                               // 红包数量
	RemainQuantity int32                       `protobuf:"varint,9,opt,name=remain_quantity,json=remainQuantity,proto3" json:"remain_quantity,omitempty"`                             // 剩余红包数量
	RemainAmount   int64                       `protobuf:"varint,10,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`                                  // 剩余红包金额(分)
	Message        string                      `protobuf:"bytes,11,opt,name=message,proto3" json:"message,omitempty"`                                                                 // 祝福语
	ReceiverUids   []int64                     `protobuf:"varint,12,rep,packed,name=receiver_uids,json=receiverUids,proto3" json:"receiver_uids,omitempty"`                           // 接收者ID
	ExpiredTs      int64                       `protobuf:"varint,13,opt,name=expired_ts,json=expiredTs,proto3" json:"expired_ts,omitempty"`                                           // 红包过期时间
	PayStatus      pb_enum.PAYMENT_STATUS      `protobuf:"varint,14,opt,name=pay_status,json=payStatus,proto3,enum=pb_enum.PAYMENT_STATUS" json:"pay_status,omitempty"`               // 支付状态 0-未支付 1-支付中 2-已支付 3-支付失败
	EnvStatus      pb_enum.RED_ENVELOPE_STATUS `protobuf:"varint,15,opt,name=env_status,json=envStatus,proto3,enum=pb_enum.RED_ENVELOPE_STATUS" json:"env_status,omitempty"`          // 状态 0-创建 1-已发放 2-已领完 3-已过期且退还剩余红包
	SenderPlatform pb_enum.PLATFORM_TYPE       `protobuf:"varint,16,opt,name=sender_platform,json=senderPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"sender_platform,omitempty"` // 发红包平台
}

func (x *RedEnvelopeInfo) Reset() {
	*x = RedEnvelopeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedEnvelopeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedEnvelopeInfo) ProtoMessage() {}

func (x *RedEnvelopeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedEnvelopeInfo.ProtoReflect.Descriptor instead.
func (*RedEnvelopeInfo) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{0}
}

func (x *RedEnvelopeInfo) GetEnvId() int64 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *RedEnvelopeInfo) GetEnvType() pb_enum.RED_ENVELOPE_TYPE {
	if x != nil {
		return x.EnvType
	}
	return pb_enum.RED_ENVELOPE_TYPE(0)
}

func (x *RedEnvelopeInfo) GetReceiverType() pb_enum.RECEIVER_TYPE {
	if x != nil {
		return x.ReceiverType
	}
	return pb_enum.RECEIVER_TYPE(0)
}

func (x *RedEnvelopeInfo) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RedEnvelopeInfo) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *RedEnvelopeInfo) GetSenderUid() int64 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *RedEnvelopeInfo) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RedEnvelopeInfo) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *RedEnvelopeInfo) GetRemainQuantity() int32 {
	if x != nil {
		return x.RemainQuantity
	}
	return 0
}

func (x *RedEnvelopeInfo) GetRemainAmount() int64 {
	if x != nil {
		return x.RemainAmount
	}
	return 0
}

func (x *RedEnvelopeInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RedEnvelopeInfo) GetReceiverUids() []int64 {
	if x != nil {
		return x.ReceiverUids
	}
	return nil
}

func (x *RedEnvelopeInfo) GetExpiredTs() int64 {
	if x != nil {
		return x.ExpiredTs
	}
	return 0
}

func (x *RedEnvelopeInfo) GetPayStatus() pb_enum.PAYMENT_STATUS {
	if x != nil {
		return x.PayStatus
	}
	return pb_enum.PAYMENT_STATUS(0)
}

func (x *RedEnvelopeInfo) GetEnvStatus() pb_enum.RED_ENVELOPE_STATUS {
	if x != nil {
		return x.EnvStatus
	}
	return pb_enum.RED_ENVELOPE_STATUS(0)
}

func (x *RedEnvelopeInfo) GetSenderPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.SenderPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

type RedEnvelopeReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvId     int64 `protobuf:"varint,1,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`             // 红包ID
	ExpiredTs int64 `protobuf:"varint,2,opt,name=expired_ts,json=expiredTs,proto3" json:"expired_ts,omitempty"` // 红包过期时间
}

func (x *RedEnvelopeReturn) Reset() {
	*x = RedEnvelopeReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedEnvelopeReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedEnvelopeReturn) ProtoMessage() {}

func (x *RedEnvelopeReturn) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedEnvelopeReturn.ProtoReflect.Descriptor instead.
func (*RedEnvelopeReturn) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{1}
}

func (x *RedEnvelopeReturn) GetEnvId() int64 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *RedEnvelopeReturn) GetExpiredTs() int64 {
	if x != nil {
		return x.ExpiredTs
	}
	return 0
}

type GiveRedEnvelopeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvType        pb_enum.RED_ENVELOPE_TYPE `protobuf:"varint,1,opt,name=env_type,json=envType,proto3,enum=pb_enum.RED_ENVELOPE_TYPE" json:"env_type,omitempty"`                  // 红包类型 1-均分红包 2-碰运气红包
	ReceiverType   pb_enum.RECEIVER_TYPE     `protobuf:"varint,2,opt,name=receiver_type,json=receiverType,proto3,enum=pb_enum.RECEIVER_TYPE" json:"receiver_type,omitempty"`       // 接收者类型 1-私聊对方 2-群聊所有人 3-群聊指定人
	ChatId         int64                     `protobuf:"varint,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`                                                    // 聊天ID
	SenderUid      int64                     `protobuf:"varint,4,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`                                           // 发红包用户ID
	SenderPlatform pb_enum.PLATFORM_TYPE     `protobuf:"varint,5,opt,name=sender_platform,json=senderPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"sender_platform,omitempty"` // 发红包平台
	Total          int64                     `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`                                                                    // 红包总金额(分)
	Quantity       int32                     `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`                                                              // 红包数量
	Message        string                    `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`                                                                 // 祝福语
	ReceiverUids   []int64                   `protobuf:"varint,9,rep,packed,name=receiver_uids,json=receiverUids,proto3" json:"receiver_uids,omitempty"`                           // 接收者ID
	PayPassword    string                    `protobuf:"bytes,10,opt,name=pay_password,json=payPassword,proto3" json:"pay_password,omitempty"`                                     // 支付密码
}

func (x *GiveRedEnvelopeReq) Reset() {
	*x = GiveRedEnvelopeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveRedEnvelopeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveRedEnvelopeReq) ProtoMessage() {}

func (x *GiveRedEnvelopeReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveRedEnvelopeReq.ProtoReflect.Descriptor instead.
func (*GiveRedEnvelopeReq) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{2}
}

func (x *GiveRedEnvelopeReq) GetEnvType() pb_enum.RED_ENVELOPE_TYPE {
	if x != nil {
		return x.EnvType
	}
	return pb_enum.RED_ENVELOPE_TYPE(0)
}

func (x *GiveRedEnvelopeReq) GetReceiverType() pb_enum.RECEIVER_TYPE {
	if x != nil {
		return x.ReceiverType
	}
	return pb_enum.RECEIVER_TYPE(0)
}

func (x *GiveRedEnvelopeReq) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GiveRedEnvelopeReq) GetSenderUid() int64 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *GiveRedEnvelopeReq) GetSenderPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.SenderPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *GiveRedEnvelopeReq) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GiveRedEnvelopeReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GiveRedEnvelopeReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GiveRedEnvelopeReq) GetReceiverUids() []int64 {
	if x != nil {
		return x.ReceiverUids
	}
	return nil
}

func (x *GiveRedEnvelopeReq) GetPayPassword() string {
	if x != nil {
		return x.PayPassword
	}
	return ""
}

type GiveRedEnvelopeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GiveRedEnvelopeResp) Reset() {
	*x = GiveRedEnvelopeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiveRedEnvelopeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiveRedEnvelopeResp) ProtoMessage() {}

func (x *GiveRedEnvelopeResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiveRedEnvelopeResp.ProtoReflect.Descriptor instead.
func (*GiveRedEnvelopeResp) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{3}
}

func (x *GiveRedEnvelopeResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GiveRedEnvelopeResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RedEnvelopePayCallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvId         int64                  `protobuf:"varint,1,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`                                         // 红包ID
	TradeNo       string                 `protobuf:"bytes,2,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`                                    // 交易编号
	PaymentAmount int64                  `protobuf:"varint,3,opt,name=payment_amount,json=paymentAmount,proto3" json:"payment_amount,omitempty"`                 // 支付金额
	PayStatus     pb_enum.PAYMENT_STATUS `protobuf:"varint,4,opt,name=pay_status,json=payStatus,proto3,enum=pb_enum.PAYMENT_STATUS" json:"pay_status,omitempty"` // 支付状态 0-未支付 1-支付中 2-已支付 3-支付失败
}

func (x *RedEnvelopePayCallbackReq) Reset() {
	*x = RedEnvelopePayCallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedEnvelopePayCallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedEnvelopePayCallbackReq) ProtoMessage() {}

func (x *RedEnvelopePayCallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedEnvelopePayCallbackReq.ProtoReflect.Descriptor instead.
func (*RedEnvelopePayCallbackReq) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{4}
}

func (x *RedEnvelopePayCallbackReq) GetEnvId() int64 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *RedEnvelopePayCallbackReq) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RedEnvelopePayCallbackReq) GetPaymentAmount() int64 {
	if x != nil {
		return x.PaymentAmount
	}
	return 0
}

func (x *RedEnvelopePayCallbackReq) GetPayStatus() pb_enum.PAYMENT_STATUS {
	if x != nil {
		return x.PayStatus
	}
	return pb_enum.PAYMENT_STATUS(0)
}

type RedEnvelopePayCallbackResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RedEnvelopePayCallbackResp) Reset() {
	*x = RedEnvelopePayCallbackResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedEnvelopePayCallbackResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedEnvelopePayCallbackResp) ProtoMessage() {}

func (x *RedEnvelopePayCallbackResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedEnvelopePayCallbackResp.ProtoReflect.Descriptor instead.
func (*RedEnvelopePayCallbackResp) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{5}
}

func (x *RedEnvelopePayCallbackResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RedEnvelopePayCallbackResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type OpenRedEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvId          int64                 `protobuf:"varint,1,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`                                                        // 红包ID
	Status         string                `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                                                                    // 成功/失败
	Desc           string                `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`                                                                        // 描述
	RemainAmount   int64                 `protobuf:"varint,4,opt,name=remain_amount,json=remainAmount,proto3" json:"remain_amount,omitempty"`                                   // 剩余红包金额(分)
	RemainQuantity int64                 `protobuf:"varint,5,opt,name=remain_quantity,json=remainQuantity,proto3" json:"remain_quantity,omitempty"`                             // 剩余红包数量
	ReceiveAmount  int64                 `protobuf:"varint,6,opt,name=receive_amount,json=receiveAmount,proto3" json:"receive_amount,omitempty"`                                // 领取红包金额(分)
	RecordId       int64                 `protobuf:"varint,7,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`                                               // 领取记录ID
	TradeNo        string                `protobuf:"bytes,8,opt,name=trade_no,json=tradeNo,proto3" json:"trade_no,omitempty"`                                                   // 交易编号
	SenderUid      int64                 `protobuf:"varint,9,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`                                            // 发红包用户ID
	SenderPlatform pb_enum.PLATFORM_TYPE `protobuf:"varint,10,opt,name=sender_platform,json=senderPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"sender_platform,omitempty"` // 发红包平台
	ReceiverUid    int64                 `protobuf:"varint,11,opt,name=receiver_uid,json=receiverUid,proto3" json:"receiver_uid,omitempty"`                                     // 接收者ID
	ChatId         int64                 `protobuf:"varint,12,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`                                                    // 聊天ID
}

func (x *OpenRedEnvelope) Reset() {
	*x = OpenRedEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_red_env_red_env_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRedEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRedEnvelope) ProtoMessage() {}

func (x *OpenRedEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_pb_red_env_red_env_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRedEnvelope.ProtoReflect.Descriptor instead.
func (*OpenRedEnvelope) Descriptor() ([]byte, []int) {
	return file_pb_red_env_red_env_proto_rawDescGZIP(), []int{6}
}

func (x *OpenRedEnvelope) GetEnvId() int64 {
	if x != nil {
		return x.EnvId
	}
	return 0
}

func (x *OpenRedEnvelope) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OpenRedEnvelope) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *OpenRedEnvelope) GetRemainAmount() int64 {
	if x != nil {
		return x.RemainAmount
	}
	return 0
}

func (x *OpenRedEnvelope) GetRemainQuantity() int64 {
	if x != nil {
		return x.RemainQuantity
	}
	return 0
}

func (x *OpenRedEnvelope) GetReceiveAmount() int64 {
	if x != nil {
		return x.ReceiveAmount
	}
	return 0
}

func (x *OpenRedEnvelope) GetRecordId() int64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *OpenRedEnvelope) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *OpenRedEnvelope) GetSenderUid() int64 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *OpenRedEnvelope) GetSenderPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.SenderPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *OpenRedEnvelope) GetReceiverUid() int64 {
	if x != nil {
		return x.ReceiverUid
	}
	return 0
}

func (x *OpenRedEnvelope) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

var File_pb_red_env_red_env_proto protoreflect.FileDescriptor

var file_pb_red_env_red_env_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x2f, 0x72, 0x65, 0x64,
	0x5f, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x62, 0x5f, 0x72,
	0x65, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x05, 0x0a, 0x0f, 0x52,
	0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15,
	0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x52, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0d,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x45,
	0x43, 0x45, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0c, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x09, 0x70, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x6e, 0x76, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x52, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3f, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x22, 0x49, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x73, 0x22, 0x95, 0x03, 0x0a, 0x12,
	0x47, 0x69, 0x76, 0x65, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x52,
	0x45, 0x44, 0x5f, 0x45, 0x4e, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x45, 0x43, 0x45, 0x49,
	0x56, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x3f,
	0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52,
	0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x69, 0x76, 0x65, 0x52, 0x65, 0x64, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0xac, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x50, 0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x15,
	0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x52, 0x09, 0x70, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x42, 0x0a, 0x1a, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x50, 0x61,
	0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x9d, 0x03, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x64, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x32, 0x5c, 0x0a, 0x06, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x12, 0x52, 0x0a,
	0x0f, 0x47, 0x69, 0x76, 0x65, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x12, 0x1e, 0x2e, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x2e, 0x47, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x2e, 0x47, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x76, 0x3b, 0x70, 0x62, 0x5f, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_red_env_red_env_proto_rawDescOnce sync.Once
	file_pb_red_env_red_env_proto_rawDescData = file_pb_red_env_red_env_proto_rawDesc
)

func file_pb_red_env_red_env_proto_rawDescGZIP() []byte {
	file_pb_red_env_red_env_proto_rawDescOnce.Do(func() {
		file_pb_red_env_red_env_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_red_env_red_env_proto_rawDescData)
	})
	return file_pb_red_env_red_env_proto_rawDescData
}

var file_pb_red_env_red_env_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_red_env_red_env_proto_goTypes = []interface{}{
	(*RedEnvelopeInfo)(nil),            // 0: pb_red_env.RedEnvelopeInfo
	(*RedEnvelopeReturn)(nil),          // 1: pb_red_env.RedEnvelopeReturn
	(*GiveRedEnvelopeReq)(nil),         // 2: pb_red_env.GiveRedEnvelopeReq
	(*GiveRedEnvelopeResp)(nil),        // 3: pb_red_env.GiveRedEnvelopeResp
	(*RedEnvelopePayCallbackReq)(nil),  // 4: pb_red_env.RedEnvelopePayCallbackReq
	(*RedEnvelopePayCallbackResp)(nil), // 5: pb_red_env.RedEnvelopePayCallbackResp
	(*OpenRedEnvelope)(nil),            // 6: pb_red_env.OpenRedEnvelope
	(pb_enum.RED_ENVELOPE_TYPE)(0),     // 7: pb_enum.RED_ENVELOPE_TYPE
	(pb_enum.RECEIVER_TYPE)(0),         // 8: pb_enum.RECEIVER_TYPE
	(pb_enum.PAYMENT_STATUS)(0),        // 9: pb_enum.PAYMENT_STATUS
	(pb_enum.RED_ENVELOPE_STATUS)(0),   // 10: pb_enum.RED_ENVELOPE_STATUS
	(pb_enum.PLATFORM_TYPE)(0),         // 11: pb_enum.PLATFORM_TYPE
}
var file_pb_red_env_red_env_proto_depIdxs = []int32{
	7,  // 0: pb_red_env.RedEnvelopeInfo.env_type:type_name -> pb_enum.RED_ENVELOPE_TYPE
	8,  // 1: pb_red_env.RedEnvelopeInfo.receiver_type:type_name -> pb_enum.RECEIVER_TYPE
	9,  // 2: pb_red_env.RedEnvelopeInfo.pay_status:type_name -> pb_enum.PAYMENT_STATUS
	10, // 3: pb_red_env.RedEnvelopeInfo.env_status:type_name -> pb_enum.RED_ENVELOPE_STATUS
	11, // 4: pb_red_env.RedEnvelopeInfo.sender_platform:type_name -> pb_enum.PLATFORM_TYPE
	7,  // 5: pb_red_env.GiveRedEnvelopeReq.env_type:type_name -> pb_enum.RED_ENVELOPE_TYPE
	8,  // 6: pb_red_env.GiveRedEnvelopeReq.receiver_type:type_name -> pb_enum.RECEIVER_TYPE
	11, // 7: pb_red_env.GiveRedEnvelopeReq.sender_platform:type_name -> pb_enum.PLATFORM_TYPE
	9,  // 8: pb_red_env.RedEnvelopePayCallbackReq.pay_status:type_name -> pb_enum.PAYMENT_STATUS
	11, // 9: pb_red_env.OpenRedEnvelope.sender_platform:type_name -> pb_enum.PLATFORM_TYPE
	2,  // 10: pb_red_env.RedEnv.GiveRedEnvelope:input_type -> pb_red_env.GiveRedEnvelopeReq
	3,  // 11: pb_red_env.RedEnv.GiveRedEnvelope:output_type -> pb_red_env.GiveRedEnvelopeResp
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pb_red_env_red_env_proto_init() }
func file_pb_red_env_red_env_proto_init() {
	if File_pb_red_env_red_env_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_red_env_red_env_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedEnvelopeInfo); i {
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
		file_pb_red_env_red_env_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedEnvelopeReturn); i {
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
		file_pb_red_env_red_env_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveRedEnvelopeReq); i {
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
		file_pb_red_env_red_env_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiveRedEnvelopeResp); i {
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
		file_pb_red_env_red_env_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedEnvelopePayCallbackReq); i {
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
		file_pb_red_env_red_env_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedEnvelopePayCallbackResp); i {
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
		file_pb_red_env_red_env_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRedEnvelope); i {
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
			RawDescriptor: file_pb_red_env_red_env_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_red_env_red_env_proto_goTypes,
		DependencyIndexes: file_pb_red_env_red_env_proto_depIdxs,
		MessageInfos:      file_pb_red_env_red_env_proto_msgTypes,
	}.Build()
	File_pb_red_env_red_env_proto = out.File
	file_pb_red_env_red_env_proto_rawDesc = nil
	file_pb_red_env_red_env_proto_goTypes = nil
	file_pb_red_env_red_env_proto_depIdxs = nil
}
