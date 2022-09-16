package qrcode

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"study-server/app/libs/encry"
)

type qrCode struct {
	qr     *qrcode.QRCode
	border bool // 是否展示边框
	size   int  // 二维码大小
}

// New 二维码生成类: 初始化,传入需要生成的二维码字符串内容
/**
 * @Example:
	q := qrcode.New("xxx")
	q.SetSize(100).SetBackgroundColor(colornames.Red)
	fmt.Println(q.CreateBase64QrCode())
*/
func New(content string) (*qrCode, error) {
	var e error
	q := &qrCode{border: true, size: 300}
	q.qr, e = qrcode.New(content, qrcode.Medium)
	if e != nil {
		return nil, e
	}
	return q, nil
}

// SetContent 设置二维码内容
func (q *qrCode) SetContent(content string) *qrCode {
	q.qr.Content = content
	return q
}

// SetBorder 设置边框
func (q *qrCode) SetBorder(border bool) *qrCode {
	q.border = border
	return q
}

// SetSize 设置二维码大小
func (q *qrCode) SetSize(size int) *qrCode {
	q.size = size
	return q
}

// SetBackgroundColor 设置二维码背景颜色。colornames.Red
func (q *qrCode) SetBackgroundColor(color color.RGBA) *qrCode {
	q.qr.BackgroundColor = color
	return q
}

// SetForegroundColor 设置二维码前景颜色。colornames.Red
func (q *qrCode) SetForegroundColor(color color.RGBA) *qrCode {
	q.qr.BackgroundColor = color
	return q
}

// CreateBase64QrCode 创建base64二维码
func (q *qrCode) CreateBase64QrCode() (string, error) {
	if q.border {
		q.qr.DisableBorder = true //去掉边框
	}
	data, e := q.qr.PNG(q.size)
	if e != nil {
		return "", e
	}
	return "data:image/png;base64," + encry.Base64Encode(data), nil
}
