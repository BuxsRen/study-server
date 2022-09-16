package config

var App *app

// 服务配置
type server struct {
	Mode string `yaml:"mode"` // 运行模式
	Url string `yaml:"url"` // 项目地址
	Host string `yaml:"host"` // 服务监听地址
	Port string `yaml:"port"`// 服务监听端口
	GrpcPort string `yaml:"grpc_port"`// grpc服务监听端口
	Ip string `yaml:"ip"`// 对外ip
	Debug bool `yaml:"debug"`// 开启debug模式
	Env string `yaml:"env"`// 运行环境 local(开发) production(线上)
	LogAccess string `yaml:"log_access"`// 访问日志保存地址
	LogError string`yaml:"log_error"` // 错误日志保存地址
	Template bool `yaml:"template"`// 加载模板
	ManageWs string `yaml:"manage_ws"`// 管理服务器地址
	ManageRpc string `yaml:"manage_rpc"`// 管理服务器地址
}

// MySQL配置
type mysql struct {
	Host string `yaml:"host"`// 数据库地址
	Port string `yaml:"port"`// 数据库端口
	Database string `yaml:"database"`// 数据库名
	UserName string `yaml:"username"`// 用户名
	PassWord string `yaml:"password"`// 密码
	Prefix string `yaml:"prefix"`// 表前缀
	Log bool `yaml:"log"`// 开启sql日志，打印sql执行日志到控制台
	SaveLog bool `yaml:"save_log"`// 保存sql日志到文本，需要先打开 "开启sql日志"
	LogPath string `yaml:"log_path"`// sql 日志保存路径，需要先打开 "保存sql日志到文本"
	Pool int `yaml:"pool"` // 连接池数量
	Ping int `yaml:"ping"` // 心跳时间
}

// Redis 配置
type redis struct {
	Host string `yaml:"host"`// 地址
	Port string `yaml:"port"`// 端口
	PassWord string `yaml:"password"`// 密码
	Prefix string `yaml:"prefix"`// redis前缀
	Pool int `yaml:"pool"` // 连接池数量
	Ping int `yaml:"ping"` // 心跳时间
	Db int `yaml:"db"` // 使用几号数据库
}

type rsa struct {
	PublicKey string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

// 其他配置
type other struct {
	PublicDir string `yaml:"public_dir"`// 静态文件保存目录，后面一定要加上 /
	PublicPrefix string `yaml:"public_prefix"`// 前端寻址前缀
	TokenKey string `yaml:"token_key"`// 接口token签发密钥
	AesKey string `yaml:"aes_key"`// AES 16位加密字符串密码
}

// 推送配置
type push struct {
	Use bool `yaml:"use"`// 开启异常推送
	Mode string `yaml:"mode"`// 推送方式：bark、dingTalk、dingTalkMarkDown、wechat
	BarkUrl string `yaml:"bark_url"`//  bark推送地址
	DingTalkUrl string `yaml:"ding_talk_url"`// 钉钉推送地址
	WechatUrl string `yaml:"wechat_url"`// 企业微信推送地址
}

// 七牛云配置
type qiniuyun struct {
	Bucket string `yaml:"bucket"`// 七牛云空间名称
	AccessKey string `yaml:"access_key"`// 七牛云AK
	SecretKey string `yaml:"secret_key"`// 七牛云SK
}

// 邮箱配置
type email struct {
	Name string `yaml:"name"`// 发件人名称
	User string `yaml:"user"`// 发件人邮箱
	Pass string `yaml:"pass"`// 发件人密码
	Host string `yaml:"host"`// 邮箱服务器
	Port int `yaml:"port"`// 邮箱端口
}

// 支付宝网页&移动应用支付配置
type alipay struct {
	AppID string `yaml:"appid"`// AppID
	PrivateKey string `yaml:"private_key"`// 应用私钥地址
	NotifyUrl string `yaml:"notify_url"`// 支付回调地址
}

type app struct {
	Server server `yaml:"server"`
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
	RSA rsa `yaml:"rsa"`
	Other other `yaml:"other"`
	Push push `yaml:"push"`
	QiNiuYun qiniuyun `yaml:"qiniuyun"`
	Email email `yaml:"email"`
	Alipay alipay `yaml:"alipay"`
}