package gui

import (
	"log"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

var (
	mainWindow *walk.MainWindow
)
var isLogin bool

func Gui() {
	// gui.Window()

	OpenGui(mainWindow)
}

// 禁用窗口最大化 禁用窗口大小修改
func freezeWindow(mw *walk.MainWindow) {
	hwnd := mw.Handle()
	currStyle := win.GetWindowLong(hwnd, win.GWL_STYLE)
	win.SetWindowLong(hwnd, win.GWL_STYLE, currStyle&^win.WS_MAXIMIZEBOX&^win.WS_SIZEBOX)
}

// 窗口居中
func Win2Center(mw *walk.MainWindow, width int, height int) {
	xScreen := win.GetSystemMetrics(win.SM_CXSCREEN)
	yScreen := win.GetSystemMetrics(win.SM_CYSCREEN)
	win.SetWindowPos(
		mw.Handle(),
		0,
		(xScreen-int32(width))/2,
		(yScreen-int32(height))/2,
		int32(width),
		int32(height),
		win.SWP_FRAMECHANGED,
	)
	win.ShowWindow(mw.Handle(), win.SW_SHOW)
}

// 给托盘图标添加 menu
func AddMenu2SysTray(ni *walk.NotifyIcon, text string, handler walk.EventHandler) {
	action := walk.NewAction()
	if err := action.SetText(text); err != nil {
		log.Println(err)
	}
	action.Triggered().Attach(handler)

	if err := ni.ContextMenu().Actions().Add(action); err != nil {
		log.Println(err)
	}
}
