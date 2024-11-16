package config

import (
	"log"
	"os"
)

// 配置信息
type Config struct {
	Host         string
	ServerPort   int
	StringConfig map[string]string
	IntConfig    map[string]int
	CsvFilePath  string
	Log2File     bool
}

var CC Config

// first init
func init() {
	// 日志初步初始化
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	{
		// server init
		CC.Host = "0.0.0.0"
		CC.ServerPort = 8080
		CC.CsvFilePath = "area.csv"
	}
	CC.Log2File = false
}

// 可调用的初始化
func InitConfig() {
	if CC.Log2File {
		var file = "log.txt"
		logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err == nil {
			// 将文件设置为loger作为输出
			log.SetOutput(logFile)
			return
		}
	}

}
