package main

import (
	"fmt"
	"flag"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

var (
	name string
)

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object")

	flag.Parse()
}

func main() {
	fmt.Printf("Hello, %s!\n", name)

	fmt.Println(sign(max(min(24, 42), max(24, 42))))
	fmt.Println(sign(max(min(24, 42), max(24, 42))))
}
