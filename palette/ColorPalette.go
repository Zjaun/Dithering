package palette

import (
	"dithering/colorspace"
	"math"
)

type ColorPalette struct {
	xyz []colorspace.XYZ
}

func (c *ColorPalette) AddColor(col colorspace.Color) {
	c.xyz = append(c.xyz, col.ToXYZ())
}

func dist(a, b colorspace.XYZ) float64 {
	return math.Pow(a.X-b.X, 2) +
		math.Pow(a.Y-b.Y, 2) +
		math.Pow(a.Z-b.Z, 2)
}

func (c *ColorPalette) NearestColor(col colorspace.Color) colorspace.XYZ {
	conv := col.ToXYZ()

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
