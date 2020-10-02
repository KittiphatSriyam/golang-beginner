package main

import (
	"fmt"
)

func main() {
	x := 70
	var p *int
	p = &x
	// *p += 2
	fmt.Println("x : ", x)
	fmt.Println("&x : ", &x)
	fmt.Println("p : ", p)
	fmt.Println("&p : ", &p)
	fmt.Println("value of *p : ", *p)
}
