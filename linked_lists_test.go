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
	t.Log(list.String())

	if !list.Equals(&expected) {
		t.Error()
	}

	var empty *Node
	empty.RemoveDups()
	if empty != nil {
		t.Error()
	}
}

// Test the RemoveDupsNoBuf method.
func TestRemoveDupsNoBuf(t *testing.T) {
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

	list.RemoveDupsNoBuf()
	t.Log(list.String())

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

	var empty *Node
	empty.RemoveDupsNoBuf()
	if empty != nil {
		t.Error()
	}
}

// Test the KthToLast function.
func TestKthToLast(t *testing.T) {
	list := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: &Node{
					data: 4, next: &Node{
						data: 5, next: &Node{
							data: 6, next: nil,
						},
					},
				},
			},
		},
	}

	if list.KthToLast(0).data != 6 {
		t.Error()
	}

	if list.KthToLast(3).data != 3 {
		t.Error()
	}

	if list.KthToLast(5).data != 1 {
		t.Error()
	}

	if list.KthToLast(6) != nil {
		t.Error()
	}

	var empty *Node
	if empty.KthToLast(0) != nil {
		t.Error()
	}
}

// Test the DeleteMiddle function.
func TestDeleteMiddle(t *testing.T) {
	list := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: &Node{
					data: 4, next: &Node{
						data: 5, next: &Node{
							data: 6, next: nil,
						},
					},
				},
			},
		},
	}

	list.next.next.DeleteMiddle()      // deletes 3
	list.next.next.next.DeleteMiddle() // deletes 5

	expected := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 4, next: &Node{
					data: 6, next: nil,
				},
			},
		},
	}

	if !list.Equals(&expected) {
		t.Error()
	}
}

// Test the Partition function.
func TestPartition(t *testing.T) {
	list := Node{
		data: 3, next: &Node{
			data: 5, next: &Node{
				data: 8, next: &Node{
					data: 5, next: &Node{
						data: 10, next: &Node{
							data: 2, next: &Node{
								data: 1, next: nil,
							},
						},
					},
				},
			},
		},
	}

	newHead := list.Partition(5)

	expected := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: &Node{
					data: 5, next: &Node{
						data: 8, next: &Node{
							data: 5, next: &Node{
								data: 10, next: nil,
							},
						},
					},
				},
			},
		},
	}

	if !newHead.Equals(&expected) {
		t.Error()
	}
}
