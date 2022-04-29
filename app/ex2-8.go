package main

import "fmt"

type vehicle struct {
	doors int 
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle{2, "blue"}, 
		false,
	}
	sedan1 := sedan{
		vehicle{4, "grey"}, 
		true,
	}
	
	fmt.Println(truck1)
	fmt.Println(truck1.color)
	fmt.Println(sedan1)
	fmt.Println(sedan1.luxury)
}