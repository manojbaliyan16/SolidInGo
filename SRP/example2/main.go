package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type shape interface {
	name() string
	area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

func (c Circle) name() string {
	return "Circle"
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) name() string {
	return "Square"
}
func (s Square) area() float64 {
	return s.length * s.length
}

// Need the out in multiple format

type output struct{}

// Text format o/p
func (o output) Text(sh shape) string {
	return fmt.Sprintf("Area of the shape %s  is %f", sh.name(), sh.area())
}

func (o output) JSON(sh shape) string {
	res := struct {
		Name string  `josn:"shape"`
		Area float64 `josn:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}
	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func main() {
	c := Circle{
		radius: 3,
	}
	o := output{}
	fmt.Println(o.Text(c))
	fmt.Println(o.JSON(c))

	s := Square{
		length: 15,
	}

	fmt.Println(o.Text(s))
	fmt.Println(o.JSON(s))

}
