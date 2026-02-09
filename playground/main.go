package main

// Let's talk about structs

import (
	"encoding/json"
	"fmt"
	"math"
)

type Point struct {
	X float32 `json:"x"` // well-known struct tag
	Y float32 `json:"y"`
}

type Rect struct {
	LeftUpper, RightLower Point
}

func (r Rect) Width() float32 {
	return float32(math.Abs(float64(r.RightLower.X - r.LeftUpper.X)))
}

func (r Rect) Height() float32 {
	return float32(math.Abs(float64(r.RightLower.Y - r.LeftUpper.Y)))
}

func (r Rect) Area() float32 {
	return r.Width() * r.Height()
}

func (r *Rect) Enlarge(factor float32) {
	r.RightLower.X = r.LeftUpper.X + factor * r.Width()
	r.RightLower.Y = r.LeftUpper.Y + factor * r.Height()
}

type Circle struct {
	Center Point
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float32
}

type ObjectWithHeight interface {
	Height() float32
}

type ShapeWithHeight interface {
	Shape
	ObjectWithHeight
}

const (
	WHITE int = 0xFFFFFF
	BLACK int = 0x000000
	RED   int = 0xFF0000
	GREEN int = 0x00FF00
	BLUE  int = 0x0000FF
	YELLOW int = 0xFFFF00
	MAGENTA int = 0xFF00FF
	CYAN int = 0x00FFFF
)

type ColoredCircle struct {
	Circle
	Color int
}

func main() {
	r := Rect{LeftUpper: Point{X: 0, Y: 10}, RightLower: Point{X: 10, Y: 0}}
	fmt.Printf("Width: %f, Height: %f, Area: %f\n", r.Width(), r.Height(), r.Area())
	r.Enlarge(2)
	fmt.Printf("Width: %f, Height: %f, Area: %f\n", r.Width(), r.Height(), r.Area())
	c := Circle{Center: Point{X: 0, Y: 0}, Radius: 5}
	fmt.Println(c)

	p := Point{X: 1, Y: 2}
	j, _ := json.Marshal(p)
	fmt.Println(string(j))

	cc := ColoredCircle{Circle: Circle{Center: Point{X: 0, Y: 0}, Radius: 10}, Color: WHITE}
	fmt.Printf("Radius: %f, Color: %d, Area: %f\n", cc.Radius, cc.Color, cc.Area())

	shapes := []Shape{r, c, cc}
	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}
}
