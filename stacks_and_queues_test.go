package ctci

import "testing"

// Test the threeStacks type and methods.
func TestThreeStacks(t *testing.T) {
	stacks := NewThreeStacks()

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
