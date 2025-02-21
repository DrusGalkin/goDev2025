package main

import (
	"GoLandTest/geometry"
	"GoLandTest/people"
	//"GoLandTest/people"
)

func main() {
	//1.2.1
	p1 := geometry.NewPoint(1, 5)
	p2 := geometry.NewPoint(5, 10)
	l := geometry.NewLine(p1, p2)
	l.Print()

	//1.2.2
	h := people.NewHuman("Андрей", 20)
	n := people.NewName().SetName("Андрей").SetPatronymic("Андреевич").SetSurname("Галкин").Create()
	ph := people.NewHumanWhisName(n, h)

	ph.Print()

	//1.2.3

	//1.2.4

}
