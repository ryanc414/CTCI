package ctci

import "errors"

// Stores three stacks interleaved in a single array.
type multiStacks struct {
	stackData    []int
	stackIndices []int
    numStacks int
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
        stackData: stackData,
        stackIndices: stackIndices,
        numStacks: numStacks,
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
    return (stacks.stackIndices[stackIx]*stacks.numStacks) + stackIx
}
