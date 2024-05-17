package main

import "fmt"

type Rectangle struct  {
	width, height int
}

func (r Rectangle) Area() float64 {
	return float64(r.width * r.height)
}

func structure (){
	r := Rectangle{width:20, height:30}
	area := r.Area()
	fmt.Println(area)
}