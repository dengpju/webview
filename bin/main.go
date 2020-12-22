package main

import "github.com/webview/webview"

// 基于 CygWin 执行 GoLang 编译报 GCC 错误解决 https://lzxz1234.cn/archives/278
func main()  {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://www.youbbs.org/t/2184")
	w.Run()
}
