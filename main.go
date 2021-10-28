package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"syscall"
)

var (
	SCREEN_WIDTH  int32
	SCREEN_HEIGHT int32
)

func main() {
	//GenerateRice()
	createMainWindow()
}

func createMainWindow() {
	LoadDll()
	screenSize()
	w, err := window.New(sciter.DefaultWindowCreateFlag,
		&sciter.Rect{Left: SCREEN_WIDTH/2-300, Top: SCREEN_HEIGHT/2-200, Right: SCREEN_WIDTH/2 + 300, Bottom: SCREEN_HEIGHT/2 - 124})
	if err != nil {
		log.Fatal(err)
	}

	HandleDataLoad(w.Sciter)
	w.LoadFile("rice://res/index.html")
/*	fullpath, err := filepath.Abs("res/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)*/
	Handler(w)
	//设置标题
	//w.SetTitle("表单")
	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}

func screenSize() {
	user32 := syscall.NewLazyDLL("User32.dll")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")
	width, _, _ := getSystemMetrics.Call(0)
	height, _, _ := getSystemMetrics.Call(1)
	SCREEN_WIDTH = int32(width)
	SCREEN_HEIGHT = int32(height)
}
