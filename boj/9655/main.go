package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Println("CY")
	} else {
		fmt.Println("SK")
	}
}
