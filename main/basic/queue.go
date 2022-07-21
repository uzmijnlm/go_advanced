package basic

type Queue struct {
	elements []interface{}
	size     int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Push(element interface{}) {
	queue.elements = append(queue.elements, element)
	queue.size += 1
}

func (queue *Queue) Pop() interface{} {
	if queue.size <= 0 {
		return nil
	} else {
		element := queue.elements[0]
		queue.elements = queue.elements[1:queue.size]
		queue.size -= 1
		return element
	}
}

func (queue *Queue) IsEmpty() bool {
	return queue.size <= 0

}

func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	return queue.elements[0]
}

type StackByQueue struct {
	queueForUser   *Queue
	queueForBuffer *Queue
}

func NewStackByQueue() *StackByQueue {
	return &StackByQueue{queueForUser: NewQueue(), queueForBuffer: NewQueue()}
}

func (stack *StackByQueue) Push(element interface{}) {
	stack.queueForUser.Push(element)
}

func (stack *StackByQueue) Pop() interface{} {
	for stack.queueForUser.size > 1 {
		element := stack.queueForUser.Pop()
		stack.queueForBuffer.Push(element)
	}
	res := stack.queueForUser.Pop()
	stack.queueForUser, stack.queueForBuffer = stack.queueForBuffer, stack.queueForUser
	return res
}

func (stack *StackByQueue) IsEmpty() bool {
	return stack.queueForUser.IsEmpty()
}

func (stack *StackByQueue) Peek() interface{} {
	return stack.queueForUser.Peek()
}

type Deque struct {
	elements []interface{}
	size     int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (deque *Deque) PollFirst() interface{} {
	if deque.size <= 0 {
		return nil
	} else {
		element := deque.elements[0]
		deque.elements = deque.elements[1:deque.size]
		deque.size -= 1
		return element
	}
}

func (deque *Deque) PollLast() interface{} {
	if deque.size <= 0 {
		return nil
	} else {
		element := deque.elements[deque.size-1]
		deque.elements = deque.elements[0 : deque.size-1]
		deque.size -= 1
		return element
	}
}

func (deque *Deque) AddFirst(element interface{}) {
	deque.elements = append([]interface{}{element}, deque.elements...)
	deque.size += 1
}

func (deque *Deque) AddLast(element interface{}) {
	deque.elements = append(deque.elements, element)
	deque.size += 1
}

func (deque *Deque) IsEmpty() bool {
	return deque.size <= 0
}

func (deque *Deque) PeekFirst() interface{} {
	if deque.IsEmpty() {
		return nil
	}
	return deque.elements[0]
}

func (deque *Deque) PeekLast() interface{} {
	if deque.IsEmpty() {
		return nil
	}
	return deque.elements[deque.size-1]
}
