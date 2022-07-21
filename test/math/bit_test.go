package math

import (
	"github.com/uzmijnlm/go_advanced/main/math"
	"math/rand"
	"testing"
	"time"
)

func TestSwap(t *testing.T) {
	num1 := 2
	num2 := 3
	math.Swap(&num1, &num2)
	if num1 != 3 || num2 != 2 {
		t.Fail()
	}

	arr := []int{1, 2}
	math.Swap(&arr[0], &arr[1])
	if arr[0] != 2 || arr[1] != 1 {
		t.Fail()
	}

	num := 5
	math.Swap(&num, &num)
	if num != 5 {
		t.Fail()
	}
}

func TestFindOneNumber(t *testing.T) {
	// 只有一个数出现了奇数次，其他数都出现偶数次
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4}
	result := math.FindOneNumber(arr)
	if result != 4 {
		t.Fail()
	}
}

func TestFindTwoNumbers(t *testing.T) {
	// 有两个数出现了奇数次，其他数都出现偶数次
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6}
	a, b := math.FindTwoNumbers(arr)
	if !((a == 4 && b == 5) || (a == 5 && b == 4)) {
		t.Fail()
	}
}

func TestFindRightestOne(t *testing.T) {
	num := 24
	rightestOne := math.FindRightestOne(num)
	if rightestOne != 8 {
		t.Fail()
	}
}

func TestFindLarger(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(10000) - 5000
		var b int
		for {
			b = rand.Intn(10000) - 5000
			if a != b {
				break
			}
		}

		var res int
		if a > b {
			res = a
		} else {
			res = b
		}
		if math.FindLarger(a, b) != res {
			t.Fail()
		}
	}
}
