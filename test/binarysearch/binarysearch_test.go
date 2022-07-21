package binarysearch

import (
	"fmt"
	"github.com/uzmijnlm/go_advanced/main/binarysearch"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestRecurBinarySearch(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		sort.Ints(arr)
		target := rand.Intn(maxNum)
		found := false
		for i := 0; i < len(arr); i++ {
			if arr[i] == target {
				found = true
				break
			}
		}
		result := binarysearch.RecurBinarySearch(arr, target)
		if result != found {
			t.Fail()
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		sort.Ints(arr)
		target := rand.Intn(maxNum)
		found := false
		for i := 0; i < len(arr); i++ {
			if arr[i] == target {
				found = true
				break
			}
		}
		result := binarysearch.IterBinarySearch(arr, target)
		if result != found {
			t.Fail()
		}
	}
}

func TestRecurBinarySearchSmallestBiggerOne(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		sort.Ints(arr)
		target := rand.Intn(maxNum)

		expectedResult := 0
		found := false
		for i := 0; i < len(arr); i++ {
			if arr[i] > target {
				expectedResult = i
				found = true
				break
			}
		}
		if !found {
			expectedResult = -1
		}

		result := binarysearch.RecurFindSmallestBiggerOne(arr, target)

		if result != expectedResult {
			t.Fail()
		}

	}
}

func TestIterBinarySearchSmallestBiggerOne(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		var arr []int
		for i := 0; i < length; i++ {
			arr = append(arr, rand.Intn(maxNum))
		}
		sort.Ints(arr)
		target := rand.Intn(maxNum)

		expectedResult := 0
		found := false
		for i := 0; i < len(arr); i++ {
			if arr[i] > target {
				expectedResult = i
				found = true
				break
			}
		}
		if !found {
			expectedResult = -1
		}

		result := binarysearch.IterFindSmallestBiggerOne(arr, target)

		if result != expectedResult {
			fmt.Println(arr)
			fmt.Println(target)
			t.Fail()
		}

	}
}

func TestFindLocalSmallest(t *testing.T) {
	for i := 0; i < 1000; i++ {
		arr := make([]int, 0)
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(1) + 5
		maxNum := rand.Intn(100) + 20
		var existNums = make(map[int]bool)
		for i := 0; i < length; i++ {
			for {
				value := rand.Intn(maxNum)
				_, ok := existNums[value]
				if !ok {
					arr = append(arr, value)
					existNums[value] = true
					break
				}
			}
		}
		result := binarysearch.FindLocalSmallest(arr)
		if result == 0 {
			if arr[result] > arr[result+1] {
				t.Fail()
			}
		} else if result == (len(arr) - 1) {
			if arr[result] > arr[result-1] {
				t.Fail()
			}
		} else {
			if !(arr[result] < arr[result-1] && arr[result] < arr[result+1]) {
				t.Fail()
			}
		}
	}

}
