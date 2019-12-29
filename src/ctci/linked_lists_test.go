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

	t.Log(newHead.String())
	if !newHead.Equals(&expected) {
		t.Error()
	}
}

// Test the SumLists function.
func TestSumLists(t *testing.T) {
	// Try summing two lists of equal length.
	listA := Node{
		data: 7, next: &Node{
			data: 1, next: &Node{
				data: 6, next: nil,
			},
		},
	}
	listB := Node{
		data: 5, next: &Node{
			data: 9, next: &Node{
				data: 2, next: nil,
			},
		},
	}
	expectedSum := Node{
		data: 2, next: &Node{
			data: 1, next: &Node{
				data: 9, next: nil,
			},
		},
	}

	sum := listA.SumLists(&listB)

	t.Log(sum.String())
	if !sum.Equals(&expectedSum) {
		t.Error()
	}

	// Try summing an empty list.
	var empty *Node
	if empty.SumLists(nil) != nil {
		t.Error()
	}

	// Try summing two lists of different lengths.
	listC := Node{
		data: 2, next: &Node{
			data: 9, next: &Node{
				data: 3, next: &Node{
					data: 8, next: nil,
				},
			},
		},
	}
	expectedSum = Node{
		data: 9, next: &Node{
			data: 0, next: &Node{
				data: 0, next: &Node{
					data: 9, next: nil,
				},
			},
		},
	}

	sum = listA.SumLists(&listC)

	t.Log(sum.String())
	if !sum.Equals(&expectedSum) {
		t.Error()
	}
}

// Test the SumListsForward function.
func TestSumListsForward(t *testing.T) {
	// Try summing two lists of equal length.
	listA := Node{
		data: 7, next: &Node{
			data: 1, next: &Node{
				data: 6, next: nil,
			},
		},
	}
	listB := Node{
		data: 5, next: &Node{
			data: 9, next: &Node{
				data: 2, next: nil,
			},
		},
	}
	expectedSum := Node{
		data: 1, next: &Node{
			data: 3, next: &Node{
				data: 0, next: &Node{
					data: 8, next: nil,
				},
			},
		},
	}

	sum := listA.SumListsForward(&listB)

	t.Log(sum.String())
	if !sum.Equals(&expectedSum) {
		t.Error()
	}

	// Try summing an empty list.
	var empty *Node
	if empty.SumListsForward(nil) != nil {
		t.Error()
	}

	// Try summing two lists of different lengths.
	listC := Node{
		data: 2, next: &Node{
			data: 9, next: &Node{
				data: 3, next: &Node{
					data: 8, next: nil,
				},
			},
		},
	}
	expectedSum = Node{
		data: 3, next: &Node{
			data: 6, next: &Node{
				data: 5, next: &Node{
					data: 4, next: nil,
				},
			},
		},
	}

	sum = listA.SumListsForward(&listC)

	t.Log(sum.String())
	if !sum.Equals(&expectedSum) {
		t.Error()
	}
}

// Test the IsPalindrome function.
func TestIsPalindrome(t *testing.T) {
	listA := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: &Node{
					data: 2, next: &Node{
						data: 1, next: nil,
					},
				},
			},
		},
	}
	listB := Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: &Node{
					data: 2, next: &Node{
						data: 2, next: nil,
					},
				},
			},
		},
	}

	var empty *Node

	if !listA.IsPalindrome() {
		t.Error()
	}
	if listB.IsPalindrome() {
		t.Error()
	}
	if !empty.IsPalindrome() {
		t.Error()
	}
}

// Test the FindIntersection function.
func TestFindIntersection(t *testing.T) {
	commonTail := &Node{
		data: 4, next: &Node{
			data: 5, next: &Node{
				data: 6, next: nil,
			},
		},
	}
	listA := &Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: commonTail,
			},
		},
	}
	listB := &Node{
		data: 6, next: &Node{
			data: 5, next: commonTail,
		},
	}
	listC := &Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: nil,
			},
		},
	}

	if FindIntersection(listA, listB) != listB.next.next {
		t.Error()
	}

	if FindIntersection(listA, listC) != nil {
		t.Error()
	}

	var empty *Node
	if FindIntersection(listA, empty) != nil {
		t.Error()
	}
}

// Test the FindLoop() function.
func TestFindLoop(t *testing.T) {
	loop := &Node{
		data: 3, next: &Node{
			data: 4, next: &Node{
				data: 5, next: nil,
			},
		},
	}
	loop.next.next.next = loop

	hasLoop := &Node{
		data: 1, next: &Node{
			data: 2, next: loop,
		},
	}

	if hasLoop.FindLoop() != loop {
		t.Error()
	}

	noLoop := &Node{
		data: 1, next: &Node{
			data: 2, next: &Node{
				data: 3, next: nil,
			},
		},
	}

	if noLoop.FindLoop() != nil {
		t.Error()
	}

	var empty *Node
	if empty.FindLoop() != nil {
		t.Error()
	}
}
