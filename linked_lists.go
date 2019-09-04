package ctci

import (
    "fmt"
    "strings"
)

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

// Remove duplicates from a linked-list without using any additional buffer.
func (list *Node) RemoveDupsNoBuf() {
    for currNode := list; currNode != nil; currNode = currNode.next {
        currNode.RemoveAllAfter(currNode.data)
    }
}

// Remove all nodes after the current node that contain a particular value.
func (list *Node) RemoveAllAfter(removeData int) {
    if list == nil {
        return
    }

    for currNode := list; currNode.next != nil; {
        if currNode.next.data == removeData {
            // Delete the next node.
            currNode.next = currNode.next.next
        } else {
            currNode = currNode.next
        }
    }
}

// Format all nodes in the list as a string.
func (list *Node) String() string {
    var b strings.Builder

    for currNode := list; currNode != nil; currNode = currNode.next {
        fmt.Fprintf(&b, "%d", currNode.data)
        if currNode.next != nil {
            b.WriteString(" -> ")
        }
    }

    return b.String()
}
