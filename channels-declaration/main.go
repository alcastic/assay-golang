package main

import "fmt"

func main() {
	// bidirectional unbuffered channel that works with values of type float64
	var ch1 chan float64
	// receive-only channel that works with values of type rune initialiced with make() built-in function
	var ch2 <-chan rune = make(<-chan rune)
	// send-only  channel that works with values of type rune initialiced with make() built-in function
	var ch3 chan<- rune = make(chan<- rune)
	// bidirectional buffered channel with a capacity of 10 ints.
	var ch4 chan int = make(chan int, 10)
	fmt.Printf("Type of ch1: %T\n", ch1)
	fmt.Printf("Type of ch2: %T\n", ch2)
	fmt.Printf("Type of ch3: %T\n", ch3)
	fmt.Printf("Type of ch4: %T\n", ch4)
}
