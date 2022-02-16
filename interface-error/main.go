package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	var err error // 'error' is an interface
	_ = err

	var err0 error = errors.New("simple way to create an error") // it returns an instance of errors.errorString struct
	fmt.Println(err0)

	var err1 error = fmt.Errorf("%s way to create an error", "the most common") // it returns an instance of errors.errorString struct
	fmt.Println(err1)

	// syscall.Errno is an other type of error for low-level system-call API
	var err2 error = syscall.Errno(2)
	fmt.Println(err2) // it prints 'no such file or directory'
}
