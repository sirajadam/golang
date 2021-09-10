package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	arrTest := GenerateRandomArray(5,100,false);
	
	fmt.Println(bubble(arrTest))
}

func bubble(arr []int) []int {

	n := len(arr)

	for i := 0; i < n; i++ {
		
		for j := 0; j < n - 1; j++ {
			if arr[j+1] < arr[j] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr;
}


func GenerateRandomArray(len, max int, sign bool) []int {
	
	rand.Seed(time.Now().Unix())
	var randArr []int
	for i := 0; i < len; i++ {
        value := rand.Intn(max) + 1
        if sign {
            value = -(rand.Intn(max) + 1)
        } 
	
		randArr = append(randArr, value)
	
	
	}
	return randArr
}