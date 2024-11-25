package main

import (
	"C"
	"fmt"

	"github.com/webview/webview_go"
)

var w webview.WebView

//export openWebView
func openWebView() {
	fmt.Println("go: creating webview...")
	w = webview.New(false)
	defer w.Destroy()
	fmt.Println("go: setting title...")
	w.SetTitle("WebView Example")
	w.Navigate("http://localhost:8000/")
	w.SetSize(800, 600, webview.HintNone)
	fmt.Println("go: running webview...")
	w.Run()
	fmt.Println("go: webview closed!")
}

func main() {}
