package main

import (
	"crypto/rand"
	"fmt"
)

func main(){
	//① var 変数名型
	var i int
	i = 10
	fmt.Println("Hello World", rand.Intn(10))
}