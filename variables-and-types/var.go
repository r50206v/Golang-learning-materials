package main

import "fmt"

func main() {
	const earthsGravity = 9.80665
	fmt.Println(earthsGravity)

	var numOfFlavors int
	numOfFlavors = 57
	fmt.Println(numOfFlavors)
	
	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)

	var favoriteSnack string = "chocolate"
	fmt.Println("My favorite snack is " + favoriteSnack)

	var emptyInt int8
	var emptyFloat float32
	var emptyString string
	//   default number is 0, string is "", and boolean is false
	fmt.Println(emptyInt, emptyFloat, emptyString)

	//   creating variables without specifying types
	var newIntVariable = 5
	newFloatVariable := 5.0
	fmt.Println(newIntVariable, newFloatVariable)

	var variable1, variable2 int32


	name := "H. J. Simp"

	// Add your switch statement below:
	switch name {
    case "Butch": fmt.Print("Head to Robbers Roost.")

    case "Bonnie": fmt.Print("Stay put in Joplin.")

    default: fmt.Print("Just hide!")
  }
}