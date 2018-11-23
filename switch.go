package main

import "fmt"

func prize1(score int) string {
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "不及格"
	case 6, 7:
		return "及格"
	case 8:
		return "良"
	default:
		return "优"
	}
}

func prize2(score int) string {
	switch {
	case score < 60:
		return "差"
	case score < 80:
		return "及格"
	case score < 90:
		return "良"
	default:
		return "优"
	}
}

func main() {
	fmt.Println(prize1(60))
	fmt.Println(prize2(60))
}
