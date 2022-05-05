package main 

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 42

	var x int
	x = 5

	y := 6

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	x = g1
	fmt.Println(x)
	// you CANNOT assign g1 to x. They are two different types!
}