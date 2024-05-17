package main

import (
	"fmt"
)

func main(){
	//① var 変数名型
	var i int
	i = 10
	fmt.Println("変数", i)
	//② var 変数名 = 型
	var j = 10

	fmt.Println("変数", j)
	//③ 変数名 := 値
	k := 10

	fmt.Println("変数", k)
	//④ 変数名1, 変数名2 := 値1, 値2
	l, m := 10, "hoge"
	fmt.Println("変数", l, m)
}