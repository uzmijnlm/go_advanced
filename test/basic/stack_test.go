package basic

import (
	"github.com/uzmijnlm/go_advanced/main/basic"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestGetLeftBiggerAndRightBiggerNoRepeat(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(1) + 100
		maxNum := rand.Intn(100) + 100
		var arr []int
		var exists = make(map[int]interface{})
		for i := 0; i < length; i++ {
			var value int
			for {
				value = rand.Intn(maxNum)
				_, ok := exists[value]
				if ok {
					continue
				} else {
					exists[value] = nil
					break
				}
			}
			arr = append(arr, value)
		}

		res := basic.GetLeftBiggerAndRightBiggerNoRepeat(arr)
		expectedRes := getNearBigger(arr)

		if len(res) != len(expectedRes) {
			t.Fail()
		}
		for key, value := range res {
			expectedValue, ok := expectedRes[key]
			if !ok {
				t.Fail()
			} else {
				if value[0] != expectedValue[0] || value[1] != expectedValue[1] {
					t.Fail()
				}
			}
		}
		for key, expectedValue := range expectedRes {
			value, ok := res[key]
			if !ok {
				t.Fail()
			} else {
				if value[0] != expectedValue[0] || value[1] != expectedValue[1] {
					t.Fail()
				}
			}
		}

	}
}

func TestGetLeftBiggerAndRightBiggerWithRepeat(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(1) + 100
		maxNum := rand.Intn(100) + 100
		var arr []int
		var exists = make(map[int]interface{})
		for i := 0; i < length; i++ {
			var value int
			for {
				value = rand.Intn(maxNum)
				_, ok := exists[value]
				if ok {
					continue
				} else {
					exists[value] = nil
					break
				}
			}
			arr = append(arr, value)
		}

		res := basic.GetLeftBiggerAndRightBiggerWithRepeat(arr)
		expectedRes := getNearBigger(arr)

		if len(res) != len(expectedRes) {
			t.Fail()
		}
		for key, value := range res {
			expectedValue, ok := expectedRes[key]
			if !ok {
				t.Fail()
			} else {
				if value[0] != expectedValue[0] || value[1] != expectedValue[1] {
					t.Fail()
				}
			}
		}
		for key, expectedValue := range expectedRes {
			value, ok := res[key]
			if !ok {
				t.Fail()
			} else {
				if value[0] != expectedValue[0] || value[1] != expectedValue[1] {
					t.Fail()
				}
			}
		}

	}
}

func getNearBigger(arr []int) map[int][]int {
	res := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		res[i] = []int{-1, -1}
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[i] {
				res[i][0] = j
				break
			}
		}
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				res[i][1] = j
				break
			}
		}
	}
	return res
}

func TestReverse(t *testing.T) {
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		length := rand.Intn(100-10) + 10
		maxNum := rand.Intn(100) + 1
		stack1 := basic.NewStack()
		stack2 := basic.NewStack()

		for i := 0; i < length; i++ {
			num := rand.Intn(maxNum)
			stack1.Push(num)
			stack2.Push(num)
		}

		basic.Reverse(stack1)
		stack3 := reverse(stack2)

		for !stack1.IsEmpty() {
			if stack1.Pop().(int) != stack3.Pop().(int) {
				t.Fail()
			}
		}
	}
}

func reverse(stack *basic.Stack) *basic.Stack {
	newStack := basic.NewStack()
	for !stack.IsEmpty() {
		newStack.Push(stack.Pop())
	}
	return newStack
}

func TestSortByAnotherStack(t *testing.T) {
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

		stack1 := basic.NewStack()
		stack2 := basic.NewStack()
		for i := 0; i < len(arr); i++ {
			stack1.Push(arr1[i])
			stack2.Push(arr2[i])
		}

		basic.SortByAnotherStack(stack2)

		for !stack1.IsEmpty() {
			if stack1.Pop().(int) != stack2.Pop().(int) {
				t.Fail()
			}
		}
		if !stack2.IsEmpty() {
			t.Fail()
		}
	}
}
