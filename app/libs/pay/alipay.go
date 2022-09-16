package pay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	_ "github.com/go-pay/gopay/alipay"
	_ "github.com/go-pay/gopay/pkg/xlog"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

// 支付宝支付
type ALiPay struct {
	Subject string          // 交易备注
	TradeNo string          // 单号
	Amount  float64         // 支付金额
	Context context.Context // 上下文
}

// 发起网页支付
/**
 * @Example:
	ctx := context.Background()
	alipay := pay.ALiPay{
		Subject: "测试",
		TradeNo: "TEST202201222309355702",
		Amount: 0.01,
		Context: ctx,
	}
	fmt.Println(alipay.WebPay()) // 返回支付地址
*/
func (a *ALiPay) WebPay() string {
	pay := a.init()
	//请求参数
	data := make(gopay.BodyMap)
	data.Set("subject", a.Subject)
	data.Set("out_trade_no", a.TradeNo)
	data.Set("total_amount", a.Amount)
	//手机网站支付请求
	payUrl, err := pay.TradeWapPay(a.Context, data)
	if err != nil {
		utils.ExitError(err.Error(), -1)
	}
	return payUrl
}

// 初始化支付
func (a *ALiPay) init() *alipay.Client {
	if a.Subject == "" || a.TradeNo == "" {
		utils.ExitError("交易备注或单号不能为空", -1)
	}
	if a.Amount <= 0 {
		utils.ExitError("支付金额应当大于0", -1)
	}
	appId := config.App.Alipay.AppID
	privateKey := config.App.Alipay.PrivateKey
	notifyUrl := config.App.Alipay.NotifyUrl
	if appId == "" || privateKey == "" || notifyUrl == "" {
		utils.ExitError("请先配置支付宝支付配置", -1)
	}
	file, err := utils.ReadFile(privateKey)
	if err != nil {
		return nil
	}
	privateKey = string(file)
	if privateKey == "" {
		utils.ExitError("读取应用私钥失败", -1)
	}
	client, err := alipay.NewClient(appId, privateKey, false)
	if err != nil {
		utils.ExitError(err.Error(), -1)
	}
	//配置公共参数
	client.SetCharset("utf-8").SetSignType(alipay.RSA2).SetNotifyUrl(notifyUrl)
	return client
}
