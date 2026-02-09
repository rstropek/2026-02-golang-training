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
    ng := nameGreeter{name: "Alice"}
	var ng1 greeter = ng
	var ng2 greeter = &ng
	fmt.Printf("ng1: %+v, ng2: %+v\n", ng1, ng2)

	pc := pointerCounter{value: 0}
	// var pc1 counter = pc
	var pc2 counter = &pc
	fmt.Printf("pc2: %+v\n", pc2)

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
