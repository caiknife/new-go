package main

import "fmt"

func main() {
	var a [9]int
	fmt.Println(a)

	var aa = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9,}
	fmt.Println(aa)

	var b [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0,}
	fmt.Println(b)

	c := [8]int{1, 2, 3, 4, 5, 6, 7, 8,}
	fmt.Println(c)

	var squares [9]int
	for i := 0; i < len(squares); i++ {
		squares[i] = (i + 1) * (i + 1)
	}
	fmt.Println(squares)

	for index := range b {
		fmt.Println(index, b[index])
	}
}
