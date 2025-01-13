package colorspace

import "math"

type normalizedSRGB struct {
	R, G, B float64
}

type linearRGB struct {
	R, G, B float64
}

type StandardRGB struct {
	R, G, B int
}

func (s *StandardRGB) normalize() normalizedSRGB {
	return normalizedSRGB{
		float64(s.R) / 255.0,
		float64(s.G) / 255.0,
		float64(s.B) / 255.0,
	}
}

func linear(normalized float64) float64 {
	var value float64
	if normalized <= 0.04045 {
		value = normalized / 12.92
	} else {
		value = math.Pow((normalized+0.055)/1.055, 2.4)
	}
	return value
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

func (s *StandardRGB) linear() linearRGB {
	normalized := s.normalize()
	return linearRGB{
		linear(normalized.R),
		linear(normalized.G),
		linear(normalized.B),
	}
}

func (l *linearRGB) standard() StandardRGB {
	return StandardRGB{
		nonlinear(l.R),
		nonlinear(l.G),
		nonlinear(l.B),
	}
}
