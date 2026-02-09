package main

import "fmt"

type person struct {
	Name string
	Age int
}

type greeter interface {
	Greet() string
}

type counter interface {
	Increment() int
}

type nameGreeter struct {
	name string
}
func (vg nameGreeter) Greet() string {
	return "Hello, " + vg.name
}

type pointerCounter struct {
	value int
}
func (c *pointerCounter) Increment() int {
	c.value++
	return c.value
}

func main() {
	x := 42
	px := &x

	fmt.Printf("x is at address %p and has value %d\n", px, x)

	*px *= 2
	fmt.Printf("x is at address %p and has value %d\n", px, x)

	px = new(int)
	fmt.Printf("x is at address %p and has value %d\n", px, *px)

	*px = 42
	func (val *int) {
		*val *= 2
	}(px)
	fmt.Printf("x is at address %p and has value %d\n", px, *px)

	alice := &person{
		Name: "Alice",
		Age: 20,
	}
	fmt.Printf("alice is at address %p and has value %+v\n", &alice, alice)

	manipulatePerson(alice)
	fmt.Printf("alice is at address %p and has value %+v\n", &alice, alice)
	
}

func manipulatePerson(p *person) {
	p.Name = "Bob"
	p.Age = 30
	fmt.Printf("alice is at address %p and has value %+v\n", &p, p)
}
