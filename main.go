package main

import (
	"embed"
	"net/http"

	"github.com/getlantern/systray"
	icon "github.com/linexjlin/systray-icons/openai-logomark"
)

//go:embed web
var webStaticFS embed.FS
var serverAddr = "127.0.0.1:28612"

func ServeWeb() {
	fileServer := http.FileServer(http.FS(webStaticFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/web" + r.URL.Path
		fileServer.ServeHTTP(w, r)
	})
	http.ListenAndServe(serverAddr, nil)
}

func main() {
	//systray.Register(onReady, nil)
	go ServeWeb()
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
				showWebview("http://" + serverAddr)
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
	mWebUI.ClickedCh <- struct{}{}
}
