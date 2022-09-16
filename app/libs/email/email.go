package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"mime"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

/* 邮箱类 */
type Email struct {
	name string
	user string
	pass string
	host string
	port int
	em   *gomail.Message
}

// 邮箱类，初始化方法
/**
 *@Example:
	e := email.New()
	e.SetTitle("测试")
	e.SetToEmail([]string{"buxsren@qq.com"})
	e.SetBody("<h1>哈哈哈</h1>")
	e.AddAttach("test.txt","/www/wwwroot/public/test.txt")
	fmt.Println(e.SendMail())
*/
func New() *Email {
	e := Email{
		name: config.App.Email.Name,
		user: config.App.Email.User,
		pass: config.App.Email.Pass,
		host: config.App.Email.Host,
		port: config.App.Email.Port,
	}
	if e.name == "" || e.user == "" || e.pass == "" || e.host == "" || e.port == 0 {
		utils.ExitError("请先配置邮箱设置", -1)
	}
	e.em = gomail.NewMessage()
	e.em.SetHeader("From", e.em.FormatAddress(e.user, e.name)) // 设置发件人名称
	return &e
}

// 发送邮件
func (e *Email) SendMail() error {
	d := gomail.NewDialer(e.host, e.port, e.user, e.pass)
	return d.DialAndSend(e.em)
}

// 设置/更改发件人名称
func (e *Email) SetFormName(name string) *Email {
	e.em.SetHeader("From", e.em.FormatAddress(e.user, name))
	return e
}

// 设置发件人
func (e *Email) SetToEmail(to []string) *Email {
	e.em.SetHeader("To", to...)
	return e
}

// 设置邮件标题
func (e *Email) SetTitle(title string) *Email {
	e.em.SetHeader("Subject", title)
	return e
}

// 设置邮件正文,支持html格式
func (e *Email) SetBody(body string) *Email {
	e.em.SetBody("text/html", body) //设置邮件正文
	return e
}

// 添加附件.附件名称(xxx.txt),附件路径(/www/wwwwroot/public/xxx.txt)
func (e *Email) AddAttach(name, path string) *Email {
	e.em.Attach(path, gomail.Rename(name),
		gomail.SetHeader(map[string][]string{ // 附件中文乱码解决方案
			"Content-Disposition": {fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", name))},
		}),
	)
	return e
}
