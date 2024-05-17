package main

import "fmt"

func pointer() {
	i := 10
	p := &i

	*p = 11
	fmt.Println(*p)
}

