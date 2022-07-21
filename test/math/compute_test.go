package math

import (
	"github.com/uzmijnlm/go_advanced/main/math"
	math2 "math"
	"math/rand"
	"testing"
	"time"
)

func TestPow(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(10) - 5
		n := rand.Intn(10) - 5

		res := math.Pow(x, n)
		expectedRes := int(math2.Pow(float64(x), float64(n)))
		if res != expectedRes {
			t.Fail()
		}
	}
}

func TestGetPrimeNumberCount(t *testing.T) {
	for i := 1; i < 1000; i++ {
		res := math.GetPrimeNumberCount(i)
		expectedRes := getPrimeNumberCount(i)
		if res != expectedRes {
			t.Fail()
		}
	}
}

func getPrimeNumberCount(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	count := 1
	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			} else {
				if j == i-1 {
					count += 1
				}
			}
		}
	}
	return count
}

func TestSqrt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		res := math.Sqrt(i)
		expectedRes := sqrt(i)
		if res != expectedRes {
			t.Fail()
		}
	}

}

func sqrt(x int) int {
	var ans int
	for i := 0; i <= x; i++ {
		if i*i > x {
			break
		}
		ans = i
	}
	return ans
}
