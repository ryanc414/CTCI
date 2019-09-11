package ctci

import "errors"

// Stores three stacks interleaved in a single array.
type threeStacks struct {
	stackData    []int
	stackIndices [3]int
}

// Construct a new threeStacks object.
func NewThreeStacks() threeStacks {
	stackData := make([]int, 3)
	return threeStacks{stackData: stackData, stackIndices: [3]int{-1, -1, -1}}
}

// Pops an item off the top of a threeStacks stack. An error is returned if the
// stack is empty.
func (stacks *threeStacks) Pop(stackIx int) (int, error) {
	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.stackIndices[stackIx]*3 + stackIx
	stacks.stackIndices[stackIx]--

	return stacks.stackData[dataIx], nil
}

// Push an item onto the top of a threeStacks stack. The backing array is
// automatically expanded if required, so this function should always succeed.
func (stacks *threeStacks) Push(stackIx, item int) {
	stacks.stackIndices[stackIx]++
	newDataIx := stacks.stackIndices[stackIx]*3 + stackIx

	// Expand data slice if required.
	if newDataIx >= len(stacks.stackData) {
		extraReqd := newDataIx + 1 - len(stacks.stackData)
		stacks.stackData = append(stacks.stackData, make([]int, extraReqd)...)
	}

	stacks.stackData[newDataIx] = item
}

// Return the value at the top of a threeStacks stack without removing it. An
// error is returned if the stack is empty.
func (stacks *threeStacks) Peek(stackIx int) (int, error) {
	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.stackIndices[stackIx]*3 + stackIx
	return stacks.stackData[dataIx], nil
}

// Check if a threeStacks stack is empty.
func (stacks *threeStacks) IsEmpty(stackIx int) bool {
	return stacks.stackIndices[stackIx] == -1
}
