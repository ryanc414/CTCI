package ctci

import "errors"

// A generic stack must implement four methods: Pop, Push, Peek and IsEmpty.
type IntStack interface {
	Pop() (int, error)
	Push(item int)
	Peek() (int, error)
	IsEmpty() bool
}

// A generic queue must implement the same four methods as a stack. Only the
// ordering of popped elements differs - queues are popped in LIFO order
// while stacks are FIFO. For clarity, we define a separate interface
// for queues so that it is clear which ordering is expected.
type IntQueue interface {
	Pop() (int, error)
	Push(item int)
	Peek() (int, error)
	IsEmpty() bool
}

// Basic implementation of a stack, using a dynamic expanding array slice.
type basicStack struct {
	data []int
}

// Construct a new empty stack.
func NewBasicStack() *basicStack {
	return &basicStack{data: nil}
}

// Pop an element from the stack - pops last element from the slice.
func (stack *basicStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	newLen := len(stack.data) - 1
	var popped int
	popped, stack.data = stack.data[newLen], stack.data[:newLen]

	return popped, nil
}

// Push an element onto the stack - item is appended to the slice.
func (stack *basicStack) Push(item int) {
	stack.data = append(stack.data, item)
}

// Peek at the top element in the stack.
func (stack basicStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	return stack.data[len(stack.data)-1], nil
}

// Check if the stack is empty, by checking if the data slice has a non-zero
// length.
func (stack basicStack) IsEmpty() bool {
	return len(stack.data) == 0
}

// Stores three stacks interleaved in a single array.
type multiStacks struct {
	stackData    []int
	stackIndices []int
	numStacks    int
}

// Construct a new multiStacks object.
func NewMultiStacks(numStacks int) multiStacks {
	if numStacks < 1 {
		panic("numStacks must be >0")
	}

	stackData := make([]int, numStacks)
	stackIndices := make([]int, numStacks)
	for i := range stackIndices {
		stackIndices[i] = -1
	}

	return multiStacks{
		stackData:    stackData,
		stackIndices: stackIndices,
		numStacks:    numStacks,
	}
}

// Pops an item off the top of a multiStacks stack. An error is returned if the
// stack is empty.
func (stacks *multiStacks) Pop(stackIx int) (int, error) {
	stacks.checkStackIx(stackIx)

	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.getDataIx(stackIx)
	stacks.stackIndices[stackIx]--

	return stacks.stackData[dataIx], nil
}

// Push an item onto the top of a multiStacks stack. The backing array is
// automatically expanded if required, so this function should always succeed.
func (stacks *multiStacks) Push(stackIx, item int) {
	stacks.checkStackIx(stackIx)

	stacks.stackIndices[stackIx]++
	newDataIx := stacks.getDataIx(stackIx)

	// Expand data slice if required.
	if newDataIx >= len(stacks.stackData) {
		extraReqd := newDataIx + 1 - len(stacks.stackData)
		stacks.stackData = append(stacks.stackData, make([]int, extraReqd)...)
	}

	stacks.stackData[newDataIx] = item
}

// Return the value at the top of a multiStacks stack without removing it. An
// error is returned if the stack is empty.
func (stacks *multiStacks) Peek(stackIx int) (int, error) {
	stacks.checkStackIx(stackIx)

	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.getDataIx(stackIx)
	return stacks.stackData[dataIx], nil
}

// Check if a multiStacks stack is empty.
func (stacks *multiStacks) IsEmpty(stackIx int) bool {
	stacks.checkStackIx(stackIx)
	return stacks.stackIndices[stackIx] == -1
}

// Check if the stackIx is within expected range.
func (stacks multiStacks) checkStackIx(stackIx int) {
	if stackIx < 0 || stackIx > stacks.numStacks {
		panic("stackIx is not within expected range")
	}
}

// Return the current index for a stack within the shared stack data array.
func (stacks multiStacks) getDataIx(stackIx int) int {
	return (stacks.stackIndices[stackIx] * stacks.numStacks) + stackIx
}

// A MinStack and its node. This implements a stack interface with the addition
// of providing O(1) access to the minimum value in the stack. This is achieved
// by storing an extra min value on each stack node, that keeps track of the
// minimal value from that node downwards in the stack.
type MinStackNode struct {
	data int
	next *MinStackNode
	min  int
}

type MinStack struct {
	top *MinStackNode
}

func (stack *MinStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	retVal := stack.top.data
	stack.top = stack.top.next

	return retVal, nil
}

func (stack *MinStack) Push(item int) {
	var newMin int
	if stack.IsEmpty() {
		newMin = item
	} else {
		newMin = min(item, stack.top.min)
	}

	newNode := MinStackNode{
		data: item,
		next: stack.top,
		min:  newMin,
	}
	stack.top = &newNode
}

