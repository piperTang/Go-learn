package main

import (
	"code.google.com/p/go-tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	s_arr:= strings.Fields(s)
	fmt.Println(s_arr)
	return map[string]int{"x": 1}
}

func main() {
	s:="I am a good student"
	wc.Test(WordCount(s))
}