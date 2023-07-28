package util

import (
	"image"

	"github.com/fogleman/gg"
)

func GetPart(src image.Image, x int, y int, w int, h int, scale int) *gg.Context {
	dst := gg.NewContext(w*scale, h*scale)
	
}