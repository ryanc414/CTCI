package lists

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

// Compare two lists.
func (list *Node) Equals(other *Node) bool {
	listNode := list
	otherNode := other
	for listNode != nil && otherNode != nil {
		if listNode.data != otherNode.data {
			return false
		}
		listNode = listNode.next
		otherNode = otherNode.next
	}

	return listNode == nil && otherNode == nil
}

// Return the Kth-to-last element in a singly linked list.
func (list *Node) KthToLast(k int) *Node {
	_, foundNode := list.kthToLastRecur(k)
	return foundNode
}

// Recursively check each node in a list to find the Kth to last element.
// Returns the number of list nodes following this one, and a node if found.
// Returns nil if no node is found, because k exceeds the number of nodes in
// this section.
func (node *Node) kthToLastRecur(k int) (int, *Node) {
	if node == nil {
		return 0, nil
	}

	tailCount, foundNode := node.next.kthToLastRecur(k)

	if foundNode != nil {
		return tailCount + 1, foundNode
	}

	if tailCount == k {
		return tailCount + 1, node
	}

	return tailCount + 1, nil
}

// Delete a node from the middle of a singly linked list - i.e. any node
// apart from the first or last.
func (node *Node) DeleteMiddle() {
	if node == nil || node.next == nil {
		panic("Node is not in the middle of a list.")
	}

	// Update this node to match the next one's value and next pointer. The
	// next node will be GC'd so we don't need to explicitly deallocate it.
	node.data = node.next.data
	node.next = node.next.next
}

// Partition a list around a value x, such that all nodes with values less
// than x come before all nodes with values greater than or equal to x.
func (listHead *Node) Partition(x int) *Node {
	if listHead == nil {
		return listHead
	}

	currNode := listHead
	for currNode.next != nil {
		if currNode.next.data < x {
			// Delete the next node from the list and insert before the front.
			insertNode := currNode.next
			currNode.next = insertNode.next
			insertNode.next = listHead
			listHead = insertNode
		} else {
			currNode = currNode.next
		}
	}

	return listHead
}

// Add two integers represented by a list of digits in reverse order.
func (list *Node) SumLists(other *Node) *Node {
	return sumListRecur(list, other, 0)
}

// Recursive implementation.
func sumListRecur(list, other *Node, carry int) *Node {
	if list == nil && other == nil {
		if carry > 0 {
			return &Node{data: carry, next: nil}
		} else {
			return nil
		}
	}

	var nextValue int
	var nextList *Node
	var nextOther *Node
	if list == nil {
		nextValue = other.data + carry
		nextList = nil
		nextOther = other.next
	} else if other == nil {
		nextValue = list.data + carry
		nextList = list.next
		nextOther = nil
	} else {
		nextValue = list.data + other.data + carry
		nextList = list.next
		nextOther = other.next
	}

	var nextCarry int
	if nextValue > 9 {
		nextValue -= 10
		nextCarry = 1
	}

	return &Node{
		data: nextValue,
		next: sumListRecur(nextList, nextOther, nextCarry),
	}
}

// Add two integers represented by a list of digits in forward order.
func (list *Node) SumListsForward(other *Node) *Node {
	listLength := list.Length()
	otherLength := other.Length()

	if listLength > otherLength {
		other = other.prePadZeros(listLength - otherLength)
	} else if otherLength > listLength {
		list = list.prePadZeros(otherLength - listLength)
	}

	result, carry := sumListFwdRecur(list, other)

	if carry > 0 {
		return &Node{data: carry, next: result}
	} else {
		return result
	}
}

// Return the length of a linked list, by iterating to the end.
func (list *Node) Length() int {
	length := 0
	for currNode := list; currNode != nil; currNode = currNode.next {
		length++
	}

	return length
}

// Pre-pad a list with a given number of nodes containing data of 0.
func (list *Node) prePadZeros(numZeros int) *Node {
	if numZeros < 0 {
		panic("numZeros must be a positive integer")
	}

	for i := 0; i < numZeros; i++ {
		list = &Node{data: 0, next: list}
	}

	return list
}

// Recursive implementation for summing two integers represented by lists
// of digits in forward order. Both lists must be the same length to ensure
// digits for the same power of 10 are processed together.
func sumListFwdRecur(list, other *Node) (*Node, int) {
	// Terminate recursion at end of both lists.
	if list == nil && other == nil {
		return nil, 0
	}

	// Perform recursive step first to sum the tails of both lists.
	tailSum, carry := sumListFwdRecur(list.next, other.next)

	nextValue := list.data + other.data + carry
	nextCarry := 0

	if nextValue > 9 {
		nextValue -= 10
		nextCarry = 1
	}

	return &Node{data: nextValue, next: tailSum}, nextCarry
}

// Check if a singly-linked list is a palindrome: the same forwards or
// backwards
func (list *Node) IsPalindrome() bool {
	_, isPalindrome := isPalindromeRecur(list, list)
	return isPalindrome
}

func isPalindromeRecur(head, currNode *Node) (*Node, bool) {
	if currNode == nil {
		return head, true
	}

	mirrorNode, isPalindrome := isPalindromeRecur(head, currNode.next)
	if !isPalindrome {
		return nil, false
	}

	return mirrorNode.next, mirrorNode.data == currNode.data
}

// Find the intersection point of two singly linked lists. If the lists do not
// intersect, nil is returned.
func FindIntersection(listA, listB *Node) *Node {
	lengthA := listA.Length()
	lengthB := listB.Length()

	// Shorten the longer list so that both lists are the same length.
	if lengthA > lengthB {
		for i := 0; i < lengthA-lengthB; i++ {
			listA = listA.next
		}
	} else if lengthB > lengthA {
		for i := 0; i < lengthB-lengthA; i++ {
			listB = listB.next
		}
	}

	// Now iterate through both lists until we find an intersection point,
	// or until both list pointers are nil.
	for listA != listB {
		listA = listA.next
		listB = listB.next
	}

	return listA
}

// If a loop exists in a list, find and return the node at the start of the
// loop.
func (list *Node) FindLoop() *Node {
	visitedNodes := make(map[*Node]bool)

	for currNode := list; currNode != nil; currNode = currNode.next {
		if visitedNodes[currNode] {
			return currNode
		}
		visitedNodes[currNode] = true
	}

	// No loop was found so return nil.
	return nil
}
