package main

import (
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"net/url"
	"sync/atomic"
	"time"
)

func main() {
	/**
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

	 */
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Data model: number of ticks
	ticks := uint32(0)
	// Channel to connect UI events with the background ticking goroutine
	togglec := make(chan bool)
	// Bind Go functions to JS
	ui.Bind("toggle", func() { togglec <- true })
	ui.Bind("reset", func() {
		atomic.StoreUint32(&ticks, 0)
		ui.Eval(`document.querySelector('.timer').innerText = '0'`)
	})

	// Load HTML after Go functions are bound to JS
	ui.Load("data:text/html," + url.PathEscape(`
	<html>
		<body>
			<!-- toggle() and reset() are Go functions wrapped into JS -->
			<div class="timer" onclick="toggle()"></div>
			<button onclick="reset()">Reset</button>
			<iframe src="https://hsbcapi.51jiaoyujia.com/admin"></iframe>
		</body>
	</html>
	`))

	// Start ticker goroutine
	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t.C: // Every 100ms increate number of ticks and update UI
				ui.Eval(fmt.Sprintf(`document.querySelector('.timer').innerText = 0.1*%d`,
					atomic.AddUint32(&ticks, 1)))
			case <-togglec: // If paused - wait for another toggle event to unpause
				<-togglec
			}
		}
	}()
	<-ui.Done()
}