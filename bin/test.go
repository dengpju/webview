package main

import (
	"fmt"
	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("https://hsbcapi.51jiaoyujia.com/admin", "", 1280, 1024)
	defer ui.Close()

	// Bind Go function to be available in JS. Go function may be long-running and
	// blocking - in JS it's represented with a Promise.
	_ = ui.Bind("add", func(a, b int) int { return a + b })

	// Call JS function from Go. Functions may be asynchronous, i.e. return promises
	n := ui.Eval(`Math.random()`).Float()
	fmt.Println(n)

	// Call JS that calls Go and so on and so on...
	m := ui.Eval(`add(2, 3)`).Int()
	fmt.Println(m)

	// Wait for the browser window to be closed
	<-ui.Done()
}