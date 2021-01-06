package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x int
	y int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}
func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) * (x - y))
	return color.RGBA{v, v, 255, 255}
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}
func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
