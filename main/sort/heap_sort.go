package sort

import (
	"fmt"
	"strconv"
)

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index int, heapSize int) {
	for index*2+1 < heapSize {
		var bigger int
		if index*2+2 < heapSize && arr[index*2+1] < arr[index*2+2] {
			bigger = index*2 + 2
		} else {
			bigger = index*2 + 1
		}

		if arr[index] <= arr[bigger] {
			arr[index], arr[bigger] = arr[bigger], arr[index]
			index = bigger
		} else {
			break
		}
	}
}

func HeapSort(arr []int) {
	// 构建大根堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}

	heapSize := len(arr)
	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	heapSize -= 1
	for heapSize > 0 {
		// 在0到heapSize范围上将最大值放到堆顶
		heapify(arr, 0, heapSize)
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize -= 1
	}
}

func HeapSortOptimized(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	heapSize := len(arr)
	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	heapSize -= 1
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize -= 1
	}
}

type MyPriorityQueue struct {
	arr  []int
	size int
}

func (queue *MyPriorityQueue) Insert(num int) {
	queue.arr = append(queue.arr[0:queue.size], num)
	heapInsert(queue.arr, queue.size)
	queue.size += 1
}

func (queue *MyPriorityQueue) PrintArr() {
	fmt.Println(queue.arr)
}

func (queue *MyPriorityQueue) GetMax() int {
	maxNum := queue.arr[0]
	queue.arr[0], queue.arr[queue.size-1] = queue.arr[queue.size-1], queue.arr[0]
	queue.size -= 1
	heapify(queue.arr, 0, queue.size)
	return maxNum
}

func (queue *MyPriorityQueue) isEmpty() bool {
	return queue.size == 0
}

func KSort(arr []int, k int) {
	index := len(arr) - 1
	queue := MyPriorityQueue{arr: make([]int, 10)}
	for i := len(arr) - 1; i >= 0; i-- {
		queue.Insert(arr[i])
		if i < len(arr)-1-k {
			arr[index] = queue.GetMax()
			index -= 1
		}
	}

	for !queue.isEmpty() {
		arr[index] = queue.GetMax()
		index -= 1
	}
}

type MedianHeap struct {
	smallHeap []int
	smallSize int
	bigHeap   []int
	bigSize   int
}

func (heap *MedianHeap) Insert(num int) {
	if heap.smallSize+heap.bigSize == 0 {
		heap.bigHeap = append(heap.bigHeap[0:heap.bigSize], num)
		heap.bigSize += 1
	} else if num < heap.bigHeap[heap.bigSize-1] {
		heap.bigHeap = append(heap.bigHeap[0:heap.bigSize], num)
		heapInsertBig(heap.bigHeap, heap.bigSize)
		heap.bigSize += 1
	} else {
		heap.smallHeap = append(heap.smallHeap[0:heap.smallSize], num)
		heapInsertSmall(heap.smallHeap, heap.smallSize)
		heap.smallSize += 1
	}
	for heap.bigSize-heap.smallSize >= 2 {
		maxNum := heap.bigHeap[0]
		heap.bigHeap[0], heap.bigHeap[heap.bigSize-1] = heap.bigHeap[heap.bigSize-1], heap.bigHeap[0]
		heap.bigSize -= 1
		heapifyBig(heap.bigHeap, 0, heap.bigSize)

		heap.smallHeap = append(heap.smallHeap[0:heap.smallSize], maxNum)
		heapInsertSmall(heap.smallHeap, heap.smallSize)
		heap.smallSize += 1
	}

	for heap.smallSize-heap.bigSize >= 2 {
		minNum := heap.smallHeap[0]
		heap.smallHeap[0], heap.smallHeap[heap.smallSize-1] = heap.smallHeap[heap.smallSize-1], heap.smallHeap[0]
		heap.smallSize -= 1
		heapifySmall(heap.smallHeap, 0, heap.smallSize)

		heap.bigHeap = append(heap.bigHeap[0:heap.bigSize], minNum)
		heapInsertBig(heap.bigHeap, heap.bigSize)
		heap.bigSize += 1
	}
}

