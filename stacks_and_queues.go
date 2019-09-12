package ctci

import "errors"

// A generic stack must implement four methods: Pop, Push, Peek and IsEmpty.
type interface IntStack {
    Pop() (int, error)
    Push(item int)
    Peek() (int, error)
    IsEmpty() bool
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
