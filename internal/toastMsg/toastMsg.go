package toastMsg

import (
	"log"

	"github.com/go-toast/toast"
)

type toastMsg struct {
	user    string
	content string
}

func Notification() {
	notification := toast.Notification{
		AppID:   "xiaobo9",
		Title:   "Toast Title",
		Message: "这是消息内容，等等。。。",
		Actions: []toast.Action{
			{
				Type:      "protocol",
				Label:     "查看消息",
				Arguments: "http://github.com/xiaobo9",
			},
			{
				Type:      "protocol",
				Label:     "第二个按钮",
				Arguments: "http://github.com/xiaobo9",
			},
		},
		// Audio: toast.IM,
		// Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
