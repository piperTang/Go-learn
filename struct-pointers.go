package main

import "fmt"

type Veltex struct {
	X int
	Y int
}

func main()  {
	v:=Veltex{1,2}
	p:=&v
	p.X=1e9
	fmt.Println(v)

}