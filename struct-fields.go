package main

import "fmt"

type Veltex struct {
	X int
	Y int
}

func main()  {
	v:=Veltex{1,2}
	fmt.Println(v)
	v.X=4
	fmt.Println(v.X)
}