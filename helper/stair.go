package helper

func GoStairWithLoop(s int) int {
	// f(1) = 1
	// f(2) = 2
	// result是f(1)的值 tmp是f(2)-f(1)的值
	result, tmp := 1, 1

	for s != 0 {
		tmp = result + tmp
		result = tmp - result
		s--
	}

	return result
}

func GoStair(s int) int {
	if s < 1 {
		return 0
	}

	if s < 3 {
		return s
	}

	return GoStair(s-1) + GoStair(s-2)
}

func GoStairWithTail(s int, result int, tmp int) int {
	if s < 1 {
		return result
	}

	return GoStairWithTail(s-1, tmp, result+tmp)
}

func GoStairWithSM(s int, m int) int {
	if s == 0 {
		return 0
	}

	if s == 1 {
		return 1
	}

	result := 0

	if s > m {
		for i := 1; i <= m; i++ {
			result += GoStairWithSM(s-i, m)
		}
	} else {
		result = GoStairWithSM(s, s-1) + 1
	}

	return result
}
