package math

import "math"

// Pow n可以为负数
// 例如10^75，将75拆成二进制，为1001011（64+8+2+1）。
// 先令一个数t=10^1，用变量res表示最终的结果。将二进制数从右往左判断。
// 第一位是1，表明要把t的值算入res，即res=1*t=1*10^1。
// 然后t和自己相乘变为10^2，第二位是1，表明要把此时的t算入res，res=1*10^1*10^2。
// t再和自己乘变为10^4，第三位是0，表明此时的t不算入res。依此类推。
// n为负数时先计算x^-n，再用1/res
func Pow(x int, n int) int {
	if x == 0 && n < 0 {
		return int(math.Inf(1))
	}
	negative := n < 0
	var res = 1
	var factor = x
	var pow int
	if negative {
		pow = -n
	} else {
		pow = n
	}

	for pow != 0 {
		if pow&1 == 1 {
			res = res * factor
		}
		factor = factor * factor
		pow = pow >> 1
	}

	if negative {
		return 1 / res
	} else {
		return res
	}

}

// GetPrimeNumberCount 求1到n有多少个质数
func GetPrimeNumberCount(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	notPrime := make([]bool, n+1)
	count := (n - 1) / 2

	for i := 3; i < n; i = i + 2 {
		if notPrime[i] {
			continue
		}
		for j := i; j < n; j = j + 2 {
			if i*j <= n {
				if !notPrime[i*j] {
					notPrime[i*j] = true
					count -= 1
				}

			} else {
				break
			}
		}
	}
	return count + 1
}

// Sqrt 求x的平方根。x一定非负，返回结果向下取整
// 二分法
func Sqrt(x int) int {
	low := 0
	high := x
	var ans int
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
