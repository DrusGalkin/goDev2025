package people

import "fmt"

type HumanWhisName struct {
	name  *Name
	human *Human
}

func NewHumanWhisName(name *Name, human *Human) HumanWhisName {
	return HumanWhisName{name, human}
}

func (h HumanWhisName) GetName() Name {
	return *h.name
}

func (h HumanWhisName) GetHuman() Human {
	return *h.human
}

func (h HumanWhisName) Print() {
	fmt.Printf("Человек с именем %s и ростом %d", h.name.GetFIO(), h.GetHuman().GetHeight())
}
