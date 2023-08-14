//go:build !windows
// +build !windows

package webview

import "github.com/webview/webview"

func ShowWebview2(url string) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("ChatDesk Application")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
