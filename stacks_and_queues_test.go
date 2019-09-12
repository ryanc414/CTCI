package ctci

import "testing"

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
