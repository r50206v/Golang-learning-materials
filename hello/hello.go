package main 
/*
	This line is called a package declaration and every Go program
	starts with one. The package declaration informs the compiler 
	whether to create an executable or library. 
*/

// import "fmt"
// import "time"
import (
	"fmt" 
	t "time"
)
/*
	package name is enclosed with double quotes "
*/

func main(){
	fmt.Printf("hello, world!\n")
	fmt.Println(t.Now())
}
/*
	a main function is special, a file that has a 
	package main declaration will automatically run the main function!
*/

