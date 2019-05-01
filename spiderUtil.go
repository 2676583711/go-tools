package cn_zhou_tools

/*
	爬虫工具,传入相关参数爬取小说
*/

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"net/http"
)

type Task struct {
	Url           string //初始url
	Filename      string //提取数据，写入数据的文件名
	TopicSelector string //章节目录
	TextSelector  string //文本数据
	UrlsSelector  string // //下一页地址css节点
	NextIndex     int    //下一个url的下半截的索引
	NextHalf      string //下一个url上半截
}

type TaskWithSendEmail struct {
	Url           string //初始url
	Filename      string //提取数据，写入数据的文件名
	TopicSelector string //章节目录
	TextSelector  string //文本数据
	UrlsSelector  string // //下一页地址css节点
	NextIndex     int    //下一个url的下半截的索引
	NextHalf      string //下一个url上半截
	SendEmail     bool
}

//var i=3  //下一个url的位置
//var nextHalf="https://www.booktxt.net/2_2219/" //下一个url的上半截
func (t Task) Execute() {

	//在程序结束之前发送邮件
	defer Email{
		Sender:         "",
		Password:       "",
		Receiver:       []string{"liberalzhou@163.com"},
		Subject:        "",
		Text:           []byte("spider & email"),
		AttachmentPath: t.Filename,
		Host:           "",
		HostAndPort:    "",
	}.SendWithAttachment()

	//第一次爬取,以及返回相响应体
	resp := spider(t.Url)
	topic, text, nextUrl := parseDoc(resp, t.TopicSelector,
		t.TextSelector, t.UrlsSelector)

	//写入所提取到的目标数据到文件中
	writeData(t.Filename, topic, text)
	//重复爬取
	for nextUrl != nil {
		//发送请求，返回响应体
		resp := spider(t.NextHalf + nextUrl[t.NextIndex])

		//解析响应体
		topic, text, nextUrl = parseDoc(resp, t.TopicSelector,
			t.TextSelector, t.UrlsSelector)

		//写入文件
		writeData(t.Filename, topic, text)

		//打印下一章的url地址
		fmt.Print(t.NextHalf + nextUrl[t.NextIndex] + "\t")

		//判断是否存在下一页
		//判若不存在下一页，直接panic终端程序
		if (t.NextHalf + nextUrl[t.NextIndex]) == t.NextHalf {
			//panic("爬取完毕")
			fmt.Println("爬取完毕")
			return
		}

	}

}

//根据Url　发送请求，返回响应体
func spider(url string) http.Response {
	client := http.Client{} //创建客户端

	//跳过https的ssl验证
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	//创建请求
	get, err1 := http.NewRequest("GET", url, nil)
	Export{}.PrintMoreError(err1, "创建请求环节出问题了")

	//添加请求参数
	//User-Agent (谷歌浏览器)
	get.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64…) Gecko/20100101 Firefox/64.0")

	//User-Agent (火狐浏览器)
	//get.Header.Set("User-Agent  ", "Mozilla/5.0 (X11; Linux x86_64…) Gecko/20100101 Firefox/64.0")

	//发送请求
	resp, err2 := client.Do(get)
	//错误处理
	Export{}.PrintMoreError(err2, "发送请求环节出问题了")

	return *resp

}

//根据得到的响应体，解析响应体,并返回目标数据
func parseDoc(resp http.Response, topicSelector, textSelector,
	urlsSelector string) (topic, text string, urls [] string) {

	// 解析数据之前对响应体进行编码，防止乱码的出现
	newReader, encodeErr := charset.NewReader(resp.Body, "utf-8")

	//从响应体中解析数据
	doc, err1 := goquery.NewDocumentFromReader(newReader)

	Export{}.PrintMoreError(encodeErr, "编码错误")
	defer resp.Body.Close()
	Export{}.PrintError(err1)
	topic = doc.Find(topicSelector).Text() //标题
	text = doc.Find(textSelector).Text()   //文本内容

	urls = make([] string, 10) //下一页的url
	//queue:=3
	doc.Find(urlsSelector).Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		urls[i] = url
	})
	return topic, text, urls
}

//把爬取到的目标数据写入文件中
func writeData(filename, topic, text string) {
	f := FileUtil{filename}.openAddition()
	f.WriteString(topic)
	f.WriteString("\n\n")
	f.WriteString(text)
	f.WriteString("\n\n\n\n")

	defer f.Close()
	defer fmt.Println(topic, "\t 写入OK")
}
