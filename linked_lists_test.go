package ctci

import "testing"

// Test the RemoveDups method.
func TestRemoveDups(t *testing.T) {
	list := Node{
		data: 1, next: &Node{
			data: 1, next: &Node{
				data: 2, next: &Node{
					data: 2, next: &Node{
						data: 3, next: &Node{
							data: 3, next: nil,
						},
					},
				},
			},
		},
	}

	expected := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: nil,
			},
		},
	}

	list.RemoveDups()

	listNode := &list
	expectedNode := &expected
	for listNode != nil && expectedNode != nil {
		if listNode.data != expectedNode.data {
			t.Error()
		}
		listNode = listNode.next
		expectedNode = expectedNode.next
	}

	if listNode != nil || expectedNode != nil {
		t.Error()
	}

	var empty Node
	empty.RemoveDups()
}
