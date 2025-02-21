package main

import (
	"fmt"
	"time"
)

type Print interface {
	Print()
}

type User struct {
	name string
	age  int
}

type Document struct {
	Name         string
	DocumentType string
	time         time.Time
}

func (u User) Print() {
	fmt.Printf("%s, %d года\n", u.name, u.age)
}

func (d Document) Print() {
	fmt.Printf("%s, %s, %d\n", d.Name, d.DocumentType, d.time)
}

func Process(o ...Print) {
	for _, el := range o {
		el.Print()
	}
}
