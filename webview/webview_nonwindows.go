//go:build !windows
// +build !windows

package webview

import (
	"log"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
	icon "github.com/linexjlin/systray-icons/openai-logomark"
)

var URL string

func openUrl(url string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", "-a", "Safari", url)
	} else {
		cmd = exec.Command("xdg-open", url)
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func ShowWebview(url string) {
	URL = url
	systray.Run(onReady, nil)
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
