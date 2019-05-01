package cn_zhou_tools

import (
 	"os"
)

/*
	文件操作工具
  传入文件名，返回文件对象
 */

type FileUtil struct {
	Filename string
	//file
}



//打开一个文件若该文件不存在则创建一个新文件并打开
func (f FileUtil)openNew()(file *os.File){
	file,err:= os.Open(f.Filename)
	if err !=nil {
		file,err=os.Create(f.Filename)
		Export{}.PrintError(err)
 	}
	return file
 }

//以写跟追加的方式打开文件
func (fu FileUtil)openAddition() *os.File{
	f, err := os.OpenFile( fu.Filename, os.O_WRONLY|os.O_APPEND, 0666)
	if(err!=nil){
		f,err=os.Create(fu.Filename)
		Export{}.PrintError(err)
	}
	return f
}