package ctci

import "testing"

// Test calculating the number of step ways recursively and iteratively.
func TestNumStepWays(t *testing.T) {
	res := NumStepWays(6)
	if res != 24 {
		t.Error(res)
	}

	res = NumStepWaysIter(6)
	if res != 24 {
		t.Error(res)
	}
}
