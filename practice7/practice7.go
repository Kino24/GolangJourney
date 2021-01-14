package main

import (
	"fmt"
	"math"
)

var speedCarA float64
var speedCarB float64
var hour float64
var min float64
var carDistanceA float64
var carDistanceB float64
var shortDistance float64

func main() {
	fmt.Println("Input car A speed")
	fmt.Scanln(&speedCarA)
	fmt.Println("Input car B speed")
	fmt.Scanln(&speedCarB)
	fmt.Println("Input time in hours and minutes")
	fmt.Println("Enter Hours")
	fmt.Scanln(&hour)
	fmt.Println("Enter Minutes")
	fmt.Scanln(&min)

	carDistanceA = (hour + (min / 60)) * speedCarA
	carDistanceB = (hour + (min / 60)) * speedCarB
	shortDistance = math.Sqrt(math.Pow(carDistanceA, 2) + math.Pow(carDistanceB, 2))
	fmt.Println(shortDistance)

}
