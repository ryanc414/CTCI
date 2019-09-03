package ctci

type Node struct {
	next *Node
	data int
}

// Remove duplicates from a linked-list in place.
func (list *Node) RemoveDups() {
	if list == nil {
		return
	}

	valsFound := make(map[int]bool)
	valsFound[list.data] = true

	currNode := list

	for currNode.next != nil {
		if valsFound[currNode.next.data] {
			// Delete next node
			currNode.next = currNode.next.next
		} else {
			valsFound[currNode.next.data] = true
			currNode = currNode.next
		}
	}
}
