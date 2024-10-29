package gateway

import (
	"GIM/pkg/proto/pb_cm"
	"fmt"
)

func (s *gatewayServer) SendCloudMessage(req *pb_cm.CloudMessageReq) {
	s.producer.EnQueue(req, fmt.Sprintf("%d:%d", req.Topic, req.SubTopic))
}
