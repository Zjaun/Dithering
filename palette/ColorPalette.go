package palette

import (
	"dithering/colorspace"
	"dithering/colorspace/cie"
	"math"
)

type ColorPalette struct {
	xyz []cie.XYZ
}

func (c *ColorPalette) AddColor(col colorspace.Color) {
	c.xyz = append(c.xyz, col.XYZ())
}

//func Dist(a, b colorspace.Color, fx Distance) float64 {
//
//}

func (c *ColorPalette) NearestColor(col colorspace.Color, dist func(a, b cie.XYZ) float64) cie.XYZ {
	conv := col.XYZ()

	minDistance := math.MaxFloat64
	nearestColor := conv

	for _, val := range c.xyz {
		d := dist(val, conv)
		if d < minDistance {
			minDistance = d
			nearestColor = val
		}
	}

	return nearestColor
}
