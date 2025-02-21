package main

import "fmt"

type vehicle interface {
	move()
}

type Car struct {
}

type Aircraft struct {
}

func (c Car) move() {
	fmt.Println("Машина едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func drive(m vehicle) {
	m.move()
}

func main() {
	//var tesla vehicle = Car{}
	//var air vehicle = Aircraft{}
	//drive(tesla)
	//drive(air)

	d := &Document{Name: "aa2a", DocumentType: "fdf342423324324df"}
	d1 := &Document{Name: "aa3a", DocumentType: "fdrewtwefdf"}
	d2 := &Document{Name: "aa54a", DocumentType: "fddsfsdfsdfasdfdf"}
	d3 := &Document{Name: "aa5a", DocumentType: "fdfdsfawsefasdfdf"}
	d4 := &Document{Name: "aa6a", DocumentType: "fdasdfsadfsaed4ertfdf"}

	Process(d, d1, d2, d3, d4)
}
