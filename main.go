package main

import (
	"github.com/getlantern/systray"
	icon "github.com/linexjlin/systray-icons/enter-the-keyboard"
)

func main() {
	//systray.Register(onReady, nil)
	systray.Run(onReady, nil)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Webview example")
	mShowLantern := systray.AddMenuItem("Show Lantern", "")
	mShowWikipedia := systray.AddMenuItem("Show Wikipedia", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShowLantern.ClickedCh:
				showWebview("https://www.getlantern.org")
			case <-mShowWikipedia.ClickedCh:
				showWebview("https://www.wikipedia.org")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}
