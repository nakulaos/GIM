package service

import (
	"GIM/pkg/entity"
	"GIM/pkg/proto/pb_convo"
	"GIM/pkg/proto/pb_enum"
	"context"
)

func (s *convoService) ConvoChatSeqList(ctx context.Context, req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp, _ error) {
	// 主要作用是返回指定uid的用户的状态消息读取列表
	resp = &pb_convo.ConvoChatSeqListResp{List: make([]*pb_convo.ConvoChatSeq, 0)}
	var (
		q = entity.NewNormalQuery()
	)
	q.SetFilter("m.uid=?", req.Uid)
	q.SetFilter("m.chat_id!=?", req.LastCid)
	q.SetFilter("m.status<=?", pb_enum.CHAT_STATUS_BANNED)
	q.SetFilter("m.deleted_ts=?", 0)
	q.SetFilter("c.deleted_ts=?", 0)
	q.SetFilter("c.srv_ts<=?", req.LastTs)
	q.SetLimit(req.Limit)
	resp.List, _ = s.chatMemberRepo.ConvoChatSeqList(q)
	return
}
