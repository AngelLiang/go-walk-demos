package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	. "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	hWnd    	win.HWND
}


func main() {

	mw := new(MyWindow)

	MainWindow {
		AssignTo: &mw.MainWindow,
		Title:  "Test",
		Size: Size{Width: 450, Height: 300},
		Layout:  VBox{},
		Children: []Widget{
		},
	}.Create()


	// 启动时触发
	mw.Starting().Attach(func() {
		fmt.Println("Starting")
	})

	// 激活时触发
	mw.Activating().Attach(func() {
		fmt.Println("Activating")
	})

	mw.VisibleChanged().Attach(func() {
		fmt.Println("VisibleChanged")
	})

	// 改变焦点的时候会触发
	mw.FocusedChanged().Attach(func() {
		fmt.Println("FocusedChanged")
	})

	// 移动的时候会触发
	mw.BoundsChanged().Attach(func() {
		fmt.Println("BoundsChanged")
	})

	// 关闭时候回触发
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		fmt.Printf("Closing canceled=%v reason=%v\n", canceled, reason)
	})

	mw.Deactivating().Attach(func() {
		fmt.Println("Deactivating")
	})

	// 销毁的时候回触发
	mw.Disposing().Attach(func() {
		fmt.Println("Disposing")
	})

	mw.hWnd = mw.Handle()

	mw.Run()
}
