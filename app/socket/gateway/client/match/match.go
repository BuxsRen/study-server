package match

import (
	"context"
	"study-server/app/socket/gateway/gateway"
	pm "study-server/bootstrap/grpc/proto/match"
)

// 随机加入一个区域或者创建一个空的区域
func (m *Match) ChangeArea(id string) (*pm.Res, error) {
	res, e := (*m.service).ChangeArea(
		context.Background(),
		&pm.Req{
			Id:      id,
			Gateway: gateway.DeviceNumber,
		},
	)
	return res, e
}

// 退出区域
func (m *Match) ExitArea(id string) (*pm.Res, error) {
	res, e := (*m.service).ExitArea(
		context.Background(),
		&pm.Req{
			Id:      id,
			Gateway: gateway.DeviceNumber,
		},
	)
	return res, e
}

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
		return &pm.List{}, e
	}
	return list, e
}

