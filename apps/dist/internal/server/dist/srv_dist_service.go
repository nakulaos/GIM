package dist

import (
	"GIM/pkg/proto/pb_dist"
	"context"
)

func (s *distServer) ChatInviteNotification(ctx context.Context, req *pb_dist.ChatInviteNotificationReq) (resp *pb_dist.ChatInviteNotificationResp, err error) {
	return s.distService.ChatInviteNotification(ctx, req)
}
