package main

import (
	"log"
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	. "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	hWnd    	win.HWND
}

func CreateMyWindow() *MyWindow {
	mw, err := walk.NewMainWindow()
    if err != nil {
		log.Fatal(err)
	}

	mmw := &MyWindow{mw, mw.Handle()}

	// InitWrapperWindow initializes a window that wraps (embeds) another window.
	//
	// Calling this method is necessary, if you want to be able to override the
	// WndProc method of the embedded window. The embedded window should only be
	// used as inseparable part of the wrapper window to avoid undefined behavior.
	if err := walk.InitWrapperWindow(mmw); err != nil {
		log.Fatal(err)
		return nil
	}

	return mmw
}

func (mw *MyWindow) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	fmt.Println(msg)
	switch msg {
		case win.WM_ACTIVATEAPP:
			fmt.Println("WM_ACTIVATEAPP")
		case win.WM_COMMAND:
			fmt.Println("WM_COMMAND")
		case win.WM_CLOSE:
			fmt.Printf("WM_CLOSE")
		case win.WM_SYSCOMMAND:
			fmt.Printf("WM_SYSCOMMAND")
	}
	return mw.MainWindow.WndProc(hwnd, msg, wParam, lParam)
}


func main() {

	mw := CreateMyWindow()

	MainWindow {
		AssignTo: &mw.MainWindow,
		Title:  "Test",
		Size: Size{Width: 450, Height: 300},
		Layout:  VBox{},
		Children: []Widget{
		},
	}.Create()

	// mw.hWnd = mw.Handle()

	mw.Run()
}
