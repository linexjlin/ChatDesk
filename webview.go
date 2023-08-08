package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jchv/go-webview-selector"
)

/*
const initJS = `window.onload=function(){
    getWebInfo(document.title, window.location.href)
}`

func configureWebview() {
	var w = webview.New(false)
	url := "https://chat.openai.com"
	///debug := true
	//w := webview.New(false)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)

	go func() {
		time.Sleep(time.Second * 10)

		fmt.Println("goto google")
		w.Navigate("https://chat.openai.com")
		w.SetSize(800, 600, webview.HintNone)

		w.Run()

	}()
	_ = w.Bind("getWebInfo", getWebInfo)
	w.Init(initJS)
	w.Run()
	fmt.Println("exit")
}

func getWebInfo(title, uri string) {
	fmt.Println("getWebInfo", title, uri)
}
*/

func showWebview(url string) {
	w := webview.New(false)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate(url)
	w.Run()
	fmt.Println("exit")
}

