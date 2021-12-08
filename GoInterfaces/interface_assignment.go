package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}
type squre struct{
	squrelength float64
}

func main() {
	tr := triangle{
		base: 5,
		height: 7,
	}
	sq := squre{
		squrelength: 4,
	}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return (0.5) *t.base*t.height
}

func (s squre) getArea() float64 {
	return s.squrelength*s.squrelength
}