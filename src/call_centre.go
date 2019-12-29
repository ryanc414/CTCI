package main

import (
	"ctci"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	callCentre := InitCallCentre(6, 3, 1)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		caller := &Caller{id: i}
		wg.Add(1)
		go callCentre.HandleCall(&wg, caller)
	}

	wg.Wait()
}

type CallCentre struct {
	respondents *EmployeePool
	managers    *EmployeePool
	directors   *EmployeePool
	callerQueue []Caller
}

type EmployeeCategory int

const (
	Respondent = iota
	Manager
	Director
)

type EmployeePool struct {
	available *ctci.BasicQueue
	busy      map[int]*Employee
	category  EmployeeCategory
	mutex     sync.Mutex
}

type Employee struct {
	id       int
	category EmployeeCategory
}

type Caller struct {
	id int
}

// Initialise a new CallCentre.
func InitCallCentre(
	numRespondents, numManagers, numDirectors int) *CallCentre {
	return &CallCentre{
		respondents: InitEmployeePool(numRespondents, Respondent),
		managers:    InitEmployeePool(numManagers, Manager),
		directors:   InitEmployeePool(numDirectors, Director),
	}
}

// Initialise a new pool of employees.
func InitEmployeePool(
	numEmployees int, category EmployeeCategory,
) *EmployeePool {
	available := ctci.NewBasicQueue()
	for i := 0; i < numEmployees; i++ {
		available.Add(InitEmployee(i, category))
	}

	return &EmployeePool{
		available: available,
		busy:      make(map[int]*Employee),
		category:  category,
	}
}

// Initialise a new employee.
func InitEmployee(
	id int, category EmployeeCategory,
) *Employee {
	return &Employee{
		id:       id,
		category: category,
	}
}

// Handle an incoming call.
func (callCentre *CallCentre) HandleCall(wg *sync.WaitGroup, caller *Caller) {
	defer wg.Done()

	success := callCentre.respondents.HandleCall(caller)
	if success {
		return
	}

	success = callCentre.managers.HandleCall(caller)
	if success {
		return
	}

	success = callCentre.directors.HandleCall(caller)
	if success {
		return
	}

	panic("Could not handle call")
}

// Handle a call by assigning it to an available employee. Returns true if the
// call was successfully handled, false if it needs to be delegated.
func (pool *EmployeePool) HandleCall(caller *Caller) bool {
	pool.mutex.Lock()
	next, err := pool.available.Remove()
	if err != nil {
		pool.mutex.Unlock()
		fmt.Printf(
			"No available employees in category %v\n",
			categoryString(pool.category),
		)
		return false
	}

	employee := next.(*Employee)
	pool.busy[employee.id] = employee
	pool.mutex.Unlock()

	success := employee.HandleCall(caller)

	pool.mutex.Lock()
	delete(pool.busy, employee.id)
	pool.available.Add(employee)
	pool.mutex.Unlock()

	return success
}

// Handle a call as an employee.
func (employee *Employee) HandleCall(caller *Caller) bool {
	category := categoryString(employee.category)

	fmt.Printf(
		"%v #%v handles call from caller %v\n", category, employee.id, caller.id,
	)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	fmt.Printf(
		"%v #%v finished handling call from caller %v\n",
		category,
		employee.id,
		caller.id,
	)

	return true
}

// Get the employee's category as a string.
func categoryString(category EmployeeCategory) string {
	switch category {
	case Respondent:
		return "Respondent"

	case Manager:
		return "Manager"

	case Director:
		return "Director"

	default:
		panic("Unexpected employee category")
	}
}
