package main

import (
	"fmt"
	"math"
	"strconv"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info (s shape) string {
	fmt.Println(s.area())
	res := strconv.FormatFloat(s.area(), 'E', -1, 64)
	return res
}

func main() {
	sq := square{40}
	ci := circle{20}
	sq_info := info(sq)
	ci_info := info(ci)
	fmt.Println(sq_info)
	fmt.Println(ci_info)
}