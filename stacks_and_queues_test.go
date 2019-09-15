package ctci

import "testing"

// Test the BasicStack type and methods.
func TestBasicStack(t *testing.T) {
	stack := NewBasicStack()
	genericStackTest(t, stack)
}

// Test the multiStacks type and methods.
func TestMultiStacks(t *testing.T) {
	stacks := NewMultiStacks(3)

	for i := 0; i < 3; i++ {
		if !stacks.IsEmpty(i) {
			t.Error()
		}

		_, err := stacks.Pop(i)
		if err == nil {
			t.Error()
		}

		_, err = stacks.Peek(i)
		if err == nil {
			t.Error()
		}

		stacks.Push(i, 1)
		stacks.Push(i, 2)
		stacks.Push(i, 3)

		if stacks.IsEmpty(i) {
			t.Error()
		}

		val, err := stacks.Pop(i)
		if err != nil || val != 3 {
			t.Error()
		}

		val, err = stacks.Peek(i)
		if err != nil || val != 2 {
			t.Error()
		}
	}
}

// Test the minStacks type and methods.
func TestMinStack(t *testing.T) {
	var stack MinStack
	genericStackTest(t, &stack)

	if !stack.IsEmpty() {
		t.Error()
	}

	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Min()
	if val != 1 || err != nil {
		t.Error()
	}

	val, err = stack.Pop()
	if val != 1 || err != nil {
		t.Error()
	}

	val, err = stack.Min()
	if val != 2 || err != nil {
		t.Error()
	}

	stack.Push(0)
	val, err = stack.Peek()
	if val != 0 || err != nil {
		t.Error()
	}

	val, err = stack.Min()
	if val != 0 || err != nil {
		t.Error()
	}
}

// Test the setOfStacks type and methods.
func TestSetOfStacks(t *testing.T) {
	stackSet := NewSetOfStacks(3)
	genericStackTest(t, &stackSet)

	if !stackSet.IsEmpty() {
		t.Error()
	}

	// Push 3 values and peek at the top value. This will fill the first
	// internal stack.
	stackSet.Push(1)
	stackSet.Push(2)
	stackSet.Push(3)

	val, err := stackSet.Peek()
	if val != 3 || err != nil {
		t.Error()
	}

	// Push another 3 values and peek again. This will fill the second internal
	// stack.
	stackSet.Push(4)
	stackSet.Push(5)
	stackSet.Push(6)

	val, err = stackSet.Peek()
	if val != 6 || err != nil {
		t.Error()
	}

	// Pop all values.
	for expectedVal := 6; expectedVal > 0; expectedVal-- {
		val, err = stackSet.Pop()
		if val != expectedVal || err != nil {
			t.Error()
		}
	}

	if !stackSet.IsEmpty() {
		t.Error()
	}
}

// Generic test for any type that implements the stack interface.
func genericStackTest(t *testing.T, stack IntStack) {
	if !stack.IsEmpty() {
		t.Error()
	}

	_, err := stack.Pop()
	if err == nil {
		t.Error()
	}

	_, err = stack.Peek()
	if err == nil {
		t.Error()
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, err := stack.Pop()
	if val != 3 || err != nil {
		t.Error()
	}

	val, err = stack.Peek()
	if val != 2 || err != nil {
		t.Error()
	}

	// Pop the rest of the values, so that we leave the stack empty.
	val, err = stack.Pop()
	if val != 2 || err != nil {
		t.Error()
	}

	val, err = stack.Pop()
	if val != 1 || err != nil {
		t.Error()
	}

	if !stack.IsEmpty() {
		t.Error()
	}
}

// Test the MyQueue type and methods.
func TestMyQueue(t *testing.T) {
	queue := NewMyQueue()

	if !queue.IsEmpty() {
		t.Error()
	}

	_, err := queue.Pop()
	if err == nil {
		t.Error()
	}

	_, err = queue.Peek()
	if err == nil {
		t.Error()
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.IsEmpty() {
		t.Error()
	}

	// Since the queue has LIFO ordering, the first value to come out should
	// be 1.
	val, err := queue.Peek()
	if val != 1 || err != nil {
		t.Error()
	}

	// Check that the values are popped in order 1, 2, 3
	for expected := 1; expected < 4; expected++ {
		val, err = queue.Pop()
		if val != expected || err != nil {
			t.Error()
		}
	}

	// All values have been popped - queue should again be empty.
	if !queue.IsEmpty() {
		t.Error()
	}
}
