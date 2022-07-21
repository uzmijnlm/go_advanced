package basic

type Stack struct {
	elements []interface{}
	size     int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(element interface{}) {
	stack.elements = append(stack.elements, element)
	stack.size += 1
}

func (stack *Stack) Pop() interface{} {
	if stack.size <= 0 {
		return nil
	} else {
		element := stack.elements[stack.size-1]
		stack.elements = stack.elements[0 : stack.size-1]
		stack.size -= 1
		return element
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.size <= 0
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.elements[stack.size-1]
}

func GetLeftBiggerAndRightBiggerNoRepeat(arr []int) map[int][]int {
	stack := NewStack()
	if arr == nil || len(arr) == 0 {
		return nil
	}

	res := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && arr[stack.Peek().(int)] < arr[i] {
			index := stack.Pop().(int)
			var leftIndex int
			if stack.IsEmpty() {
				leftIndex = -1
			} else {
				leftIndex = stack.Peek().(int)
			}
			res[index] = []int{leftIndex, i}
		}
		stack.Push(i)
	}

	for !stack.IsEmpty() {
		index := stack.Pop().(int)
		var leftIndex int
		if stack.IsEmpty() {
			leftIndex = -1
		} else {
			leftIndex = stack.Peek().(int)
		}
		res[index] = []int{leftIndex, -1}
	}
	return res
}

func GetLeftBiggerAndRightBiggerWithRepeat(arr []int) map[int][]int {
	stack := NewStack()
	if arr == nil || len(arr) == 0 {
		return nil
	}

	res := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && arr[stack.Peek().([]int)[0]] < arr[i] {
			indexes := stack.Pop().([]int)
			var leftIndex int
			if stack.IsEmpty() {
				leftIndex = -1
			} else {
				leftIndex = stack.Peek().([]int)[0]
			}
			for j := 0; j < len(indexes); j++ {
				res[indexes[j]] = []int{leftIndex, i}
			}
		}
		if !stack.IsEmpty() && arr[stack.Peek().([]int)[0]] == arr[i] {
			indexes := stack.Peek().([]int)
			indexes = append(indexes, i)
		} else {
			indexes := []int{i}
			stack.Push(indexes)
		}
	}

	for !stack.IsEmpty() {
		indexes := stack.Pop().([]int)
		var leftIndex int
		if stack.IsEmpty() {
			leftIndex = -1
		} else {
			leftIndex = stack.Peek().([]int)[0]
		}
		for j := 0; j < len(indexes); j++ {
			res[indexes[j]] = []int{leftIndex, -1}
		}
	}
	return res
}

// Reverse 仅用递归函数和栈操作实现逆序一个栈
func Reverse(stack *Stack) {
	if stack.IsEmpty() {
		return
	}
	element := getAndRemoveLastElement(stack)
	Reverse(stack)
	stack.Push(element)
}

func getAndRemoveLastElement(stack *Stack) interface{} {
	result := stack.Pop()
	if stack.IsEmpty() {
		return result
	} else {
		last := getAndRemoveLastElement(stack)
		stack.Push(result) // 弹出最后一个元素，并且先前弹出的元素都放回后，最后放入最先弹出的元素，由此其余元素顺序不变
		return last
	}
}

func SortByAnotherStack(stack *Stack) {
	helpStack := NewStack()
	for !stack.IsEmpty() {
		cur := stack.Pop().(int)
		for !helpStack.IsEmpty() && helpStack.Peek().(int) < cur {
			stack.Push(helpStack.Pop())
		}
		helpStack.Push(cur)
	}
	for !helpStack.IsEmpty() {
		stack.Push(helpStack.Pop())
	}
}

// GetMinStack1 有GetMin功能的栈
// 所有操作的时间复杂度均为O(1)
type GetMinStack1 struct {
	stackData *Stack
	stackMin  *Stack
}

func NewGetMinStack1() *GetMinStack1 {
	return &GetMinStack1{
		stackData: NewStack(),
		stackMin:  NewStack(),
	}
}

func (stack *GetMinStack1) Push(element int) {
	if stack.stackMin.IsEmpty() {
		stack.stackMin.Push(element)
	} else if element <= stack.GetMin() {
		stack.stackMin.Push(element)
	}
	stack.stackData.Push(element)
}

func (stack *GetMinStack1) Pop() int {
	value := stack.stackData.Pop().(int)
	if value == stack.GetMin() {
		stack.stackMin.Pop()
	}
	return value
}

func (stack *GetMinStack1) GetMin() int {
	return stack.stackMin.Peek().(int)
}

type GetMinStack2 struct {
	stackData *Stack
	stackMin  *Stack
}

func NewGetMinStack2() *GetMinStack2 {
	return &GetMinStack2{
		stackData: NewStack(),
		stackMin:  NewStack(),
	}
}

func (stack *GetMinStack2) Push(element int) {
	if stack.stackMin.IsEmpty() {
		stack.stackMin.Push(element)
	} else if element < stack.GetMin() {
		stack.stackMin.Push(element)
	} else {
		newMin := stack.stackMin.Peek()
		stack.stackMin.Push(newMin)
	}
	stack.stackData.Push(element)
}

func (stack *GetMinStack2) Pop() int {
	value := stack.stackData.Pop().(int)
	stack.stackMin.Pop()
	return value
}

func (stack *GetMinStack2) GetMin() int {
	return stack.stackMin.Peek().(int)
}

type QueueByStack struct {
	stackPush *Stack
	stackPop  *Stack
}

func NewQueueByStack() *QueueByStack {
	return &QueueByStack{stackPush: NewStack(), stackPop: NewStack()}
}

func (queue *QueueByStack) Push(element interface{}) {
	queue.stackPush.Push(element)
	queue.pushToPop()
}

func (queue *QueueByStack) pushToPop() {
	if queue.stackPop.IsEmpty() {
		for !queue.stackPush.IsEmpty() {
			queue.stackPop.Push(queue.stackPush.Pop())
		}
	}
}

func (queue *QueueByStack) Pop() interface{} {
	queue.pushToPop()
	return queue.stackPop.Pop()
}

func (queue *QueueByStack) Peek() interface{} {
	queue.pushToPop()
	return queue.stackPop.Peek()
}

func (queue *QueueByStack) IsEmpty() bool {
	queue.pushToPop()
	return queue.stackPop.IsEmpty()
}
