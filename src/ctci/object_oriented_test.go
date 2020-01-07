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

// Test setting, getting and deleting values in the hash table.
func TestHashTable(t *testing.T) {
	table := InitHashTable(10)

	// Try to get value from empty table.
	_, err := table.GetItem(9)
	if err == nil {
		t.Error(err)
	}

	// Set a key and check it can be retrieved.
	table.SetItem(25, "Hello, world")

	val, err := table.GetItem(25)
	checkValAndErr(t, "Hello, world", val.(string), err)

	// Set a key with the same hash and check it can still be retrieved.
	table.SetItem(55, "Another string")
	val, err = table.GetItem(55)
	checkValAndErr(t, "Another string", val.(string), err)

	// Try updating an existing key.
	table.SetItem(25, "and it's gone")
	val, err = table.GetItem(25)
	checkValAndErr(t, "and it's gone", val.(string), err)

	// Finally, delete the key and check it cannot be retrieved.
	err = table.Delete(25)
	if err != nil {
		t.Error(err)
	}

	_, err = table.GetItem(25)
	if err == nil {
		t.Error()
	}
}

// Check there is no error and the value matches what is expected.
func checkValAndErr(t *testing.T, expected, actual string, err error) {
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Error(actual)
	}
}
