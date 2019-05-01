package cn_zhou_tools

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/**
发送邮件
*/

/**
//发送人
	Sender string
	//发送人密码
	Password string
	//收件人,多个（群发邮件）
	Receiver []string
	//邮件主题
	Subject string
	//邮件内容
	Text [] byte
	//附件路径
	AttachmentPath string
	//服务器域名
	Host string
	//服务器家端口
	HostAndPort string
*/
//首字母大写，才可以访问
//邮件使用入口
type Email struct {
	//发送人
	Sender string
	//发送人密码
	Password string
	//收件人,多个（群发邮件）
	Receiver []string
	//邮件主题
	Subject string
	//邮件内容
	Text [] byte
	//附件路径
	AttachmentPath string
	//服务器域名
	Host string
	//服务器家端口
	HostAndPort string
}

//发送带附件的邮件
func (e Email) SendWithAttachment() {
	email := email.NewEmail()

	//如果没有给定发件人信息，就使用默认的信息
	if len(e.Sender) == 0 || e.Sender == " " {
		e.Sender = "liberalzhou@163.com"
	}

	if e.Password == " " || len(e.Password) == 0 {
		e.Password = "zhou123456"
	}

	if e.Host == " " || len(e.Host) == 0 {
		e.Host = "smtp.163.com"
	}
	if e.HostAndPort == " " || len(e.HostAndPort) == 0 {
		e.HostAndPort = "smtp.163.com:25"
	}

	//否则使用给定的发件人信息
	//发件人
	email.From = e.Sender
	//收件人
	email.To = e.Receiver
	//邮件的主题
	email.Subject = e.Subject
	//邮件的内容　
	email.Text = e.Text
	//附件路径
	email.AttachFile(e.AttachmentPath)
	//发送邮件
	err := email.Send(e.HostAndPort, smtp.PlainAuth("", e.Sender, e.Password, e.Host))
	if err == nil {
		fmt.Println("email send ok")
	}
}
