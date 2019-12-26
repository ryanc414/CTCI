package ctci

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
	available *BasicQueue
	busy       map[int]*Employee
	callCentre *CallCentre
    category EmployeeCategory
}

type Employee struct {
	id       int
	category EmployeeCategory
	pool     *EmployeePool
}

type Caller struct {
	name string
}

// Initialise a new CallCentre.
func InitCallCentre(
	numRespondents, numManagers, numDirectors int) *CallCentre {
	callCentre := &CallCentre{}
	callCentre.respondents = InitEmployeePool(
		numRespondents, Respondent, callCentre,
	)
	callCentre.managers = InitEmployeePool(
		numManagers, Manager, callCentre,
	)
	callCentre.directors = InitEmployeePool(
		numDirectors, Director, callCentre,
	)

	return callCentre
}

// Initialise a new pool of employees.
func InitEmployeePool(numEmployees int,
	category EmployeeCategory,
	callCentre *CallCentre) *EmployeePool {
	pool := &EmployeePool{
		available:  NewBasicQueue(),
		busy:       make(map[int]*Employee),
		callCentre: callCentre,
        category: category,
	}

	for i := 0; i < numEmployees; i++ {
		pool.available.Add(InitEmployee(i, category, pool))
	}

	return pool
}

// Initialise a new employee.
func InitEmployee(
	id int, category EmployeeCategory, pool *EmployeePool,
) *Employee {
	return &Employee{
		id:       id,
		category: category,
		pool:     pool,
	}
}

// Handle an incoming call.
func (callCentre *CallCentre) HandleCall(caller *Caller) {
    callCentre.respondents.HandleCall(caller)
}

// Delegate a call.
func (callCentre *CallCentre) Delegate(caller *Caller, from EmployeeCategory) {
    switch from {
    case Employee:
        callCentre.managers.HandleCall(caller)

    case Manager:
        callCentre.directors.HandleCall(caller)

    default:
        panic(fmt.Sprintf("Cannot delegate from employee type %v"))
    }
}

// Handle a call by assigning it to an available employee.
func (pool *EmployeePool) HandleCall(caller *Caller) {
    employee, err := pool.available.Remove()
    if err != nil {
        pool.callCentre.Delegate(caller, pool.category)
    }

    pool.busy[employee.id] = employee
    employee.HandleCall()
}

// Delegate a call from an employee.
func (pool *EmployeePool) Delegate(caller *Caller, from *Employee) {
    delete(pool.busy, employee.id)
    pool.available.Add(employee)
    pool.callCentre.Delegate(caller, from.category)
}

// Handle a call as an employee.
func (employee *Employee) HandleCall(caller *Caller) {
    category := employee.categoryString()

    fmt.Printf(
        "%v #%v handles call from %v\n", category, employee.id, caller.name,
    )

    time.Sleep(rand.Intn(1000) * time.Millisecond)

    fmt.Printf(
        "%v #%v finished handling call from %v\n",
        category,
        employee.id,
        caller.name,
    )
}

// Get the employee's category as a string.
func (employee *Employee)
