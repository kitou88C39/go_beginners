package main

import "fmt"

func pointer() {
//②var ポインタ変数 *型
	var p *int
	i := 10
	p = &i

	*p = 11
	fmt.Println(*p)
}

