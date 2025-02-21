package people

type Department struct {
	name string
	emp  Employees
}

func (d Department) GetName() string {
	return d.name
}

func (d Department) GetEmployee() Employees {
	return d.emp
}
