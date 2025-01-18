package rgb

import (
	"dithering/colorspace/cie"
	"math"
)

type normalizedSRGB struct {
	R, G, B float64
}

type linearRGB struct {
	R, G, B float64
}

type SRGB struct {
	R, G, B int
}

var sRGBmatrix = [][]float64{
	{0.4124564, 0.3575761, 0.1804375},
	{0.2126729, 0.7151522, 0.0721750},
	{0.0193339, 0.1191920, 0.9503041},
}

// Ideally, you should check if the first parameter is symmetrical, and
// check if the length of the first row is equal to the length of col
//
// But since we're using this internally, no such checks are needed; just ensure
// that the two requirements are satisfied
func matmul(mat [][]float64, col []float64) []float64 {
	product := make([]float64, len(col))
	for i := range mat {
		for j := range col {
			product[i] += mat[i][j] * col[j]
		}
	}
	return product
}

func (s *SRGB) XYZ() cie.XYZ {
	norm := s.linear()
	prod := matmul(sRGBmatrix, []float64{norm.R, norm.G, norm.B})
	return cie.XYZ{prod[0], prod[1], prod[2]}
}

func (s *SRGB) normalize() normalizedSRGB {
	return normalizedSRGB{
		float64(s.R) / 255.0,
		float64(s.G) / 255.0,
		float64(s.B) / 255.0,
	}
}

func linear(normalized float64) (value float64) {
	if normalized <= 0.04045 {
		value = normalized / 12.92
	} else {
		value = math.Pow((normalized+0.055)/1.055, 2.4)
	}
	return
}

func nonlinear(linear float64) int {
	var value float64
	if linear <= 0.0031308 {
		value = 12.92 * linear
	} else {
		value = math.Pow(1.055*linear, 1/2.4) - 0.055
	}
	return int(value)
}

func (s *SRGB) linear() linearRGB {
	normalized := s.normalize()
	return linearRGB{
		linear(normalized.R),
		linear(normalized.G),
		linear(normalized.B),
	}
}

func (l *linearRGB) standard() SRGB {
	return SRGB{
		nonlinear(l.R),
		nonlinear(l.G),
		nonlinear(l.B),
	}
}
