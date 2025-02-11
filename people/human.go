package people

import "fmt"

type Human struct {
	name   string
	height int
}

func NewHuman(name string, height int) *Human {
	return &Human{name, height}
}

func (h Human) Print() {
	fmt.Printf("Человек с именем %s и ростом %d", h.name, h.height)
}
