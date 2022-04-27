package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(s)
	fmt.Println( []byte( s ) )
	fmt.Println( string( []byte( s ) ))

	for i, v := range []byte(s) {
		fmt.Println(i, v)
	}

	fmt.Println([]byte(s)[2])
	fmt.Println(string([]byte(s)[2]))
	
	fmt.Println(string([]byte(s)[3:]))
}