package employee

type Employee struct {
	FirstName  string
	LastName   string
	TotalLeave int
	Leavetaken int
}

func New(
	Firstname string,
	lastname string,
	totalleave int,
	leavetaken int) Employee {

	e := Employee{
		FirstName:  Firstname,
		LastName:   lastname,
		TotalLeave: totalleave,
		Leavetaken: leavetaken,
	}
	return e
}

func (e Employee) LeavesRemaining() int {
	return e.TotalLeave - e.Leavetaken
}
