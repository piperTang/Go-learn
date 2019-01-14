package main;

import "fmt"

type Vertex struct {
	Lat,Long float64
}
var m=map[string]Vertex{
	"Bell Labs":{40.232,4242.232},
	"Google":{32.3233,53.24224},
}

func main()  {
	fmt.Println(m)
}