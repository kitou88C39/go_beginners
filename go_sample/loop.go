package main

import (
	"fmt"
)

func loop (){
//① for
//    sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
//② whieみたいな書き方
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

}