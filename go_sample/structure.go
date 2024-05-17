package main

import "fmt"

type Rectangle struct  {
	width, height int
}

func structure (){
	r := Rectangle{width:20, height:30}
	fmt.Println(r)
	fmt.Println(r.height)
}