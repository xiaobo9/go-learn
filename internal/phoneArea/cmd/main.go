package main

import (
	"github.com/xiaobo9/go-learn/config"
	"github.com/xiaobo9/go-learn/internal/phoneArea"
)

//
func main() {
	config.CC.CsvFilePath = "../area.csv"
	phoneArea.DemoMain()
}
