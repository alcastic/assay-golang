package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done() // alternative expresion: (*wg).Done()
}

func main() {
	wg := sync.WaitGroup{} // alternative initialization: var wg *sycn.WaitGroup
	nSubTasks := 1
	wg.Add(nSubTasks)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()
}
