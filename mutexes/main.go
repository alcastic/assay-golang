package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock() // alternative expresion form: (*mutex).Lock()
	*b += n
	mutex.Unlock() // alternative expresion form: (*mutex).Unlock()
	wg.Done()
}

func withdraw(b *int, n int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock() // alternative expresion form: (*mutex).Lock()
	*b -= n
	mutex.Unlock() // alternative expresion form: (*mutex).Unlock()
	wg.Done()
}

func main() {
	var mutex sync.Mutex  // alternative initialization form: mutex := sync.Mutex{}
	var wg sync.WaitGroup // alternative initialization form: wg := sync.WaitGroup{}
	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &mutex, &wg)
		go withdraw(&balance, i, &mutex, &wg)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
