package push

import (
	"errors"
	"study-server/bootstrap/config"
)

// 推送
type push struct {
	mode string // 推送方式
}

// 消息推送初始化
/**
 * @Example:
	p := push.New()
	p.Push("test")
*/
func New() *push {
	p := new(push)
	p.mode = config.App.Push.Mode
	if config.App.Push.Mode == "" {
		p.mode = "bark"
	}
	return p
}

// 推送信息。推送的内容
func (p *push) Push(message string) error {
	if fun[p.mode] == nil {
		return errors.New("不支持的推送方式")
	}
	err := fun[p.mode](message)
	if err != nil {
		return err
	}
	return nil
}

// 更改推送方式，只对当前实例有效。模式(bark，dingTalk，dingTalkMarkDown，wechat)
func (p *push) SetPushMode(mode string) {
	p.mode = mode
}
