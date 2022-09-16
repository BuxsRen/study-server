package match

import (
	"context"
	"study-server/app/socket/gateway/gateway"
	pm "study-server/bootstrap/grpc/proto/match"
)

// 获取用户区域内信息(获取区域内所有用户id)，返回区域名称，区域内用户列表
func (m *Match) GetAreaInfo(id string) (*pm.List, error) {
	list, e := (*m.service).GetAreaInfo(
		context.Background(),
		&pm.Req{
			Id:      id,
			Gateway: gateway.DeviceNumber,
		},
	)
	if e != nil {
		return nil, e
	}
	return list, e
}

// 获取用户区域内信息(获取区域内所有用户id)，返回区域名称，区域内用户列表
func (m *Match) GetArea(area string) (*pm.List, error) {
	list, e := (*m.service).GetAreaInfo(
		context.Background(),
		&pm.Req{
			Area:    area,
			Gateway: gateway.DeviceNumber,
		},
	)
	if e != nil {
		return nil, e
	}
	return list, e
}

// 获取用户区域内信息(获取区域内所有用户id)，返回区域名称，区域内用户列表
func (m *Match) GetUserInfo(id string) (*pm.Res, error) {
	res, e := (*m.service).GetUserInfo(
		context.Background(),
		&pm.Req{
			Id: id,
		},
	)
	if e != nil {
		return nil, e
	}
	return res, e
}