package server

import ws "study-server/bootstrap/websocket/server"

// WebSocket 事件路由
func (gws *GatewayServer) Route() []*ws.Route {
	return []*ws.Route{
		{
			Action: "SendMsg",
			Fun:    gws.SendMsg,
		}, {
			Action: "SendMsgArea",
			Fun:    gws.SendMsgArea,
		}, {
			Action: "ChangeArea",
			Fun:    gws.ChangeArea,
		}, {
			Action: "GetAreaInfo",
			Fun:    gws.GetAreaInfo,
		},
	}
}
