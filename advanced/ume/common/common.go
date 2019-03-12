package common

import (
	"fmt"
	"log"
	"os"
)

//记录错误日志
const (
	LOG_DIRECTORY = "/usr/local/var/log/ume_error.log"
)

var logFile *os.File
var logger *log.Logger

func LogError() {
	//打开日志文件
	logFile, err := os.OpenFile(LOG_DIRECTORY, os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("日志文件创建失败！！！")
		os.Exit(-1)
	}
	defer logFile.Close()

	//使用标准库log生成日志对象
	logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
}

//检查是否有错误
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}
}
