package ctci

import "testing"

// Test the CircularArray type and methods.
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

// Test the file tree types and methods.
func TestFileTree(t *testing.T) {
	root := CreateRootDir("root")

	_, err := root.CreateFile("file1", []byte("text goes here"))
	if err != nil {
		t.Error(err)
	}

	_, err = root.CreateDir("subdir1")
	if err != nil {
		t.Error(err)
	}

	expectedContents := []string{"subdir1", "file1"}
	actualContents := root.List()

	if len(actualContents) != len(expectedContents) {
		t.Error(actualContents)
		return
	}

	for i := range actualContents {
		if actualContents[i] != expectedContents[i] {
			t.Error(actualContents)
			return
		}
	}
}
