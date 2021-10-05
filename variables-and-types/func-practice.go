package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Printf("\nonly %d of fuel left", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
switch planet {
  case "Venus": fuel = 300000
  case "Mercury": fuel = 500000
  case "Mars": fuel = 700000
  default: fuel = 0
}
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Printf("you are visiting %s", planet)
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We dont have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)

  if (fuelRemaining > fuelCost){
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}

func main() {
  // Test your functions!
  
  // Create `planetChoice` and `fuel`
  planetChoice := "Venus"
  fuel := 1000000

  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}