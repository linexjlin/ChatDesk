package main

import (
	"embed"
	"net/http"
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
	go ServeWeb()
	showWebview("http://" + serverAddr)
}
