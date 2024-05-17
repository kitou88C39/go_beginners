package main

import "fmt"

func isEven(num int) {
	if num % 2 == 0 {
		fmt.Printf("%d is even \n", num)
	}
}

func branch (){
	arr := []int{2,3,5}
	for _, num := range arr {
		isEven(num)
  }
}