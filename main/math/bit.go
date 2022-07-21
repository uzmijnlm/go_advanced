package math

func Swap(num1 *int, num2 *int) {
	if num1 == num2 {
		return
	}
	*num1 ^= *num2
	*num2 ^= *num1
	*num1 ^= *num2
}

func FindOneNumber(arr []int) int {
	var result = arr[0]
	for i := 1; i < len(arr); i++ {
		result ^= arr[i]
	}
	return result
}

func FindTwoNumbers(arr []int) (int, int) {
	var a, b int
	var p int
	p = arr[0]
	for i := 1; i < len(arr); i++ {
		p ^= arr[i]
	}

	rightestOne := FindRightestOne(p)

	for i := 0; i < len(arr); i++ {
		if arr[i]&rightestOne != 0 {
			a ^= arr[i]
		}
	}

	b = a ^ p

	return a, b
}

func FindRightestOne(num int) int {
	return num & (^num + 1)
}

func FindLarger(a int, b int) int {
	sa := sign(a)
	sb := sign(b)
	sc := sign(a - b)
	diff := sa ^ sb
	same := flip(diff)
	returnA := diff*sa + same*sc
	returnB := flip(returnA)
	return returnA*a + returnB*b
}

func sign(num int) int {
	return flip(num >> 31 & 1)
}

func flip(num int) int {
	return num ^ 1
}