// Golang has no builtin "min" function...
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (stack MinStack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return stack.top.data, nil
}

func (stack MinStack) IsEmpty() bool {
	return stack.top == nil
}

func (stack MinStack) Min() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return stack.top.min, nil
}

// A setOfStacks provides the interface of a single stack while actually
// being composed of a dynamic number of stacks which are each of a fixed
// size. Since the size of each internal stack is fixed, it is simplest
// to implement using arrays.
type setOfStacks struct {
	stacks    [][]int // internal stacks implemented as arrays
	stackSize int     // fixed size of each internal stack

	// Indices that specify the current top stack and position within that top
	// stack respecively. -1 is used as a special value for the topStack to
	// indicate an empty stack.
	topStack int
	topIx    int
}

// Construct a new set of stacks. Note that no arrays are allocated yet -
// the array backing the first internal stack will be allocated when the
// first value is pushed.
func NewSetOfStacks(stackSize int) setOfStacks {
	return setOfStacks{
		stacks:    nil,
		stackSize: stackSize,
		topStack:  -1,
		topIx:     stackSize - 1,
	}
}

// Pop an item from the stack. Returns an error if the stack is empty. If
// the item popped was the last one in an internal stack, move the top indices
// to point at the previous stack.
func (stacks *setOfStacks) Pop() (int, error) {
	if stacks.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	retVal := stacks.stacks[stacks.topStack][stacks.topIx]

	if stacks.topIx > 0 {
		stacks.topIx--
	} else {
		// Last item in an internal stack is being popped. Reduce the number
		// of internal stacks by one.
		stacks.stacks[stacks.topStack] = nil
		stacks.stacks = stacks.stacks[:stacks.topStack]
		stacks.topStack--
		stacks.topIx = stacks.stackSize - 1
	}

	return retVal, nil
}

// Push an item onto the stack. If the current internal stack is full, allocate
// a new one and update internal indices to point at the new top stack.
func (stacks *setOfStacks) Push(item int) {
	if stacks.topIx < stacks.stackSize-1 {
		stacks.topIx++
		stacks.stacks[stacks.topStack][stacks.topIx] = item
	} else {
		// Create a new internal stack.
		newStack := make([]int, stacks.stackSize)
		stacks.stacks = append(stacks.stacks, newStack)
		stacks.topStack++
		stacks.topIx = 0
		newStack[0] = item
	}
}

// Peek at the top value in the stack without removing it.
func (stacks *setOfStacks) Peek() (int, error) {
	if stacks.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	return stacks.stacks[stacks.topStack][stacks.topIx], nil
}

// Check if a stack is empty.
func (stacks *setOfStacks) IsEmpty() bool {
	return stacks.topStack == -1
}

// The MyQueue type implements the Queue interface but is actually implemented
// using two stacks.
type myQueue struct {
	inStack  *basicStack
	outStack *basicStack
}

// Construct a new myQueue.
func NewMyQueue() *myQueue {
	return &myQueue{inStack: NewBasicStack(), outStack: NewBasicStack()}
}

// Pop an item from the front of the queue.
func (queue *myQueue) Pop() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	if !queue.inStack.IsEmpty() {
		queue.emptyInStack()
	}

	return queue.outStack.Pop()
}

// Push an item onto the back of the queue.
func (queue *myQueue) Push(item int) {
	if !queue.outStack.IsEmpty() {
		queue.emptyOutStack()
	}

	queue.inStack.Push(item)
}

// Peek at the front of the queue.
func (queue myQueue) Peek() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	if !queue.inStack.IsEmpty() {
		queue.emptyInStack()
	}

	return queue.outStack.Peek()
}

// Check if the queue is empty.
func (queue myQueue) IsEmpty() bool {
	return queue.inStack.IsEmpty() && queue.outStack.IsEmpty()
}

// Empty the inStack into the outStack.
func (queue *myQueue) emptyInStack() error {
	if !queue.outStack.IsEmpty() {
		return errors.New("outStack is not empty")
	}

	for queue.inStack.IsEmpty() {
		val, err := queue.inStack.Pop()
		if err != nil {
			return err
		}

		queue.outStack.Push(val)
	}

	return nil
}

// Empty the outStack into the inStack.
func (queue *myQueue) emptyOutStack() error {
	if !queue.inStack.IsEmpty() {
		return errors.New("inStack is not empty")
	}

	for queue.outStack.IsEmpty() {
		val, err := queue.outStack.Pop()
		if err != nil {
			return err
		}

		queue.inStack.Push(val)
	}

	return nil
}
