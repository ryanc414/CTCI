package stacks

import (
	"container/list"
	"errors"
)

// A generic stack must implement four methods: Pop, Push, Peek and IsEmpty.
type Stack interface {
	Pop() (interface{}, error)
	Push(item interface{})
	Peek() (interface{}, error)
	IsEmpty() bool
}

// A generic queue must implement the same four methods as a stack. Only the
// ordering of popped elements differs - queues are popped in LIFO order
// while stacks are FIFO.
type Queue interface {
	Add(item interface{})
	Remove() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}

// Basic implementation of a stack, using a dynamic expanding array slice.
type BasicStack struct {
	Data []interface{}
}

// Construct a new empty stack.
func NewBasicStack() *BasicStack {
	return &BasicStack{}
}

// Pop an element from the stack - pops last element from the slice.
func (stack *BasicStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	newLen := len(stack.Data) - 1
	var popped interface{}
	popped, stack.Data = stack.Data[newLen], stack.Data[:newLen]

	return popped, nil
}

// Push an element onto the stack - item is appended to the slice.
func (stack *BasicStack) Push(item interface{}) {
	stack.Data = append(stack.Data, item)
}

// Peek at the top element in the stack.
func (stack BasicStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	return stack.Data[len(stack.Data)-1], nil
}

// Check if the stack is empty, by checking if the data slice has a non-zero
// length.
func (stack BasicStack) IsEmpty() bool {
	return len(stack.Data) == 0
}

// Implement a basic queue using a linked list.
type BasicQueue struct {
	list list.List
}

func NewBasicQueue() *BasicQueue {
	return &BasicQueue{}
}

func (queue *BasicQueue) Add(item interface{}) {
	queue.list.PushBack(item)
}

func (queue *BasicQueue) Remove() (interface{}, error) {
	frontEl := queue.list.Front()
	if frontEl == nil {
		return nil, errors.New("Queue is empty")
	}

	retVal := frontEl.Value
	queue.list.Remove(frontEl)
	return retVal, nil
}

func (queue BasicQueue) Peek() (interface{}, error) {
	frontEl := queue.list.Front()
	if frontEl == nil {
		return nil, errors.New("Queue is empty")
	}

	return frontEl.Value, nil
}

func (queue BasicQueue) IsEmpty() bool {
	return queue.list.Front() == nil
}

// Stores three stacks interleaved in a single array.
type multiStacks struct {
	stackData    []int
	stackIndices []int
	numStacks    int
}

// Construct a new multiStacks object.
func NewMultiStacks(numStacks int) multiStacks {
	if numStacks < 1 {
		panic("numStacks must be >0")
	}

	stackData := make([]int, numStacks)
	stackIndices := make([]int, numStacks)
	for i := range stackIndices {
		stackIndices[i] = -1
	}

	return multiStacks{
		stackData:    stackData,
		stackIndices: stackIndices,
		numStacks:    numStacks,
	}
}

// Pops an item off the top of a multiStacks stack. An error is returned if the
// stack is empty.
func (stacks *multiStacks) Pop(stackIx int) (int, error) {
	stacks.checkStackIx(stackIx)

	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.getDataIx(stackIx)
	stacks.stackIndices[stackIx]--

	return stacks.stackData[dataIx], nil
}

// Push an item onto the top of a multiStacks stack. The backing array is
// automatically expanded if required, so this function should always succeed.
func (stacks *multiStacks) Push(stackIx, item int) {
	stacks.checkStackIx(stackIx)

	stacks.stackIndices[stackIx]++
	newDataIx := stacks.getDataIx(stackIx)

	// Expand data slice if required.
	if newDataIx >= len(stacks.stackData) {
		extraReqd := newDataIx + 1 - len(stacks.stackData)
		stacks.stackData = append(stacks.stackData, make([]int, extraReqd)...)
	}

	stacks.stackData[newDataIx] = item
}

// Return the value at the top of a multiStacks stack without removing it. An
// error is returned if the stack is empty.
func (stacks *multiStacks) Peek(stackIx int) (int, error) {
	stacks.checkStackIx(stackIx)

	if stacks.stackIndices[stackIx] == -1 {
		return 0, errors.New("Stack is empty")
	}

	dataIx := stacks.getDataIx(stackIx)
	return stacks.stackData[dataIx], nil
}

// Check if a multiStacks stack is empty.
func (stacks *multiStacks) IsEmpty(stackIx int) bool {
	stacks.checkStackIx(stackIx)
	return stacks.stackIndices[stackIx] == -1
}

// Check if the stackIx is within expected range.
func (stacks multiStacks) checkStackIx(stackIx int) {
	if stackIx < 0 || stackIx > stacks.numStacks {
		panic("stackIx is not within expected range")
	}
}

// Return the current index for a stack within the shared stack data array.
func (stacks multiStacks) getDataIx(stackIx int) int {
	return (stacks.stackIndices[stackIx] * stacks.numStacks) + stackIx
}

