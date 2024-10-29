package ws

import (
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/utils"
	"bytes"
	"google.golang.org/protobuf/proto"
)

func (c *Client) overloadReply(buf []byte) {
	var (
		pkt    *pb_msg.Packet
		buffer = bytes.NewBuffer(buf)
	)
	pkt = utils.Decode(buffer, false)
	if pkt.Topic == pb_enum.TOPIC_UNKNOWN_TOPIC {
		return
	}
	c.messageReply(pkt, ERROR_CODE_WS_SERVER_OVERLOAD, ERROR_WS_SERVER_OVERLOAD)
}

func (c *Client) messageReply(pkt *pb_msg.Packet, code int32, msg string) {
	var (
		resp *pb_msg.MessageResp
		buf  []byte
		err  error
	)
	resp = &pb_msg.MessageResp{
		Code:  code,
		Msg:   msg,
		MsgId: pkt.MsgId,
		IsSrv: false,
	}
	buf, err = proto.Marshal(resp)
	if err != nil {
		return
	}
	buf, err = utils.Encode(int32(pkt.Topic), int32(pkt.SubTopic), int32(pb_enum.MESSAGE_TYPE_RESP), buf)
	if err != nil {
		return
	}
	c.Send(buf)
}
