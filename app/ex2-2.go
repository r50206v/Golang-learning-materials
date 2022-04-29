package main

import "fmt"

func main() {
	m := map[string]int{"Juan": 87, "Allon": 18, "Rodd": 42}

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}