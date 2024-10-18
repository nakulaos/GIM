// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pb_chat_member/chat_member.proto

package pb_chat_member

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatMember_GetDistMemberList_FullMethodName   = "/pb_chat_member.ChatMember/GetDistMemberList"
	ChatMember_GetChatMemberInfo_FullMethodName   = "/pb_chat_member.ChatMember/GetChatMemberInfo"
	ChatMember_ChatMemberOnOffLine_FullMethodName = "/pb_chat_member.ChatMember/ChatMemberOnOffLine"
	ChatMember_GetChatMemberList_FullMethodName   = "/pb_chat_member.ChatMember/GetChatMemberList"
	ChatMember_GetContactList_FullMethodName      = "/pb_chat_member.ChatMember/GetContactList"
	ChatMember_GetGroupChatList_FullMethodName    = "/pb_chat_member.ChatMember/GetGroupChatList"
)

// ChatMemberClient is the client API for ChatMember service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatMemberClient interface {
	GetDistMemberList(ctx context.Context, in *GetDistMemberListReq, opts ...grpc.CallOption) (*GetDistMemberListResp, error)
	GetChatMemberInfo(ctx context.Context, in *GetChatMemberInfoReq, opts ...grpc.CallOption) (*GetChatMemberInfoResp, error)
	ChatMemberOnOffLine(ctx context.Context, in *ChatMemberOnOffLineReq, opts ...grpc.CallOption) (*ChatMemberOnOffLineResp, error)
	GetChatMemberList(ctx context.Context, in *GetChatMemberListReq, opts ...grpc.CallOption) (*GetChatMemberListResp, error)
	GetContactList(ctx context.Context, in *GetContactListReq, opts ...grpc.CallOption) (*GetContactListResp, error)
	GetGroupChatList(ctx context.Context, in *GetGroupChatListReq, opts ...grpc.CallOption) (*GetGroupChatListResp, error)
}

type chatMemberClient struct {
	cc grpc.ClientConnInterface
}

func NewChatMemberClient(cc grpc.ClientConnInterface) ChatMemberClient {
	return &chatMemberClient{cc}
}

