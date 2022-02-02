package main

import "math"

func power(n int, ch chan float64) {
	ch <- math.Sqrt(float64(n))
}

func main() {
	ch := make(chan float64)
	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		println(<-ch)
	}
}