func (heap *MedianHeap) GetMedian() float64 {
	if heap.bigSize > heap.smallSize {
		return float64(heap.bigHeap[0])
	} else if heap.smallSize > heap.bigSize {
		return float64(heap.smallHeap[0])
	} else {
		return float64(heap.bigHeap[0]+heap.smallHeap[0]) / float64(2)
	}
}

func heapInsertBig(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapInsertSmall(arr []int, index int) {
	for arr[index] < arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapifyBig(arr []int, index int, heapSize int) {
	for index*2+1 < heapSize {
		var bigger int
		if index*2+2 < heapSize && arr[index*2+1] < arr[index*2+2] {
			bigger = index*2 + 2
		} else {
			bigger = index*2 + 1
		}

		if arr[index] < arr[bigger] {
			arr[index], arr[bigger] = arr[bigger], arr[index]
			index = bigger
		} else {
			break
		}
	}
}

func heapifySmall(arr []int, index int, heapSize int) {
	for index*2+1 < heapSize {
		var smaller int
		if index*2+2 < heapSize && arr[index*2+1] > arr[index*2+2] {
			smaller = index*2 + 2
		} else {
			smaller = index*2 + 1
		}

		if arr[index] > arr[smaller] {
			arr[index], arr[smaller] = arr[smaller], arr[index]
			index = smaller
		} else {
			break
		}
	}
}

type MyDictionary struct {
	Word2Freq  map[string]int
	smallHeap  []string
	K          int
	heapSize   int
	Word2Index map[string]int
}

func (dict *MyDictionary) Insert(word string) {
	freq, ok := dict.Word2Freq[word]
	if !ok {
		dict.Word2Freq[word] = 1
	} else {
		dict.Word2Freq[word] = freq + 1
	}

	if dict.heapSize < dict.K {
		dict.smallHeap = append(dict.smallHeap[0:dict.heapSize], word)
		dict.Word2Index[word] = dict.heapSize
		dict.heapInsertSmall(dict.smallHeap, dict.heapSize)
		dict.heapSize += 1
	} else {
		if dict.Word2Freq[word] > dict.Word2Freq[dict.smallHeap[0]] {
			dict.smallHeap[0], dict.smallHeap[dict.heapSize-1] = dict.smallHeap[dict.heapSize-1], dict.smallHeap[0]
			delete(dict.Word2Index, dict.smallHeap[dict.heapSize-1])
			dict.heapifySmall(dict.smallHeap, 0, dict.heapSize-1)
			dict.smallHeap[dict.heapSize-1] = word
			dict.heapInsertSmall(dict.smallHeap, dict.heapSize-1)
		}
	}
}

func (dict *MyDictionary) heapInsertSmall(arr []string, index int) {
	for dict.Word2Freq[arr[index]] > dict.Word2Freq[arr[(index-1)/2]] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		dict.Word2Index[arr[index]], dict.Word2Index[arr[(index-1)/2]] = dict.Word2Index[arr[(index-1)/2]], dict.Word2Index[arr[index]]
		index = (index - 1) / 2
	}
}

func (dict *MyDictionary) heapifySmall(arr []string, index int, heapSize int) {
	for index*2+1 < heapSize {
		var smaller int
		if index*2+2 < heapSize && dict.Word2Freq[arr[index*2+1]] > dict.Word2Freq[arr[index*2+2]] {
			smaller = index*2 + 2
		} else {
			smaller = index*2 + 1
		}

		if dict.Word2Freq[arr[index]] > dict.Word2Freq[arr[smaller]] {
			arr[index], arr[smaller] = arr[smaller], arr[index]
			dict.Word2Index[arr[index]], dict.Word2Index[arr[smaller]] = dict.Word2Index[arr[smaller]], dict.Word2Index[arr[index]]
			index = smaller
		} else {
			break
		}
	}
}

func (dict *MyDictionary) GetTopK() {
	for _, word := range dict.smallHeap {
		freq := dict.Word2Freq[word]
		fmt.Println("word is " + word + " and freq is " + strconv.Itoa(freq))
	}
}
