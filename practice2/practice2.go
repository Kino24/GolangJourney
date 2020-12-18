package main

import "fmt"

func main() {

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

	fmt.Println("The element contains ", moles2, " moles.")

}
