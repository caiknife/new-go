package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Block 1
	ch1 := make(chan int, 5)
	ch1 <- 3
	ch1 <- 2
	ch1 <- 1
	ch1 <- 4

	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)

	rand.Seed(time.Now().UnixNano())
	// Block 2
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// Block 3
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("The index: %d \n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[1]:
		fmt.Println("The second chan is selected.")
	case <-intChannels[0]:
		fmt.Println("The first chan is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third chan is selected. The element is %d. \n", elem)
	default:
		fmt.Println("No chan is selected")
	}
}

/*
函数声明，返回一个单向通道，只能发送，不能接收
 */
func getIntChan() <-chan int {
	num := 5 // 通道容量 5

	ch := make(chan int, num) // 创建一个通道

	for i := 0; i < num; i++ {
		ch <- rand.Intn(100) // 每次传入时随机生成一个[0, 100)的整数
	}

	close(ch)

	return ch
}