func (c *chatMemberClient) GetDistMemberList(ctx context.Context, in *GetDistMemberListReq, opts ...grpc.CallOption) (*GetDistMemberListResp, error) {
	out := new(GetDistMemberListResp)
	err := c.cc.Invoke(ctx, ChatMember_GetDistMemberList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberInfo(ctx context.Context, in *GetChatMemberInfoReq, opts ...grpc.CallOption) (*GetChatMemberInfoResp, error) {
	out := new(GetChatMemberInfoResp)
	err := c.cc.Invoke(ctx, ChatMember_GetChatMemberInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) ChatMemberOnOffLine(ctx context.Context, in *ChatMemberOnOffLineReq, opts ...grpc.CallOption) (*ChatMemberOnOffLineResp, error) {
	out := new(ChatMemberOnOffLineResp)
	err := c.cc.Invoke(ctx, ChatMember_ChatMemberOnOffLine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetChatMemberList(ctx context.Context, in *GetChatMemberListReq, opts ...grpc.CallOption) (*GetChatMemberListResp, error) {
	out := new(GetChatMemberListResp)
	err := c.cc.Invoke(ctx, ChatMember_GetChatMemberList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetContactList(ctx context.Context, in *GetContactListReq, opts ...grpc.CallOption) (*GetContactListResp, error) {
	out := new(GetContactListResp)
	err := c.cc.Invoke(ctx, ChatMember_GetContactList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatMemberClient) GetGroupChatList(ctx context.Context, in *GetGroupChatListReq, opts ...grpc.CallOption) (*GetGroupChatListResp, error) {
	out := new(GetGroupChatListResp)
	err := c.cc.Invoke(ctx, ChatMember_GetGroupChatList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatMemberServer is the server API for ChatMember service.
// All implementations must embed UnimplementedChatMemberServer
// for forward compatibility
type ChatMemberServer interface {
	GetDistMemberList(context.Context, *GetDistMemberListReq) (*GetDistMemberListResp, error)
	GetChatMemberInfo(context.Context, *GetChatMemberInfoReq) (*GetChatMemberInfoResp, error)
	ChatMemberOnOffLine(context.Context, *ChatMemberOnOffLineReq) (*ChatMemberOnOffLineResp, error)
	GetChatMemberList(context.Context, *GetChatMemberListReq) (*GetChatMemberListResp, error)
	GetContactList(context.Context, *GetContactListReq) (*GetContactListResp, error)
	GetGroupChatList(context.Context, *GetGroupChatListReq) (*GetGroupChatListResp, error)
	mustEmbedUnimplementedChatMemberServer()
}

// UnimplementedChatMemberServer must be embedded to have forward compatible implementations.
type UnimplementedChatMemberServer struct {
}

func (UnimplementedChatMemberServer) GetDistMemberList(context.Context, *GetDistMemberListReq) (*GetDistMemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistMemberList not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberInfo(context.Context, *GetChatMemberInfoReq) (*GetChatMemberInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberInfo not implemented")
}
func (UnimplementedChatMemberServer) ChatMemberOnOffLine(context.Context, *ChatMemberOnOffLineReq) (*ChatMemberOnOffLineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatMemberOnOffLine not implemented")
}
func (UnimplementedChatMemberServer) GetChatMemberList(context.Context, *GetChatMemberListReq) (*GetChatMemberListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatMemberList not implemented")
}
func (UnimplementedChatMemberServer) GetContactList(context.Context, *GetContactListReq) (*GetContactListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactList not implemented")
}
func (UnimplementedChatMemberServer) GetGroupChatList(context.Context, *GetGroupChatListReq) (*GetGroupChatListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupChatList not implemented")
}
func (UnimplementedChatMemberServer) mustEmbedUnimplementedChatMemberServer() {}

// UnsafeChatMemberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatMemberServer will
// result in compilation errors.
type UnsafeChatMemberServer interface {
	mustEmbedUnimplementedChatMemberServer()
}

func RegisterChatMemberServer(s grpc.ServiceRegistrar, srv ChatMemberServer) {
	s.RegisterService(&ChatMember_ServiceDesc, srv)
}

func _ChatMember_GetDistMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistMemberListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetDistMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_GetDistMemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetDistMemberList(ctx, req.(*GetDistMemberListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_GetChatMemberInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberInfo(ctx, req.(*GetChatMemberInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_ChatMemberOnOffLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMemberOnOffLineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).ChatMemberOnOffLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_ChatMemberOnOffLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).ChatMemberOnOffLine(ctx, req.(*ChatMemberOnOffLineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetChatMemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatMemberListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetChatMemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_GetChatMemberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetChatMemberList(ctx, req.(*GetChatMemberListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetContactList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetContactList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_GetContactList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetContactList(ctx, req.(*GetContactListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatMember_GetGroupChatList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupChatListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatMemberServer).GetGroupChatList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatMember_GetGroupChatList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatMemberServer).GetGroupChatList(ctx, req.(*GetGroupChatListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatMember_ServiceDesc is the grpc.ServiceDesc for ChatMember service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatMember_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_chat_member.ChatMember",
	HandlerType: (*ChatMemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDistMemberList",
			Handler:    _ChatMember_GetDistMemberList_Handler,
		},
		{
			MethodName: "GetChatMemberInfo",
			Handler:    _ChatMember_GetChatMemberInfo_Handler,
		},
		{
			MethodName: "ChatMemberOnOffLine",
			Handler:    _ChatMember_ChatMemberOnOffLine_Handler,
		},
		{
			MethodName: "GetChatMemberList",
			Handler:    _ChatMember_GetChatMemberList_Handler,
		},
		{
			MethodName: "GetContactList",
			Handler:    _ChatMember_GetContactList_Handler,
		},
		{
			MethodName: "GetGroupChatList",
			Handler:    _ChatMember_GetGroupChatList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chat_member/chat_member.proto",
}