// A MinStack and its node. This implements a stack interface with the addition
// of providing O(1) access to the minimum value in the stack. This is achieved
// by storing an extra min value on each stack node, that keeps track of the
// minimal value from that node downwards in the stack.
type MinStackNode struct {
	data int
	next *MinStackNode
	min  int
}

type MinStack struct {
	top *MinStackNode
}

func (stack *MinStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	retVal := stack.top.data
	stack.top = stack.top.next

	return retVal, nil
}

func (stack *MinStack) Push(item interface{}) {
	var newMin int
	if stack.IsEmpty() {
		newMin = item.(int)
	} else {
		newMin = min(item.(int), stack.top.min)
	}

	newNode := MinStackNode{
		data: item.(int),
		next: stack.top,
		min:  newMin,
	}
	stack.top = &newNode
}

// Golang has no builtin "min" function...
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (stack MinStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return stack.top.data, nil
}

func (stack MinStack) IsEmpty() bool {
	return stack.top == nil
}

func (stack MinStack) Min() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return stack.top.min, nil
}

// A setOfStacks provides the interface of a single stack while actually
// being composed of a dynamic number of stacks which are each of a fixed
// size. Since the size of each internal stack is fixed, it is simplest
// to implement using arrays.
type setOfStacks struct {
	stacks    [][]interface{} // internal stacks implemented as arrays
	stackSize int             // fixed size of each internal stack

	// Indices that specify the current top stack and position within that top
	// stack respecively. -1 is used as a special value for the topStack to
	// indicate an empty stack.
	topStack int
	topIx    int
}

// Construct a new set of stacks. Note that no arrays are allocated yet -
// the array backing the first internal stack will be allocated when the
// first value is pushed.
func NewSetOfStacks(stackSize int) setOfStacks {
	return setOfStacks{
		stacks:    nil,
		stackSize: stackSize,
		topStack:  -1,
		topIx:     stackSize - 1,
	}
}

// Pop an item from the stack. Returns an error if the stack is empty. If
// the item popped was the last one in an internal stack, move the top indices
// to point at the previous stack.
func (stacks *setOfStacks) Pop() (interface{}, error) {
	if stacks.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	retVal := stacks.stacks[stacks.topStack][stacks.topIx]

	if stacks.topIx > 0 {
		stacks.topIx--
	} else {
		// Last item in an internal stack is being popped. Reduce the number
		// of internal stacks by one.
		stacks.stacks[stacks.topStack] = nil
		stacks.stacks = stacks.stacks[:stacks.topStack]
		stacks.topStack--
		stacks.topIx = stacks.stackSize - 1
	}

	return retVal, nil
}

// Push an item onto the stack. If the current internal stack is full, allocate
// a new one and update internal indices to point at the new top stack.
func (stacks *setOfStacks) Push(item interface{}) {
	if stacks.topIx < stacks.stackSize-1 {
		stacks.topIx++
		stacks.stacks[stacks.topStack][stacks.topIx] = item
	} else {
		// Create a new internal stack.
		newStack := make([]interface{}, stacks.stackSize)
		stacks.stacks = append(stacks.stacks, newStack)
		stacks.topStack++
		stacks.topIx = 0
		newStack[0] = item
	}
}

// Peek at the top value in the stack without removing it.
func (stacks *setOfStacks) Peek() (interface{}, error) {
	if stacks.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	return stacks.stacks[stacks.topStack][stacks.topIx], nil
}

// Check if a stack is empty.
func (stacks *setOfStacks) IsEmpty() bool {
	return stacks.topStack == -1
}

// The MyQueue type implements the Queue interface but is actually implemented
// using two stacks.
type myQueue struct {
	inStack  BasicStack
	outStack BasicStack
}

// Construct a new myQueue.
func NewMyQueue() *myQueue {
	return &myQueue{}
}

// Remove an item from the front of the queue.
func (queue *myQueue) Remove() (interface{}, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	if !queue.inStack.IsEmpty() {
		err := queue.emptyInStack()
		if err != nil {
			panic(err)
		}
	}

	return queue.outStack.Pop()
}

// Add an item onto the back of the queue.
func (queue *myQueue) Add(item interface{}) {
	if !queue.outStack.IsEmpty() {
		err := queue.emptyOutStack()
		if err != nil {
			panic(err)
		}
	}

	queue.inStack.Push(item)
}

// Peek at the front of the queue.
func (queue *myQueue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	if !queue.inStack.IsEmpty() {
		err := queue.emptyInStack()
		if err != nil {
			panic(err)
		}
	}

	return queue.outStack.Peek()
}

// Check if the queue is empty.
func (queue myQueue) IsEmpty() bool {
	return queue.inStack.IsEmpty() && queue.outStack.IsEmpty()
}

// Empty the inStack into the outStack.
func (queue *myQueue) emptyInStack() error {
	if !queue.outStack.IsEmpty() {
		return errors.New("outStack is not empty")
	}

	for !queue.inStack.IsEmpty() {
		val, err := queue.inStack.Pop()
		if err != nil {
			return err
		}

		queue.outStack.Push(val)
	}

	return nil
}

