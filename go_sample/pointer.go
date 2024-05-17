package main

import "fmt"

func pointer() {
	i := 10
	p := &i
	fmt.Println(*p)
}

