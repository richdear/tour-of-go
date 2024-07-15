package main

import (
	"fmt"
)

func add(x, y int) int{
	return x+y
}

func main(){
	fmt.Printf("Sum of %g and %g is %g",3,4,add(3,4))
}

