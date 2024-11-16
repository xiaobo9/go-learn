// 程序主入口
package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/lxn/walk"
	wd "github.com/lxn/walk/declarative"
	. "github.com/xiaobo9/go-learn/config"
	"github.com/xiaobo9/go-learn/internal/gui"
	"github.com/xiaobo9/go-learn/server"
)

var (
	port     = flag.String("port", "8080", "server port")
	data     = flag.String("dir", "./", "date file dir")
	iconName = "favicon.ico"
	srv      server.Server
)

func main() {
	log.Println("I'm running!")

	// TODO 可配置
	CC.Host = "127.0.0.1"

	flag.Parse()

	srv = server.Server{Host: CC.Host, Port: *port, Data: *data}
	go srv.Serve()

	var mw, err = walk.NewMainWindow()
	if err != nil {
		log.Fatalln(err)
		srv.Shutdown()
		return
	}
	var width, height = 400, 260
	wd.MainWindow{
		AssignTo: &mw,
		Title:    "测试",
		MinSize:  wd.Size{Width: width, Height: height},
		Layout:   wd.VBox{},
		Children: []wd.Widget{
			// wd.HSplitter
			wd.PushButton{
				Text:      "测试",
				OnClicked: func() { walk.MsgBox(mw, "测试", "测试", walk.MsgBoxIconInformation) },
			},
		},
	}.Create()

	gui.Win2Center(mw, width, height)

	notifyIcon, _ := walk.NewNotifyIcon(mw)
	defer notifyIcon.Dispose()

	{
		setSysTray(notifyIcon)
		mw.Run()
	}

}

// 设置系统托盘
func setSysTray(notifyIcon *walk.NotifyIcon) {
	icon, err := walk.Resources.Image(iconName)
	if err == nil {
		notifyIcon.SetIcon(icon)
	}
	notifyIcon.SetVisible(true)

	notifyIcon.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			openBrowser()
		}
	})

	notifyIcon.SetToolTip("我是悬浮提示")

	{
		gui.AddMenu2SysTray(notifyIcon, "打开浏览器", func() { openBrowser() })
		gui.AddMenu2SysTray(notifyIcon, "退出", func() {
			srv.Shutdown()
			walk.App().Exit(0)
		})
	}
	// notifyIcon.ShowInfo("title", "info")
}


func openBrowser() {
	log.Println("打开浏览器")
	// command := exec.Command(`cmd`, `/C`, `start`, `http://127.0.0.1:8080`)
	command := exec.Command("cmd", "/C", "start", "http://127.0.0.1:8080")
	if err := command.Start(); err != nil {
		log.Println(err)
	}
}
