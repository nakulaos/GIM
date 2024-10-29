package message

import (
	"GIM/pkg/proto/pb_msg"
	"context"
)

func (s *messageServer) SendChatMessage(ctx context.Context, req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp, _ error) {
	return s.messageService.SendChatMessage(ctx, req)
}
