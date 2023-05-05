package main

import "fmt"

// https://ueokande.github.io/go-slice-tricks/

func main () {
	var x int = 0
	a := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	x, a = a[len(a) - 1], a[:len(a)-1]

	fmt.Println(x, a)
}