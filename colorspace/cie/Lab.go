package cie

import "math"

type Lab struct {
	L, A, B float64
}

func LabFunc(normComp float64) float64 {
	if normComp > E {
		return math.Cbrt(normComp)
	} else {
		return (K*normComp + 16) / 116
	}
}

func YFunc(normComp float64) (val float64) {
	if normComp > K*E {
		return math.Pow((normComp+16)/116, 3)
	} else {
		return normComp / K
	}
}

// used for calculating xr and zr (Lab to XYZ)
func xzFunc(val float64) float64 {
	if temp := math.Pow(val, 3); temp > E {
		return temp
	} else {
		return (116*val - 16) / K
	}
}

func (l *Lab) XYZ() XYZ {
	fy := (l.L + 16) / 116
	fz := fy - l.B/200
	fx := l.A/500 + fy
	xr, yr, zr := xzFunc(fx), YFunc(l.L), xzFunc(fz)
	return XYZ{
		X: xr * D65.X,
		Y: yr * D65.Y,
		Z: zr * D65.Z,
	}
}
