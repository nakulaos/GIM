package service

import (
	gw_client "GIM/apps/msg_gateway/client"
	"GIM/pkg/common/xetcd"
	"GIM/pkg/common/xkafka"
	"GIM/pkg/common/xlog"
	"GIM/pkg/conf"
	"GIM/pkg/utils"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

func (s *distService) watchMsgGatewayServer() {
	var (
		catalog = s.cfg.Etcd.Schema + ":///" + s.cfg.MsgGatewayServer.Name
		members = map[string]string{}
		err     error
	)
	s.watcher, err = xetcd.NewWatcher(catalog, s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, s.changeWatcher)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	err = s.watcher.Run()
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	s.watcher.EachKvs(func(k string, v *xetcd.KeyValue) bool {
		var (
			name, port       = utils.GetServer(k)
			portVal, _       = utils.ToInt(port)
			member           = name + ":" + cast.ToString(portVal+1)
			_, serverId, _   = utils.GetMsgGatewayServer(member)
			msgGatewayServer *conf.GrpcServer
		)
		members[member] = member

		msgGatewayServer = &conf.GrpcServer{}
		copier.Copy(msgGatewayServer, s.cfg.MsgGatewayServer)
		msgGatewayServer.Name = name

		cli := gw_client.NewMsgGwClient(s.cfg.Etcd, msgGatewayServer, s.cfg.GrpcServer.Jaeger, s.cfg.Name)
		producer := xkafka.NewKafkaProducer(s.cfg.GwMsgProducer.Address, s.cfg.GwMsgProducer.Topic+"_"+cast.ToString(serverId))

		s.mutex.Lock()
		s.clients[serverId] = cli
		s.producers[serverId] = producer
		s.mutex.Unlock()
		return true
	})
	s.watcher.Watch()
}

func (s *distService) changeWatcher(kv *xetcd.KeyValue, eventType int) {
	var (
		name, port     = utils.GetServer(kv.Key)
		portVal, _     = utils.ToInt(port)
		member         = name + ":" + cast.ToString(portVal+1)
		_, serverId, _ = utils.GetMsgGatewayServer(member)
		cli            gw_client.MessageGatewayClient
		producer       *xkafka.Producer
		ok             bool
	)
	switch eventType {
	case xetcd.EVENT_TYPE_PUT:
		s.mutex.RLock()
		_, ok = s.clients[serverId]
		s.mutex.RUnlock()

		if ok == false {
			cli = s.NewPushOnlineClient(name)
			producer = xkafka.NewKafkaProducer(s.cfg.GwMsgProducer.Address, s.cfg.GwMsgProducer.Topic+"_"+cast.ToString(serverId))

			s.mutex.Lock()
			s.clients[serverId] = cli
			s.producers[serverId] = producer
			s.mutex.Unlock()
		}
	case xetcd.EVENT_TYPE_DELETE:
		s.mutex.Lock()
		if cli, ok = s.clients[serverId]; ok {
			delete(s.clients, serverId)
		}
		if producer, ok = s.producers[serverId]; ok {
			delete(s.producers, serverId)
		}
		s.mutex.Unlock()
		if producer != nil {
			producer.Close()
		}
	}
}

func (s *distService) NewPushOnlineClient(serverName string) (cli gw_client.MessageGatewayClient) {
	msgGatewayServer := &conf.GrpcServer{}
	copier.Copy(msgGatewayServer, s.cfg.MsgGatewayServer)
	msgGatewayServer.Name = serverName
	cli = gw_client.NewMsgGwClient(s.cfg.Etcd, msgGatewayServer, s.cfg.GrpcServer.Jaeger, s.cfg.Name)
	return cli
}
