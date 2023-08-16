//go:build !windows
// +build !windows

package webview

/*
#cgo linux pkg-config: webkit2gtk-4.0
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa -framework Webkit

#include "webview.h"
*/
import "C"

import (
	"github.com/getlantern/systray"
	icon "github.com/linexjlin/systray-icons/openai-logomark"
)

func configureWebview(title string, width, height int) {
	C.configureAppWindow(C.CString(title), C.int(width), C.int(height))
}

func openUrl(url string) {
	configureWebview("ChatDesk", 1024, 768)
	C.showAppWindow(C.CString(url))
}

var URL = ""
var W = 1024
var H = 768

func ShowWebview2(url string) {
	URL = url
	systray.Register(onReady, nil)
	configureWebview("ChatDesk", 1024, 768)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(UText(""))
	mWebUI := systray.AddMenuItem("Open ChatDesk", UText("Open ChatDesk"))
	mQuit := systray.AddMenuItem(UText("Quit"), UText("Quit the whole app"))
	go func() {
		for {
			select {
			case <-mWebUI.ClickedCh:
				openUrl(URL)
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
	mWebUI.ClickedCh <- struct{}{}
}
