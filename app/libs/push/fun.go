package push

import (
	"encoding/json"
	"net/url"
	"study-server/app/libs/net"
	"study-server/bootstrap/config"
)

// bark 推送
func bark(message string) error {
	n := net.New(config.App.Push.BarkUrl+url.QueryEscape(message), "GET", "")
	_, e := n.Do()
	return e
}

// 钉钉推送
func dingTalk(message string) error {
	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": message,
		},
	}
	param, _ := json.Marshal(data)
	n := net.New(config.App.Push.DingTalkUrl, "POST", string(param))
	_, e := n.SetHeader("Content-Type", "application/json;charset=utf-8").Do()
	return e
}

// 钉钉推送 markdown
func dingTalkMarkDown(message string) error {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": "标题",
			"text":  message,
		},
	}
	param, _ := json.Marshal(data)
	n := net.New(config.App.Push.DingTalkUrl, "POST", string(param))
	_, e := n.SetHeader("Content-Type", "application/json;charset=utf-8").Do()
	return e
}

// 企业微信推送
func wechat(message string) error {
	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": message,
		},
	}
	param, _ := json.Marshal(data)
	n := net.New(config.App.Push.WechatUrl, "POST", string(param))
	_, e := n.SetHeader("Content-Type", "application/json;charset=utf-8").Do()
	return e
}
