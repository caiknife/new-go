package main

import "fmt"

type Person interface {
	Name() string
	Age() uint
}

func main() {
	fmt.Println("Hello, world!")
}
