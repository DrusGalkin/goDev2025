package geometry

type Point struct {
	x, y int
}

func (p Point) GetX() int {
	return p.x
}

func (p Point) GetY() int {
	return p.y
}

func (p *Point) SetX(x int) {
	(*p).x = x
}

func (p *Point) SetY(y int) {
	(*p).y = y
}
