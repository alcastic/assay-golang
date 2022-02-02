package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) plusWithMethodReceiver(scalar int) {
	p.X += scalar
	p.Y += scalar
}

func (p *Point) plusWithMethodPointerReceiver(scalar int) {
	p.X += scalar
	p.Y += scalar
}

func main() {
	p := Point{X: 1, Y: 1}
	fmt.Printf("Point Original: (%d, %d)\n", p.X, p.Y)

	p.plusWithMethodReceiver(1)
	fmt.Printf("Point after Method Receiver: (%d, %d)\n", p.X, p.Y)

	p.plusWithMethodPointerReceiver(1)
	fmt.Printf("Point after Method Pointer Receiver: (%d, %d)\n", p.X, p.Y)
}
