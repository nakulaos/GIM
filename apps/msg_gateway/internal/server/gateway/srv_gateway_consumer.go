package gateway

import (
	"GIM/pkg/proto/pb_gw"
	"context"
	"google.golang.org/protobuf/proto"
)

func (s *gatewayServer) MessageHandler(msg []byte, msgKey string) (err error) {
	var req = new(pb_gw.SendTopicMessageReq)
	err = proto.Unmarshal(msg, req)
	if err != nil {
		err = nil
		return
	}
	s.SendTopicMessage(context.Background(), req)
	return
}
