package main

import "fmt"

func genNaturals() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func print(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	print(square(genNaturals()))
}
