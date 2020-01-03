package ctci

import "testing"

func TestCircularArray(t *testing.T) {
	data := []interface{}{1, 2, 3, 4, 5, 6}
	circArray := InitCircularArray(data)
	checkCircOrder(t, circArray, data)

	newCirc := circArray.Rotate(3)
	expectedOrder := []interface{}{4, 5, 6, 1, 2, 3}
	checkCircOrder(t, newCirc, expectedOrder)
}

func checkCircOrder(
	t *testing.T, circArray CircularArray, expected []interface{},
) {
	var actual []interface{}
	for i := circArray.First(); i != EndOfArray; i = circArray.Next(i) {
		actual = append(actual, circArray.Data[i])
	}

	if !equalSlices(actual, expected) {
		t.Error(actual)
	}
}

func equalSlices(actual, expected []interface{}) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i := range actual {
		if actual[i].(int) != expected[i].(int) {
			return false
		}
	}

	return true
}
