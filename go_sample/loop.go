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
	// 配列の宣言
	arr := []int{2,4,6}
	sum := 0  // sumの初期化
	for _, num := range arr {
		sum += num  // 各要素をsumに加算
		fmt.Println(num)
	}
	fmt.Println(sum)  // 最終的なsumの値を出力
}
