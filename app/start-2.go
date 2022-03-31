package main
import "fmt"

type hotdog int
type person struct {
	fname string
	lname string
	// Fname (uppercase) can make this parameter visible outside of pacakge
	// fname (lowercase) will make this parameter invisible outside of pacakge
}


func main(){
	var t hotdog
	t = 6
	fmt.Println(t)

	xi := [] int {1,2,3,4,5}
	fmt.Println(xi)

	m := map[string] int {
		"Todd": 45,
		"Richard": 19,
		"Joe": 3,
	}
	fmt.Println(m)

	pi := person{
		"Yi Ping",
		"Tseng",
	}
	fmt.Println(pi)
}