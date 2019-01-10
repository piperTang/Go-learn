package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 外层slice
	a := make([][]uint8, dy)
	for x := range a {
		// 里层slice
		b := make([]uint8, dx)
		for y := range b {
			// 给里层slice每个元素赋值
			b[y] = uint8(x*y - 1)
		}
		// 给外层slice每个元素赋值
		a[x] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
