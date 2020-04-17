package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
	. "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	hWnd        win.HWND
	minimizeBox *walk.CheckBox
	maximizeBox *walk.CheckBox
	closeBox    *walk.CheckBox
	sizeBox     *walk.CheckBox
}

func (mw *MyWindow) addStyle(style int32) {
	currStyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, currStyle|style)
}

func (mw *MyWindow) removeStyle(style int32) {
	currStyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, currStyle&style)
}


func (mw *MyWindow) SetMinimizeBox() {
	if mw.minimizeBox.Checked() {
		mw.addStyle(win.WS_MINIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MINIMIZEBOX)
}

func (mw *MyWindow) SetMaximizeBox() {
	if mw.maximizeBox.Checked() {
		mw.addStyle(win.WS_MAXIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MAXIMIZEBOX)
}

func (mw *MyWindow) SetCloseBox() {
	if mw.closeBox.Checked() {
		win.GetSystemMenu(mw.hWnd, true)
		return
	}
	hMenu := win.GetSystemMenu(mw.hWnd, false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
}

func (mw *MyWindow) SetSizePersistent() {
	if mw.sizeBox.Checked() {
		mw.addStyle(win.WS_SIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_SIZEBOX)
}


func main() {

	mw := new(MyWindow)

	MainWindow {
		AssignTo: &mw.MainWindow,
		Title:  "Test",
		Size: Size{Width: 450, Height: 300},
		Layout:  VBox{},
		Children: []Widget{
			CheckBox{
				AssignTo:            &mw.minimizeBox,
				Text:                "显示最小化按钮",
				Checked:             false,
				OnCheckStateChanged: mw.SetMinimizeBox,
			},
			CheckBox{
				AssignTo:            &mw.maximizeBox,
				Text:                "显示最大化按钮",
				Checked:             false,
				OnCheckStateChanged: mw.SetMaximizeBox,
			},
			CheckBox{
				AssignTo:            &mw.closeBox,
				Text:                "显示关闭按钮",
				Checked:             false,
				OnCheckStateChanged: mw.SetCloseBox,
			},
			CheckBox{
				AssignTo:            &mw.sizeBox,
				Text:                "允许修改大小",
				Checked:             false,
				OnCheckStateChanged: mw.SetSizePersistent,
			},
		},
	}.Create()

	mw.hWnd = mw.Handle()
	
	// 设置 ^win.WS_MINIMIZEBOX 禁用最小化按钮
	// 设置 ^win.WS_MAXIMIZEBOX 禁用最大化按钮
	// 设置 ^win.WS_SIZEBOX 禁用窗口大小改变
	win.SetWindowLong(
		mw.Handle(),
		win.GWL_STYLE, 
		win.GetWindowLong(mw.Handle(), win.GWL_STYLE) &
			^win.WS_MINIMIZEBOX &
			^win.WS_MAXIMIZEBOX &
			^win.WS_SIZEBOX,
	)
	// 禁用关闭按钮
	hMenu := win.GetSystemMenu(mw.Handle(), false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)

	mw.Run()
}
