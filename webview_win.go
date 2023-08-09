//go:build windows
// +build windows

package main

import (
	"log"

	"github.com/jchv/go-webview2"
	"github.com/lxn/win"
)

func init() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
}

func showWebview(url string) {
	w := webview2.New(true)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetSize(1024, 768, webview2.HintNone)
	w.SetTitle(UText("ChatDesk"))
	w.Navigate(url)
	w.Run()
}
