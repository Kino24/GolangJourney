package main

import (
	"fmt"
	"time"
)

func main() {

	i := time.Now()

	fmt.Println("The time now is:")
	fmt.Println(i)

	switch {
	case i.Hour() < 6:
		fmt.Println("Go to sleep bitch!")
		fmt.Println("type this:", "nakakapagpapabagabag")
		var hoho = "nakakapagpapabagabag"
		var eh string
		fmt.Scanln(&eh)
		if eh == hoho {
			fmt.Println("very good now go to sleep.")
		} else {
			fmt.Println("You seriously need to go to sleep")
		}

	case i.Hour() < 12:
		fmt.Println("Good morning mother fucking bitch")
		fmt.Println(" ")
		fmt.Println("This is your morning math exercise!")
		fmt.Println(" ")
		fmt.Println("Type in two numbers.")
		fmt.Println("Please enter value of x.")
		var one float32
		fmt.Scanln(&one)
		fmt.Println("Please enter value of y.")
		var two float32
		fmt.Scanln(&two)
		var A = ((one + 1) / (two + 5)) * 5
		fmt.Println("What is ((x+1)/(y+1)) * 5")
		var ans float32
		fmt.Scanln(&ans)
		if ans == A {
			fmt.Println("You are correct! the answer is", A)
		} else {
			fmt.Println(ans, "is not the answer.")
			fmt.Println(A, "is the right answer.")
		}

	case i.Hour() < 17:
		fmt.Println("Good afternoon bitch!")

	default:
		fmt.Println("Nighty night")

	}
}
