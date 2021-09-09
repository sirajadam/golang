package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	
	fmt.Println(GenerateRandomArray(5,100, true))
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