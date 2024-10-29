package service

import (
	"GIM/domain/po"
	"GIM/pkg/common/xlog"
	"GIM/pkg/proto/pb_chat_msg"
	"GIM/pkg/proto/pb_enum"
	"GIM/pkg/proto/pb_msg"
	"GIM/pkg/utils"
	"context"
	"github.com/jinzhu/copier"
	"sort"
)

func (s *chatMessageService) GetChatMessageList(ctx context.Context, req *pb_chat_msg.GetChatMessageListReq) (resp *pb_chat_msg.GetChatMessageListResp, _ error) {

	/*
		1. 确定最大的消息边界
		2. 看看请求的时间有没有超过缓存的,没有就从缓存拿
		3. 缓存拿完之后，缺的从数据库拿
		4. 返回最新的seq
	*/
	resp = &pb_chat_msg.GetChatMessageListResp{Msgs: &pb_chat_msg.ChatMessages{List: make([]*pb_msg.SrvChatMessage, 0)}}
	var (
		length    = len(req.SeqIds)
		nowTs     = utils.NowUnix()
		list      []*po.Message
		cacheList []*po.Message
		maxSeqId  uint64
		err       error
	)
	defer func() {
		if resp.Msgs != nil && len(resp.Msgs.List) > 0 {
			sort.Slice(resp.Msgs.List, func(i, j int) bool {
				return resp.Msgs.List[i].SeqId < resp.Msgs.List[j].SeqId
			})
		}
	}()

	if req.Order == pb_enum.ORDER_TYPE_ASC {
		// 1、消息边界
		maxSeqId, err = s.chatMessageCache.GetMaxSeqID(req.ChatId)
		if err != nil {
			return
		}
		resp.Msgs.LastSeqId = int64(maxSeqId)
	}
	// 2、从redis缓存中读取
	if nowTs-req.MsgTs < s.cfg.MsgCache.L1Duration {
		cacheList, err = s.GetCacheChatMessageList(req)
		if err != nil {
			resp.Set(ERROR_CODE_CHAT_MSG_REDIS_GET_FAILED, ERROR_CHAT_MSG_REDIS_GET_FAILED)
			xlog.Warn(resp.Code, resp.Msg, err.Error())
			return
		}
		if len(cacheList) == length {
			// 全部读取缓存
			copier.Copy(&resp.Msgs.List, cacheList)
			return
		}
	}
	// 3、从mysql数据库中读取
	list, err = s.GetHistoryMessageList(req)
	if err != nil {
		resp.Set(ERROR_CODE_CHAT_MSG_QUERY_DB_FAILED, ERROR_CHAT_MSG_QUERY_DB_FAILED)
		xlog.Warn(resp.Code, resp.Msg, err.Error())
		return
	}
	if len(cacheList) > 0 && len(list) > 0 {
		list = append(list, cacheList...)
		copier.Copy(&resp.Msgs.List, list)
		return
	}
	if len(list) > 0 {
		copier.Copy(&resp.Msgs.List, list)
		return
	}
	if len(cacheList) > 0 {
		copier.Copy(&resp.Msgs.List, cacheList)
		return
	}
	return
}

// 弃用
//func (s *chatMessageService) GetChatMessages(_ context.Context, req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp, _ error) {
//	resp = &pb_chat_msg.GetChatMessagesResp{List: make([]*pb_msg.SrvChatMessage, 0)}
//	var (
//		nowTs     = utils.NowUnix()
//		list      = make([]*po.Message, 0)
//		cacheList []*po.Message
//		//hotList     []*po.Message
//		historyList []*po.Message
//		msgCount    int
//		maxSeqId    uint64
//		next        bool
//		err         error
//	)
//	defer func() {
//		if req.New == true && len(resp.List) > 0 {
//			sort.Slice(resp.List, func(i, j int) bool {
//				return resp.List[i].SeqId < resp.List[j].SeqId
//			})
//		}
//	}()
//	// 1、消息边界
//	maxSeqId, _ = s.chatMessageCache.GetMaxSeqID(req.ChatId)
//	if req.SeqId >= int64(maxSeqId) {
//		if req.New == true {
//			// 1.1 消息越界
//			return
//		}
//		req.SeqId = int64(maxSeqId)
//	}
//
//	if nowTs-req.MsgTs < s.cfg.MsgCache.L1Duration {
//		// 2、从redis缓存中读取
//		cacheList, next, err = s.GetCacheMessages(req, int64(maxSeqId))
//		if next == false || err != nil {
//			if len(cacheList) > 0 {
//				copier.Copy(&resp.List, cacheList)
//			}
//			return
//		}
//	}
//	/*
//		if nowTs-req.MsgTs < s.cfg.MsgCache.L2Duration {
//			// 3、从mongo缓存中读取
//			hotList, next, err = s.GetHotMessages(req, int64(maxSeqId))
//			if next == false || err != nil {
//				if len(cacheList) > 0 {
//					hotList = append(cacheList, hotList...)
//					if req.New == false {
//						sortMessageList(hotList, true)
//					}
//				}
//				if len(hotList) > 0 {
//					copier.Copy(&resp.List, hotList)
//				}
//				return
//			}
//		}
//	*/
//
//	// 4、从mysql缓存中读取
//	historyList, err = s.GetHistoryMessages(req)
//	if err != nil {
//		return
//	}
//
//	if len(cacheList) > 0 {
//		list = append(list, cacheList...)
//	}
//	/*
//		if len(hotList) > 0 {
//			list = append(list, hotList...)
//		}
//	*/
//	if len(historyList) > 0 {
//		list = append(list, historyList...)
//	}
//
//	msgCount = len(list)
//	if msgCount == 0 {
//		return
//	}
//	copier.Copy(&resp.List, list)
//	return
//}

//func sortMessageList(list []*po.Message, asc bool) {
//	sort.Slice(list, func(i, j int) bool {
//		if asc == true {
//			return list[i].SeqId < list[j].SeqId
//		} else {
//			return list[i].SeqId > list[j].SeqId
//		}
//	})
//}
