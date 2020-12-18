package main

import (
	"fmt"
	"math"
)

func main() {

	var choice = 1
	var choice2 = 2

	fmt.Println("Please enter a number.")
	fmt.Println("moles[1] grams[2]")
	var find int
	fmt.Scanln(&find)

	if find == choice {
		fmt.Println("Enter grams of first element: ")
		var g1 float64
		fmt.Scanln(&g1)

		fmt.Println("Enter enter molar mass of first element: ")
		var mol1 float64
		fmt.Scanln(&mol1)

		var moles1 = g1 * 1 / mol1

		fmt.Println("Enter Mole Ratio of second element: ")
		var mr1 float64
		fmt.Scanln(&mr1)

		fmt.Println("Enter mole ratio of first element: ")
		var mr2 float64
		fmt.Scanln(&mr2)

		var mr = mr1 / mr2
		var moles2 = moles1 * mr
		var finale = math.Round(moles2)

		fmt.Println("The element contains ", finale, " moles.")
	} else {
		if find == choice2 {
			fmt.Println("Enter grams of first element: ")
			var g1 float32
			fmt.Scanln(&g1)

			fmt.Println("Enter enter molar mass of first element: ")
			var mol1 float32
			fmt.Scanln(&mol1)

			var moles1 = g1 * 1 / mol1

			fmt.Println("Enter Mole Ratio of second element: ")
			var mr1 float32
			fmt.Scanln(&mr1)

			fmt.Println("Enter mole ratio of first element: ")
			var mr2 float32
			fmt.Scanln(&mr2)

			var mr = mr1 / mr2
			var moles2 = moles1 * mr

			fmt.Println("Enter molar mass of second element")
			var mol2 float32
			fmt.Scanln(&mol2)

			var mm = moles2 * mol2

			fmt.Println("The element weighs ", mm, " grams.")
		} else {
			fmt.Println(" ")
			fmt.Println(find, "is not part of the choices bitch!")
		}
	}
}
