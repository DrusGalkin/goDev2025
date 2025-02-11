package people

import "fmt"

type Name struct {
	surname,
	name,
	patronymic string
}

type Builder struct {
	name *Name
}

func NewName() *Builder {
	return &Builder{name: new(Name)}
}

func (b *Builder) SetName(name string) *Builder {
	(*b).name.name = name
	return b
}

func (b *Builder) SetSurname(surname string) *Builder {
	(*b).name.surname = surname
	return b
}

func (b *Builder) SetPatronymic(patronymic string) *Builder {
	(*b).name.patronymic = patronymic
	return b
}

func (b Builder) Create() *Name {
	return b.name
}

func (n Name) Print() {
	fmt.Printf("%s %s %s\n", n.surname, n.name, n.patronymic)
}
