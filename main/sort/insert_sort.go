package sort

func InsertSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] >= arr[j-1] {
				break
			}
			arr[j] ^= arr[j-1]
			arr[j-1] ^= arr[j]
			arr[j] ^= arr[j-1]
		}
	}
}
