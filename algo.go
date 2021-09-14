package main

import "fmt"

// Arrays in GO


func main() {

	a := []int{1,2,3,4}
	b := []int{6,7,8,9}


	if (len(a) > 0 ) {

		fmt.Println(a)
	}

	c := append(a, 5)
	c = append(c, b...)

	fmt.Println(c)
	
}

