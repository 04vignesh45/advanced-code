package main

import "fmt"

// First struct
type Circle struct {
    radius float64
}

// Second struct
type Rectangle struct {
    length, width float64
}

// Same method name, but different receivers
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
    return r.length * r.width
}

func main() {
    c := Circle{radius: 5}
    r := Rectangle{length: 4, width: 6}

    fmt.Println("Circle Area:", c.Area())      // Calls Circle's Area method
    fmt.Println("Rectangle Area:", r.Area())   // Calls Rectangle's Area method
}
