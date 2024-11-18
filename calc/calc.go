package calc

func Add(nums... int) int {
	var res int
	for _, num := range nums {
		res += num
	}
	return res
}

func Sub(a, b int) int {
	return a - b
}

func Mul(nums... int) int {
	var res int = 1
	for _, num := range nums {
		res *= num
	}
	return res
}
