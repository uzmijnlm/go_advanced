package sort

import "math"

func CountSort(arr []int) {
	max := math.MinInt
	min := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	temp := make([]int, max-min+1)

	base := min

	for i := 0; i < len(arr); i++ {
		value := arr[i]
		index := value - base
		temp[index] += 1
	}

	index := 0
	for i := 0; i < len(temp); i++ {
		for temp[i] != 0 {
			arr[index] = i + base
			index += 1
			temp[i] -= 1
		}
	}
}
