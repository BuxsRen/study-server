package grpc

import (
	"context"
	"errors"
	"study-server/app/libs/utils"
	"study-server/app/socket/manage/manage"
	"study-server/bootstrap/area"
	pm "study-server/bootstrap/grpc/proto/manage"
	"sync"
)

var _ pm.ManageGrpcServer = (*Server)(nil)

// 获取指定服务中的节点列表
func (s *Server) GetService(_ context.Context, req *pm.Node) (*pm.List, error) {
	utils.Println("[manage] 获取服务", req.Mode)
	group, e := manage.WebSocket.GetAreaGroup(req.Mode)
	if e != nil {
		return nil, errors.New("服务不存在")
	}
	var node []*pm.Node
	group.Range(func(key, value interface{}) bool {
		list := value.(*sync.Map)
		list.Range(func(key, value interface{}) bool {
			v := value.(*area.Client)
			node = append(node, &pm.Node{
				Id:   v.Id,
				Mode: v.Group,
				Rpc:  v.Rpc,
				Udp:  v.Udp,
			})
			return true

		})
		return true
	})
	return &pm.List{List: node}, nil
}