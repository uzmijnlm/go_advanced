package sort

import "math"

// BucketSort 桶排序，又称RadixSort，基数排序
func BucketSort(arr []int) {
	// 获取数组最大值有几位
	maxBits := getMaxBits(arr)

	buckets := make([]Bucket, 10)

	// 最大值有几位，表示所有数要入桶出桶几次
	for bit := 0; bit < maxBits; bit++ {
		// 入桶，将每个数插入到对应桶的队列中
		for j := 0; j < len(arr); j++ {
			divided := int(math.Pow(10, float64(bit)))
			mod := (arr[j] / divided) % 10
			buckets[mod].insert(arr[j])
		}

		// 出桶，按桶的顺序依次将数从队列中弹出放回原数组
		index := 0
		for j := 0; j < 10; j++ {
			for buckets[j].size != 0 {
				arr[index] = buckets[j].pop()
				index += 1
			}
		}
	}
}

func getMaxBits(arr []int) int {
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	result := 0
	for max != 0 {
		result += 1
		max /= 10
	}
	return result
}

type Bucket struct {
	queue []int
	size  int
}

func (bucket *Bucket) insert(num int) {
	bucket.queue = append(bucket.queue[0:bucket.size], num)
	bucket.size += 1
}

func (bucket *Bucket) pop() int {
	result := bucket.queue[0]
	bucket.queue = bucket.queue[1:bucket.size]
	bucket.size -= 1
	return result
}

func BucketSortOptimized(arr []int) {
	// 获取数组最大值有几位
	maxBits := getMaxBits(arr)

	// 每一位index表示某一位上值小于等于index的数有多少个
	// 例如arr = {11, 53, 43, 24}，则遍历个位时，count = {0, 1, 0, 3, 4, 4, 4, 4, 4, 4}
	// count[1] = 1表示arr中的数，个位小于等于1的数有1个，即11
	// count[3] = 3表示arr中的数，个位小于等于3的数有3个，即11、53、43
	count := make([]int, 10)

	for bit := 0; bit < maxBits; bit++ {
		// 计数
		for i := 0; i < len(arr); i++ {
			divided := int(math.Pow(10, float64(bit)))
			mod := (arr[i] / divided) % 10
			count[mod] += 1
		}
		// 将计数改为前缀和
		for i := 1; i < 10; i++ {
			count[i] = count[i] + count[i-1]
		}

		// 出桶，原数组从右往左，依次看该进制位上这个数在count数组中的值
		// 假设这个数是192，本次循环看个位，则看count[2]的值，若值为5，则说明按照个位排序，这个数应排在临时数组第4（5-1）位
		// 然后count[2] -= 1，下一次遇到另一个个位为2的数，则排到临时数组第3（4-1）位
		// 注意在这种算法下，一定要从右往左，这样才能保证排序的稳定性，否则排到后面进制时，前面进制的顺序就乱了
		// 如果想从左往右遍历，则计算临时数组索引时则很麻烦，就不能用count[2]-1这种方法
		// 因为count[2]-1表示的是将这个数放在位数为2的数的最后，正好与从右往左的顺序一致，由此保证排序的稳定性
		temp := make([]int, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			divided := int(math.Pow(10, float64(bit)))
			mod := (arr[i] / divided) % 10
			temp[count[mod]-1] = arr[i]
			count[mod] -= 1
		}

		for i := 0; i < len(arr); i++ {
			arr[i] = temp[i]
		}

		count = make([]int, 10)
	}

}
