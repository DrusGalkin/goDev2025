package people

import "fmt"

type Employees struct {
	name string
	dep  *Department
}

func NewEmployee(name string) Employees {
	return Employees{name: name}
}

func (e Employees) Print() {
	if e.dep == nil {
		fmt.Println("Сотрудник с именем " + e.name)
	} else if e == e.dep.GetEmployee() {
		fmt.Printf("%s начальник отдела %")
	} else {
		fmt.Printf("%s работает в отделе %s, начальник %s", e.name, e.dep.GetName(), e.dep.GetEmployee().name)
	}
}
