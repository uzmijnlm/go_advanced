package sort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, low int, high int) {
	if low < high {
		mid := low + (high-low)/2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}

func merge(arr []int, low int, mid int, high int) {
	var temp = make([]int, high-low+1)
	l := low
	r := mid + 1
	k := 0
	for l <= mid && r <= high {
		if arr[l] <= arr[r] {
			temp[k] = arr[l]
			k += 1
			l += 1
		} else {
			temp[k] = arr[r]
			k += 1
			r += 1
		}
	}

	for l <= mid || r <= high {
		if l > mid {
			temp[k] = arr[r]
			k += 1
			r += 1
		} else {
			temp[k] = arr[l]
			k += 1
			l += 1
		}
	}

	p := low
	for i := 0; i < len(temp); i++ {
		arr[p] = temp[i]
		p += 1
	}
}

// GetSmallSum 获取一个数组的小和
// 求一个数左边比它小的数的和，对数组中每个数都做这个操作，将所有和累加
func GetSmallSum(arr []int) int {
	return getSmallSum(arr, 0, len(arr)-1)
}

func getSmallSum(arr []int, low int, high int) int {
	if low < high {
		mid := low + (high-low)/2
		return getSmallSum(arr, low, mid) + getSmallSum(arr, mid+1, high) + mergeSmallSum(arr, low, mid, high)
	} else {
		return 0
	}
}

func mergeSmallSum(arr []int, low int, mid int, high int) int {
	temp := make([]int, high-low+1)
	l := low
	r := high
	k := 0
	sum := 0

	for l <= mid && r <= high {
		if arr[l] < arr[r] {
			sum += arr[l] * (high - r + 1)
			temp[k] = arr[l]
			k += 1
			l += 1
		} else {
			temp[k] = arr[r]
			k += 1
			r += 1
		}
	}

	for l <= mid || r <= high {
		if l > mid {
			temp[k] = arr[r]
			k += 1
			r += 1
		} else {
			temp[k] = arr[l]
			k += 1
			l += 1
		}
	}

	p := low
	for i := 0; i < len(temp); i++ {
		arr[p] = temp[i]
		p += 1
	}

	return sum
}
