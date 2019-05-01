package cn_zhou_tools

/*
错误输出工具
*/
type Export struct {
	print
}

//打印错误信息
type print interface {
	PrintError(err error)
	PrintMoreError(err error, message string)
}

//判断错误信息
func (p Export) PrintError(err error) {
	if err != nil {
		panic(err)
	}
}

//判断错误信息
//err error,message string
func (p Export) PrintMoreError(err error, message string) {
	if err != nil {
		println(message)
		panic(err)
	}
}
