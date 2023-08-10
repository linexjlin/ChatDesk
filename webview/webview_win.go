//go:build windows
// +build windows

package webview

import (
	"log"

	"github.com/jchv/go-webview2"
	"github.com/lxn/win"
)

func init() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
}

func ShowWebview(url string) {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "ChatDesk Application",
			Width:  1024,
			Height: 768,
			IconId: 2, // icon resource id
			Center: true,
		},
	})

	//w := webview2.New(true)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.Navigate(url)
	w.Run()
}
