package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var shell [][]uint8
	for y := 0; y <= dy; y++ {
		var core []uint8
		for x := 0; x <= dx; x++ {
			core = append(core, uint8(y*y*x*x))
		}
		shell = append(shell, core)
	}
	return shell
}

func main() {
	pic.Show(Pic)
}
