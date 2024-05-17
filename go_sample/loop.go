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
	// sum := 1
	// for sum < 10 {
	// 	sum += sum
	// }
	// fmt.Println(sum)
//③ 配列の宣言方法
	
	arr := []string{"hoge","hoge","hoge"}
	for _, num := range arr {
		fmt.Println(num)
	}
}
