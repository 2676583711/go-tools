package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"gopkg.in/gomail.v2"
	"io"
	"net/smtp"
)

/*
	SMTP服务器就是邮件代收发服务器，由邮件服务商提供，常见的SMTP服务器端口号：
	QQ邮箱：SMTP服务器地址：smtp.qq.com（端口：587）
	雅虎邮箱: SMTP服务器地址：smtp.yahoo.com（端口：587）
	163邮箱：SMTP服务器地址：smtp.163.com（端口：25）
	126邮箱: SMTP服务器地址：smtp.126.com（端口：25）
	新浪邮箱: SMTP服务器地址：smtp.sina.com（端口：25）
*/
func main() {
	//sendEmail()
	emailWithAttach()
}

/*
使用go自带的smtp模块
*/
func sendEmailUseSmtp() {

}

/**
使用第三方的包
get -v github.com/jordan-wright/email
*/
func emailWithAttach() {
	//调用方法返回email类型的结构体
	//该结构体装载email发送过程中的信息
	newEmail := email.NewEmail()
	//email的发件人
	newEmail.From = "liberalzhou@163.com"
	//收件人
	newEmail.To = []string{"liberalzhou@163.com"}
	//附件文件，传入文件地址即可
	newEmail.AttachFile("/home/zhou/美漫世界的武者.txt")
	newEmail.Text = []byte("小说附件邮件")


	//
	//newEmail.Attach(strings.NewReader("hello,file"),
	//	"/home/zhou/美漫世界的武者.txt", "text/plain")

	newEmail.Send("smtp.163.com:25",
		smtp.PlainAuth("",
			"liberalzhou@163.com", "zhou123456", "smtp.163.com"))

}

/**
使用第三方的包
 gopkg.in/gomail.v2
*/
func sendEmail() {
	//创建消息体
	message := gomail.NewMessage()
	//设置发件人,From,f必须大写
	message.SetHeader("From", "liberalzhou@163.com")
	//设置收件人,To,t必须大写
	message.SetHeader("To", "liberalzhou@163.com")
	//设置消息头
	message.SetHeader("go email")
	//设置消息体，也就是正文
	message.SetBody("text/html", "hello,go,email")

	//新键拨号，并调用DialAndSend(message)方法拨号并发送消息体
	gomail.NewDialer("smtp.163.com", 25, "liberalzhou@163.com",
		"zhou123456").DialAndSend(message)
}

func testSendMail() {
	//创建消息体
	message := gomail.NewMessage()
	//设置发件人
	message.SetHeader("From", "liberalzhou@163.com")
	//设置收件人
	message.SetHeader("To", "liberalzhou@163.com")
	//设置消息头
	message.SetHeader("go email")
	//设置消息体，也就是正文
	message.SetBody("text/html", "hello,go,email")

	//新键拨号，并调用DialAndSend(message)方法拨号并发送消息体
	gomail.NewDialer("smtp.163.com", 25, "liberalzhou@163.com",
		"zhou123456").DialAndSend(message)

	sendFunc := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {

		return nil
	})

	gomail.Send(sendFunc, message)

}

func testGoMap() {
	/*
		go 新建map 3种方法:
			map类型变量的大小是8，实际上这是一个指针，也就是说map类型就是一个指针。因此
			m1实际上定义了一个map指针，这个指针指向NULL。
			既然m1是一个空指针，并没有一个真实的map存在，所以也就不能对m1进行内存访问操作，
			比如m1[key] = value，但奇怪的是可以读，包括ss := m1[key]和遍历(我不知道为什么这样设计)
			m2和m3定义了一个map指针，指向了一个已经生成了一个map对象。
			从打印出m1/m2/m3的值，我们可以看出m1的值为0，m2/m3处的值不为零。
			也说明go里面函数传递map的时候是传的指针，不是map对象的拷贝。

			注意m2和m3两种方式是等价的。
	*/
	//1.仅仅声明map，没有初始化（赋值），此种map被称为nil map,只能取值，不能写值
	//所以必须声明同时赋值，否则空map没有意义
	var m1 map[string]string

	// or m2 := map[string]string{}
	var m2 map[string]string = map[string]string{}

	//2.局部map,可以仅仅声明，用时赋值
	//创建go-map
	m := make(map[string]string)
	m["m1"] = "abc"
	fmt.Println(m["m1"])

	// or m3 := make(map[string]string)
	var m3 map[string]string = make(map[string]string, 10)

	fmt.Println(m1, m2, m3)
}
