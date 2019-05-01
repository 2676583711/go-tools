package main

import (
	"2018/cn.zhou.tools"
)

/*
测试或运行，工具包程序
*/

/**
使用爬虫工具，爬去小说
*/

func main() {

	//time_4_29()
	//time_4_29_2()
	time_5_1()

}

func time_5_1() {
	cn_zhou_tools.Task{"https://www.biquge.biz/28_28122/12307619.html",
		"/home/zhou/我能看见熟练度.txt",
		"div.bookname h1", //章节
		"div#content",     //小说内容
		"div.bottem1 a",   //下一页地址css节点
		2,                 //下一页地址索引
		"https://www.biquge.biz",  }.Execute()
}

/*
year:2019
*/
//https 网页  (顶点小说网)
func time_4_29() {
	cn_zhou_tools.Task{"http://www.dingdiann.com/ddk177015/9255511.html",
		"/home/zhou/美漫世界的武者.txt",
		"div.bookname h1", //章节
		"div#content",     //小说内容
		"div.bottem2 a",   //下一页地址css节点
		3,                 //下一页地址索引
		"http://www.dingdiann.com"}.Execute()
}

/*
year:2019
*/
//http网页 (笔趣阁小说网)
//http://www.biquge.li/220700/
func time_4_29_2() {
	cn_zhou_tools.Task{"http://www.biquge.li/220700/1131117.html",
		"/home/zhou/美漫世界的武者.txt",
		"div.bookname h1", //章节
		"div#content",     //小说内容
		"div.bottem1 a",   //下一页地址数组
		2,                 //下一页地址索引
		"http://www.biquge.li/220700/"}.Execute()
}