// Empty the outStack into the inStack.
func (queue *myQueue) emptyOutStack() error {
	if !queue.inStack.IsEmpty() {
		return errors.New("inStack is not empty")
	}

	for !queue.outStack.IsEmpty() {
		val, err := queue.outStack.Pop()
		if err != nil {
			return err
		}

		queue.inStack.Push(val)
	}

	return nil
}

// Sort a stack in-place so that the smallest elements are on top, using only
// an additional temporary stack.
func SortStack(stack Stack) {
	tmpStack := NewBasicStack()
	isSorted := false

	for !isSorted {
		sortStackPass(stack, tmpStack, false)
		isSorted = sortStackPass(tmpStack, stack, true)
	}
}

// Make a single pass at sorting inStack into outStack. Adjacent out-of-order
// elements are swapped.
func sortStackPass(inStack, outStack Stack, reverse bool) bool {
	isSorted := true

	var cmp func(x, y int) bool
	if reverse {
		cmp = func(x, y int) bool { return x >= y }
	} else {
		cmp = func(x, y int) bool { return x <= y }
	}

	for {
		nextVal, err := inStack.Pop()
		if err != nil { // inStack is empty
			return isSorted
		}

		outStackTop, err := outStack.Pop()
		if err != nil {
			outStack.Push(nextVal)
		} else if cmp(outStackTop.(int), nextVal.(int)) {
			outStack.Push(outStackTop)
			outStack.Push(nextVal)
		} else {
			isSorted = false
			outStack.Push(nextVal)
			outStack.Push(outStackTop)
		}
	}
}

// The animal shelter holds both dogs and cats. Animals may be enqueued any
// time. There are three dequeue operations: DequeueCat/Dog returns the
// Cat or Dog that has been in the shelter for longest, and DequeueAny
// returns whichever animal has been in the shelter the longest regardless
// of its type.
type Animal struct {
	animalType int
	name       string
	age        int
	entrySeq   int64
}

type animalNode struct {
	animal Animal
	next   *animalNode
}

type animalQueue struct {
	first *animalNode
	last  *animalNode
}

type animalShelter struct {
	dogQueue animalQueue
	catQueue animalQueue
	currSeq  int64
}

func NewAnimalShelter() *animalShelter {
	return &animalShelter{}
}

const (
	Dog int = iota
	Cat
)

// Adds a new animal to the shelter.
func (shelter *animalShelter) Enqueue(animal Animal) error {
	animal.entrySeq = shelter.currSeq

	switch animal.animalType {
	case Dog:
		shelter.dogQueue.Add(animal)
		shelter.currSeq++
		return nil

	case Cat:
		shelter.catQueue.Add(animal)
		shelter.currSeq++
		return nil

	default:
		return errors.New("Unexpected animal type")
	}
}

// Dequeue the animal that has been in the shelter for longest - either a dog
// or a cat.
func (shelter *animalShelter) DequeueAny() (Animal, error) {
	nextCat, catErr := shelter.catQueue.Peek()
	nextDog, dogErr := shelter.dogQueue.Peek()

	// If both queues returned errors that means the shelter is empty.
	if catErr != nil && dogErr != nil {
		return Animal{}, errors.New("No animals in shelter")
	}

	// If the cat queue is empty of the next cat has a higher sequence number
	// than the next dog, we return the next dog.
	if catErr != nil || nextCat.entrySeq > nextDog.entrySeq {
		_, err := shelter.dogQueue.Remove()
		if err != nil {
			panic(err)
		}
		return nextDog, nil
	}

	// Otherwise, we return the next cat.
	_, err := shelter.catQueue.Remove()
	if err != nil {
		panic(err)
	}
	return nextCat, nil
}

// Dequeue the next dog.
func (shelter *animalShelter) DequeueDog() (Animal, error) {
	return shelter.dogQueue.Remove()
}

// Dequeue the next cat.
func (shelter *animalShelter) DequeueCat() (Animal, error) {
	return shelter.catQueue.Remove()
}

// Add an animal to a queue.
func (queue *animalQueue) Add(animal Animal) {
	newNode := &animalNode{animal: animal}
	if queue.last != nil {
		queue.last.next = newNode
	}
	queue.last = newNode
	if queue.first == nil {
		queue.first = queue.last
	}
}

// Remove an animal from a queue.
func (queue *animalQueue) Remove() (Animal, error) {
	if queue.IsEmpty() {
		return Animal{}, errors.New("Queue is empty")
	}

	animal := queue.first.animal
	queue.first = queue.first.next
	if queue.first == nil {
		queue.last = nil
	}
	return animal, nil
}

// Peek at the next element in the queue.
func (queue animalQueue) Peek() (Animal, error) {
	if queue.IsEmpty() {
		return Animal{}, errors.New("Queue is empty")
	}

	return queue.first.animal, nil
}

// Check if the queue is empty.
func (queue animalQueue) IsEmpty() bool {
	return queue.first == nil
}
