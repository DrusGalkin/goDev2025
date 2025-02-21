package geometry

import "fmt"

type Line struct {
	p1 Point
	p2 Point
}

func NewLine(p1, p2 Point) Line {
	return Line{
		p1: p1,
		p2: p2,
	}
}

func (l Line) GetStart() Point {
	return l.p1
}

func (l Line) GetEnd() Point {
	return l.p2
}

func (l *Line) SetStart(p Point) {
	l.p1 = p
}

func (l *Line) SetEnd(p Point) {
	l.p2 = p
}

func (l Line) Print() {
	fmt.Println("Линия от " + fmt.Sprint(l.p1) + " до " + fmt.Sprint(l.p2))
}
