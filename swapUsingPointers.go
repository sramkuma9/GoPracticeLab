package main

import (
	"fmt"
)

func swap(x *int, y *int){
	t := *x
	*x = *y
	*y = t
}

func main() {
	x,y := 10,20
	fmt.Println("before swapping",x,y)
	swap(&x,&y)
	fmt.Print("After swapping ", x, y)	
}
