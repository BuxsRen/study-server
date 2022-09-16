package server

import "encoding/json"

type ErrorMsg struct {
	Action  string `json:"action"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func (gws *GatewayServer) Error(action, msg string) []byte {
	b, _ := json.Marshal(&ErrorMsg{
		Action:  "Error",
		Method:  action,
		Message: msg,
	})
	return b
}
