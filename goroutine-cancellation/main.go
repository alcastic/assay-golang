package main

import (
	"fmt"
	"time"
)

func spinner(d time.Duration) {
	fmt.Println("spinner started")
	var se = `-\|/`
	var i = 0
	for {
		select {
		case <-cancelationBroadcast:
			fmt.Println("\nspinner finished")
			return
		default:
			fmt.Printf("\r%c", se[i])
			i = (i + 1) % len(se)
			time.Sleep(d)
		}
	}

}

var cancelationBroadcast = make(chan struct{})

func main() {
	fmt.Println("main finished")

	go spinner(1 * time.Second)

	time.Sleep(10 * time.Second)
	close(cancelationBroadcast)
	time.Sleep(3 * time.Second)
	fmt.Println("main finished")
}
