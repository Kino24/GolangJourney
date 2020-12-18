package main

import "fmt"

var i, j int = 1, 2

func fprint() {
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
}

func main() {
	fprint()
	var c, python, java = false, true, "kainin tite"
	fmt.Println(i, j, c, python, java)
}
