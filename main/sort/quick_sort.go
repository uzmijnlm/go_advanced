package sort

import (
	"math/rand"
	"time"
)

func SplitTwoParts(arr []int, num int) {
	left := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= num {
			arr[i], arr[left+1] = arr[left+1], arr[i]
			left += 1
		}
	}
}

func SplitThreeParts(arr []int, num int) {
	left := -1
	right := len(arr)
	i := 0
	for i < right && left <= right {
		if arr[i] < num {
			arr[i], arr[left+1] = arr[left+1], arr[i]
			left += 1
			i += 1
			continue
		}
		if arr[i] > num {
			arr[i], arr[right-1] = arr[right-1], arr[i]
			right -= 1
			continue
		}
		if arr[i] == num {
			i += 1
		}
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(high-low+1) + low
		arr[randomIndex], arr[high] = arr[high], arr[randomIndex]
		left, right := partition(arr, low, high)
		quickSort(arr, low, left-1)
		quickSort(arr, right+1, high)
	}
}

func partition(arr []int, low int, high int) (int, int) {
	num := arr[high]
	left := low - 1
	right := high + 1
	i := low
	for i < right && left < right {
		if arr[i] < num {
			arr[i], arr[left+1] = arr[left+1], arr[i]
			left += 1
			i += 1
			continue
		}
		if arr[i] > num {
			arr[i], arr[right-1] = arr[right-1], arr[i]
			right -= 1
			continue
		}
		if arr[i] == num {
			i += 1
			continue
		}
	}
	return left + 1, right - 1
}
