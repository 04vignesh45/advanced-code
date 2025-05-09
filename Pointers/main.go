package main

import "fmt"

type Record struct {
	Name string
	Age  int
}

func (a Record) value() {
	fmt.Println("Name is:", a.Name)
	fmt.Println("Age is:", a.Age)
}

func (c *Record) pointer(name string, age int) {
	c.Name = name
	c.Age = age
}

func main() {
	value := Record{
		Name: "John",
		Age:  25,
	}
	fmt.Println("value type")
	value.value()

	ref := &Record{
		Name: "vignesh",
		Age:  25,
	}
	ref.pointer("vigneshKanth", 26)
	ref.value()
}
