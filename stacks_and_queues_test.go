package ctci

import "testing"

// Test the BasicStack type and methods.
func TestBasicStack(t *testing.T) {
	stack := NewBasicStack()
	genericStackTest(t, stack)
}

// Test the multiStacks type and methods.
func TestMultiStacks(t *testing.T) {
	stacks := NewMultiStacks(3)

	for i := 0; i < 3; i++ {
		if !stacks.IsEmpty(i) {
			t.Error()
		}

		_, err := stacks.Pop(i)
		if err == nil {
			t.Error()
		}

		_, err = stacks.Peek(i)
		if err == nil {
			t.Error()
		}

		stacks.Push(i, 1)
		stacks.Push(i, 2)
		stacks.Push(i, 3)

		if stacks.IsEmpty(i) {
			t.Error()
		}

		val, err := stacks.Pop(i)
		if err != nil || val != 3 {
			t.Error()
		}

		val, err = stacks.Peek(i)
		if err != nil || val != 2 {
			t.Error()
		}
	}
}

// Test the minStacks type and methods.
func TestMinStack(t *testing.T) {
	var stack MinStack
	genericStackTest(t, &stack)

	if !stack.IsEmpty() {
		t.Error()
	}

	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	minVal, err := stack.Min()
	if minVal != 1 || err != nil {
		t.Error()
	}

	val, err := stack.Pop()
	if val != 1 || err != nil {
		t.Error()
	}

	val, err = stack.Min()
	if val != 2 || err != nil {
		t.Error()
	}

	stack.Push(0)
	val, err = stack.Peek()
	if val != 0 || err != nil {
		t.Error()
	}

	val, err = stack.Min()
	if val != 0 || err != nil {
		t.Error()
	}
}

// Test the setOfStacks type and methods.
func TestSetOfStacks(t *testing.T) {
	stackSet := NewSetOfStacks(3)
	genericStackTest(t, &stackSet)

	if !stackSet.IsEmpty() {
		t.Error()
	}

	// Push 3 values and peek at the top value. This will fill the first
	// internal stack.
	stackSet.Push(1)
	stackSet.Push(2)
	stackSet.Push(3)

	val, err := stackSet.Peek()
	if val != 3 || err != nil {
		t.Error()
	}

	// Push another 3 values and peek again. This will fill the second internal
	// stack.
	stackSet.Push(4)
	stackSet.Push(5)
	stackSet.Push(6)

	val, err = stackSet.Peek()
	if val != 6 || err != nil {
		t.Error()
	}

	// Pop all values.
	for expectedVal := 6; expectedVal > 0; expectedVal-- {
		val, err = stackSet.Pop()
		if val != expectedVal || err != nil {
			t.Error()
		}
	}

	if !stackSet.IsEmpty() {
		t.Error()
	}
}

// Generic test for any type that implements the stack interface.
func genericStackTest(t *testing.T, stack Stack) {
	if !stack.IsEmpty() {
		t.Error()
	}

	_, err := stack.Pop()
	if err == nil {
		t.Error()
	}

	_, err = stack.Peek()
	if err == nil {
		t.Error()
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, err := stack.Pop()
	if val != 3 || err != nil {
		t.Error()
	}

	val, err = stack.Peek()
	if val != 2 || err != nil {
		t.Error()
	}

	// Pop the rest of the values, so that we leave the stack empty.
	val, err = stack.Pop()
	if val != 2 || err != nil {
		t.Error()
	}

	val, err = stack.Pop()
	if val != 1 || err != nil {
		t.Error()
	}

	if !stack.IsEmpty() {
		t.Error()
	}
}

// Test the BasicQueue type and methods.
func TestBasicQueue(t *testing.T) {
	queue := NewBasicQueue()
	genericQueueTest(t, queue)
}

// Test the MyQueue type and methods.
func TestMyQueue(t *testing.T) {
	queue := NewMyQueue()
	genericQueueTest(t, queue)
}

// Test the queue interface.
func genericQueueTest(t *testing.T, queue Queue) {
	if !queue.IsEmpty() {
		t.Error()
	}

	_, err := queue.Remove()
	if err == nil {
		t.Error()
	}

	_, err = queue.Peek()
	if err == nil {
		t.Error()
	}

	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	if queue.IsEmpty() {
		t.Error()
	}

	// Since the queue has LIFO ordering, the first value to come out should
	// be 1.
	val, err := queue.Peek()
	if val != 1 || err != nil {
		t.Error(val, err)
	}

	// Check that the values are popped in order 1, 2, 3
	for expected := 1; expected < 4; expected++ {
		val, err = queue.Remove()
		if val != expected || err != nil {
			t.Error(val, err)
		}
	}

	// All values have been popped - queue should again be empty.
	if !queue.IsEmpty() {
		t.Error("Queue not empty")
	}
}

// Test the SortStack function.
func TestSortStack(t *testing.T) {
	stack := NewBasicStack()
	stack.Push(2)
	stack.Push(1)
	stack.Push(4)
	stack.Push(3)
	stack.Push(5)

	SortStack(stack)

	for expected := 1; expected < 6; expected++ {
		val, err := stack.Pop()
		if val != expected || err != nil {
			t.Error(val, err)
		}
	}

	if !stack.IsEmpty() {
		t.Fail()
	}
}

// Test the animalShelter type and methods.
func TestAnimalShelter(t *testing.T) {
	shelter := NewAnimalShelter()

	// Try dequeuing animals from an empty shelter - an error should be
	// returned.
	_, err := shelter.DequeueAny()
	if err == nil {
		t.Error()
	}

	_, err = shelter.DequeueDog()
	if err == nil {
		t.Error()
	}

	_, err = shelter.DequeueCat()
	if err == nil {
		t.Error()
	}

	// Now try enqueing some animals.
	err = shelter.Enqueue(Animal{
		animalType: Cat,
		name:       "Lottie",
		age:        12,
	})
	if err != nil {
		t.Error(err)
	}

	err = shelter.Enqueue(Animal{
		animalType: Dog,
		name:       "Spot",
		age:        5,
	})
	if err != nil {
		t.Error(err)
	}

	err = shelter.Enqueue(Animal{
		animalType: Dog,
		name:       "Lucy",
		age:        2,
	})
	if err != nil {
		t.Error(err)
	}

	err = shelter.Enqueue(Animal{
		animalType: Cat,
		name:       "Penny",
		age:        11,
	})
	if err != nil {
		t.Error(err)
	}

	// Let's try dequeuing them.
	animal, err := shelter.DequeueAny()
	if err != nil || animal.name != "Lottie" {
		t.Error(animal, err)
	}

	animal, err = shelter.DequeueCat()
	if err != nil || animal.name != "Penny" {
		t.Error(animal, err)
	}

	animal, err = shelter.DequeueDog()
	if err != nil || animal.name != "Spot" {
		t.Error(animal, err)
	}

	animal, err = shelter.DequeueAny()
	if err != nil || animal.name != "Lucy" {
		t.Error(animal, err)
	}

	// Shelter is now empty.
	_, err = shelter.DequeueAny()
	if err == nil {
		t.Error()
	}
}
