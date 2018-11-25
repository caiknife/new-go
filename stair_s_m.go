package main

import (
	"fmt"
	"new-go/helper"
)

func main() {
	// 默认走10阶，每次最多2阶
	fmt.Println(helper.GoStairWithSM(3, 4))
}
