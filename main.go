package main

import (
	"embed"
	"net/http"

	"github.com/linexjlin/ChatDesk/webview"
)

//go:embed all:ChatGPT-Next-Web/out
var webStaticFS embed.FS
var serverAddr = "127.0.0.1:38612"

func ServeWeb() {
	fileServer := http.FileServer(http.FS(webStaticFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/ChatGPT-Next-Web/out" + r.URL.Path
		fileServer.ServeHTTP(w, r)
	})
	http.ListenAndServe(serverAddr, nil)
}

func main() {
	go ServeWeb()
	webview.ShowWebview("http://" + serverAddr)
}
