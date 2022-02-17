package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Reader
	fmt.Printf("w: %T, %v\n", w, w) // interface value dynamic type is nil

	w = os.Stdin
	fmt.Printf("w: %T, %v\n", w, w) // interface value dynamic type is *os.File

	// if the asserted type (*os.File) is a concrete type, then the type assertion checks
	// wheter the dynamic value of the interface value (w) is identical to the asserted concrete type.
	// and extract the concrete type
	var f1 *os.File = w.(*os.File) // success
	fmt.Printf("f1: %T, %v\n", f1, f1)

	// When asserted concrete type does not match the dynamic type of the interface value (w), then type assertion panics
	// var f1p *bytes.Buffer = w.(*bytes.Buffer) // panic: interface conversion: io.Reader is *os.File, not *bytes.Buffer

	// if the asserted type (*os.File) is a interface type, then the type assertion checks
	// wheter the dynamic value of the interface value (w) satisfy the asserted interface type.
	// Changes the expresion type, preserves de dynamic type and dynamic value inside the result interface value (f2)
	var f2 io.ReadWriter = w.(io.ReadWriter) // success
	fmt.Printf("f2: %T, %v\n", f2, f2)

	// When asserted interface type is not satisfied by the dynamic value of the interface value (w), then type assertion panics
	// var f2p error = w.(error) // panic: interface conversion: *os.File is not error: missing method Error

	// To avoid panic on type assertion, it must appears in an assignemt with two results expected
	f3, ok := w.(error)
	if !ok {
		fmt.Printf("f3: %T, %v\n", f3, f3)
	}
	f4, ok := w.(io.ReadWriter)
	if ok {
		fmt.Printf("f4: %T, %v\n", f4, f4)
	}
}
