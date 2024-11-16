package gui

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	SIZE_WIDTH  = 400
	SIZE_HEIGHT = 100
)

// 多窗口 demo 入口函数
func OpenGui(mw *walk.MainWindow) {
	if !isLogin {
		loginWindow(mw)
		return
	}
	mainWD()
}

// 登录窗口
func loginWindow(mw *walk.MainWindow) {
	var username, password *walk.LineEdit
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatalln("LoginUI create error")
		return
	}
	MainWindow{
		AssignTo: &mw,
		Title:    "登录",
		Size:     Size{Width: SIZE_WIDTH, Height: SIZE_HEIGHT},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "                            用户名："},
					LineEdit{AssignTo: &username, MaxSize: Size{Width: 150, Height: 40}},
				},
			},
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "                            密    码："},
					LineEdit{AssignTo: &password, MaxSize: Size{Width: 150, Height: 40}, PasswordMode: true},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text:    "登录",
						MaxSize: Size{Width: 50, Height: 40},
						OnClicked: func() {
							if !loginCheck(username.Text(), password.Text()) {
								walk.MsgBox(mw, "错误", "登录信息有误", walk.MsgBoxIconError)
								return
							}
							saveLoginStatus()
							mw.Close()
							mainWD()
						},
					},
				},
			},
		},
	}.Create()
	freezeWindow(mw)
	mw.SetIcon(getIcon())
	Win2Center(mw, SIZE_WIDTH, SIZE_HEIGHT)
	setSysTray(mw)
}

// 主窗口
func mainWD() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		walk.MsgBox(mw, "错误", "程序打开失败", walk.MsgBoxIconError)
		walk.App().Exit(0)
		return
	}
	MainWindow{
		AssignTo: &mw,
		Title:    "主窗口",
		Size:     Size{Width: SIZE_WIDTH, Height: SIZE_HEIGHT},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "                        主窗口信息："},
					LineEdit{Text: "这是主窗口", MaxSize: Size{Width: 150, Height: 40}, ReadOnly: true},
				},
			},
		},
	}.Create()
	mw.SetIcon(getIcon())
	Win2Center(mw, SIZE_WIDTH, SIZE_HEIGHT)
	setSysTray(mw)
}

// 设置系统托盘
func setSysTray(mw *walk.MainWindow) {
	notifyIcon, _ := walk.NewNotifyIcon(mw)
	defer notifyIcon.Dispose()
	icon := getIcon()
	if icon != nil {
		notifyIcon.SetIcon(icon)
	}
	notifyIcon.SetVisible(true)

	AddMenu2SysTray(notifyIcon, "测试按钮", func() {
		log.Println("测试按钮")
	})
	AddMenu2SysTray(notifyIcon, "退出", func() { walk.App().Exit(0) })
	mw.Run()
}

// 获取icon
func getIcon() walk.Image {
	icon, err := walk.Resources.Image("favicon.ico")
	if err != nil {
		log.Println("ERROR:获取icon失败", err)
	}
	return icon
}

func loginCheck(name, pwd string) bool {
	return pwd == "123"
}

func saveLoginStatus() {
	isLogin = true
}
