package main

import "fmt"

func main() {
	var m map[int]string = make(map[int]string)
	fmt.Println(m, len(m))

	var n map[int]string = map[int]string{
		90: "优秀",
		80: "良好",
		60: "及格", // 注意这里逗号不可缺少，否则会报语法错误
	}
	fmt.Println(n, len(n))
}
