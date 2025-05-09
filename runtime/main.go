package main

import "fmt"

type Shape interface {
	Area() float64
}

type Calculator struct {
	Radius float64
}

type Calculators struct {
	Radius float64
}

func printarea(s Shape) {
	fmt.Println("Area ", s.Area())
}

func (a Calculator) Area() float64 {
	return a.Radius + a.Radius
}

func (a Calculators) Area() float64 {
	return a.Radius - a.Radius
}
func main() {

	a := Calculator{
		Radius: 10,
	}

	b := Calculators{
		Radius: 10,
	}
	printarea(a)
	printarea(b)

}
