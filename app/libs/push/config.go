package push

var fun = map[string]func(message string) error {
	"bark": bark,
	"dingTalk": dingTalk,
	"dingTalkMarkDown": dingTalkMarkDown,
	"wechat": wechat,
}