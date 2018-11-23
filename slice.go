package main

import "fmt"

func main() {
	var s1 []int = make([]int, 5, 8)
	var s2 []int = make([]int, 8)

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
