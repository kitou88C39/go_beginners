package main

import "fmt"


func branch (){
	arr := []int{2,3,5}
	for _, num := range arr {
		if num % 2 == 0 {
		fmt.Println("%d is even \n", num)
	}
  }
}