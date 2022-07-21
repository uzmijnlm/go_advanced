package sort

import (
	"fmt"
	mySort "github.com/uzmijnlm/go_advanced/main/sort"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSelectSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.SelectSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.BubbleSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestInsertSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.InsertSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.MergeSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestGetSmallSum(t *testing.T) {
	var arr = []int{3, 4, 5, 6, 1, 2}
	sum := mySort.GetSmallSum(arr)
	expectedSum := 3 + (3 + 4) + (3 + 4 + 5) + 1

	if sum != expectedSum {
		t.Fail()
	}
}

func TestSplitTwoParts(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var num = rand.Intn(maxNum)
		var isSmaller = true
		mySort.SplitTwoParts(arr, num)
		for i := 0; i < len(arr); i++ {
			if !isSmaller {
				if arr[i] <= num {
					t.Fail()
				}
			} else {
				if arr[i] > num {
					isSmaller = false
				}
			}
		}
	}
}

func TestSplitThreeParts(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 1
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var num = rand.Intn(maxNum)

		var meetEqual = false
		var meetBigger = false
		mySort.SplitThreeParts(arr, num)
		for i := 0; i < len(arr); i++ {
			if meetEqual {
				if arr[i] < num {
					fmt.Println(arr)
					fmt.Println(num)
					t.Fail()
				}
			}
			if meetBigger {
				if arr[i] <= num {
					fmt.Println(arr)
					fmt.Println(num)
					t.Fail()
				}
			}
			if !meetEqual && !meetBigger {
				if arr[i] == num {
					meetEqual = true
				} else if arr[i] > num {
					meetBigger = true
				}
			}
			if meetEqual {
				if arr[i] > num {
					meetBigger = true
				}
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.QuickSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.HeapSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestHeapSortOptimized(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.HeapSortOptimized(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestMyPriorityQueue(t *testing.T) {
	var arr = []int{3, 4, 5, 1, 2}
	var queue = mySort.MyPriorityQueue{}
	for i := 0; i < len(arr); i++ {
		queue.Insert(arr[i])
	}
	lastNum := 6
	for i := 0; i < len(arr); i++ {
		last := queue.GetMax()
		fmt.Print(last)
		if lastNum < last {
			t.Fail()
		}
		lastNum = last
	}
}

func TestSortK(t *testing.T) {
	k := 2
	var arr = []int{1, 3, 0, 2, 5, 4, 5}
	var arr1 = make([]int, len(arr))
	var arr2 = make([]int, len(arr))
	copy(arr1, arr)
	copy(arr2, arr)
	mySort.KSort(arr1, k)
	mySort.QuickSort(arr2)
	if len(arr1) != len(arr2) {
		t.Fatal("长度不一致")
	}
	for i := 0; i < len(arr); i++ {
		if arr1[i] != arr2[i] {
			fmt.Println(arr)
			t.Fatal("排序错误")
		}
	}
}

func TestGetMedian(t *testing.T) {
	var arr1 = []int{1, 3, 0, 2, 5, 4, 5}
	var expectedMedian1 = float64(3)
	heap1 := mySort.MedianHeap{}
	for i := 0; i < len(arr1); i++ {
		heap1.Insert(arr1[i])
	}
	median1 := heap1.GetMedian()
	if median1 != expectedMedian1 {
		t.Fail()
	}

	var arr2 = []int{1, 3, 0, 2, 5, 4, 5, 4}
	var expectedMedian2 = 3.5
	heap2 := mySort.MedianHeap{}
	for i := 0; i < len(arr2); i++ {
		heap2.Insert(arr2[i])
	}
	median2 := heap2.GetMedian()
	if median2 != expectedMedian2 {
		t.Fail()
	}
}

func TestMyDictionary(t *testing.T) {
	var arr = []string{"a", "b", "a", "b", "c", "d", "d", "d"}
	dict := mySort.MyDictionary{
		K:          2,
		Word2Freq:  make(map[string]int),
		Word2Index: make(map[string]int),
	}
	for i := 0; i < len(arr); i++ {
		dict.Insert(arr[i])
	}
	dict.GetTopK()
}

func TestCountSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.CountSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestBucketSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 1
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.BucketSort(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}

func TestBucketSortOptimized(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 1
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		var arr1 = make([]int, len(arr))
		var arr2 = make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		sort.Ints(arr1)
		mySort.BucketSortOptimized(arr2)
		if len(arr1) != len(arr2) {
			t.Fatal("长度不一致")
		}
		for i := 0; i < len(arr); i++ {
			if arr1[i] != arr2[i] {
				fmt.Println(arr)
				t.Fatal("排序错误")
			}
		}
	}
}
