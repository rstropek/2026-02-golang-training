package main

// Let's talk about structs

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X float32 `json:"x"` // well-known struct tag
	Y float32 `json:"y"`
}

type Rect struct {
	LeftUpper, RightLower Point
}

type Circle struct {
	Center Point
	Radius float32
}

func main() {
	r := Rect{LeftUpper: Point{X: 0, Y: 10}, RightLower: Point{X: 10, Y: 0}}
	fmt.Println(r)

	c := Circle{Center: Point{X: 0, Y: 0}, Radius: 5}
	fmt.Println(c)

	p := Point{X: 1, Y: 2}
	j, _ := json.Marshal(p)
	fmt.Println(string(j))
}
