package binarysearch

func RecurBinarySearch(arr []int, num int) bool {
	return recurBinarySearchExist(arr, num, 0, len(arr)-1)
}

func recurBinarySearchExist(arr []int, num int, low int, high int) bool {
	if low < high {
		mid := low + (high-low)/2
		return recurBinarySearchExist(arr, num, low, mid) || recurBinarySearchExist(arr, num, mid+1, high)
	} else if low == high {
		return arr[low] == num
	} else {
		return false
	}
}

func IterBinarySearch(arr []int, num int) bool {
	l := 0
	r := len(arr) - 1
	exist := false
	for l <= r {
		mid := l + ((r - l) / 2)
		if arr[mid] < num {
			l = mid + 1
		} else if arr[mid] > num {
			r = mid - 1
		} else {
			exist = true
			break
		}
	}
	return exist
}

func RecurFindSmallestBiggerOne(arr []int, num int) int {
	return recurBinarySearchSmallestBiggerOne(arr, num, 0, len(arr)-1)
}

func recurBinarySearchSmallestBiggerOne(arr []int, num int, low int, high int) int {
	if low < high {
		mid := low + (high-low)/2
		if arr[mid] <= num {
			return recurBinarySearchSmallestBiggerOne(arr, num, mid+1, high)
		} else {
			return recurBinarySearchSmallestBiggerOne(arr, num, low, mid)
		}
	} else if low == high {
		if arr[low] <= num {
			return -1
		} else {
			return low
		}
	} else {
		return -1
	}
}

func IterFindSmallestBiggerOne(arr []int, num int) int {
	l := 0
	r := len(arr) - 1
	index := -1
	for l <= r {
		mid := l + ((r - l) / 2)
		if arr[mid] > num {
			index = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	if arr[len(arr)-1] < num {
		index = -1
	}

	return index
}

// FindLocalSmallest 数组相邻两个数不等，找到局部最小
func FindLocalSmallest(arr []int) int {
	if arr[0] < arr[1] {
		return 0
	} else if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	} else {
		l := 0
		r := len(arr) - 1
		index := 0
		for l <= r {
			mid := l + ((r - l) >> 1)
			if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
				index = mid
				break
			} else if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
				l = mid
			} else {
				r = mid
			}
		}
		return index
	}
}
