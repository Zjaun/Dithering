package cie

import (
	"math"
)

type XYZ struct {
	X, Y, Z float64
}

type lch struct {
	l, c, h float64
}

func (x *XYZ) XYZ() XYZ {
	return *x
}

func radToDeg(rad float64) float64 {
	return rad * (math.Pi / 180)
}

func luvFunc(num float64, x *XYZ) float64 {
	return num / (x.X + 15*x.Y + 3*x.Z)
}

func labFunc(normComp float64) (val float64) {
	if normComp > E {
		val = math.Cbrt(normComp)
	} else {
		val = (K*normComp + 16) / 116
	}
	return
}

func lchFunc(a, b, c float64) lch {
	var h float64
	if temp := math.Atan2(c, b); temp >= 0 {
		h = radToDeg(temp)
	} else {
		h = radToDeg(temp) + 360
	}
	return lch{
		l: a,
		c: math.Sqrt(math.Pow(b, 2) + math.Pow(c, 2)),
		h: h,
	}
}

func (x *XYZ) Lab() Lab {
	rw := D65
	xr, yr, zr := x.X/rw.X, x.Y/rw.Y, x.Z/rw.Z
	fx, fy, fz := labFunc(xr), labFunc(yr), labFunc(zr)
	return Lab{
		L: 116*fy - 16,
		A: 500 * (fx - fy),
		B: 200 * (fy - fz),
	}
}

func (x *XYZ) Luv() Luv {
	rw := D65
	yr := x.Y / rw.Y
	wU, wV := luvFunc(4*x.X, x), luvFunc(9*x.X, x)
	rU, rV := luvFunc(4*x.X, &rw), luvFunc(9*x.X, &rw)
	var l float64
	if yr > E {
		l = 116*math.Cbrt(yr) - 16
	} else {
		l = K * yr
	}
	return Luv{
		L: l,
		U: 13 * l * (wU - rU),
		V: 13 * l * (wV - rV),
	}
}

func (x *XYZ) LCHab() LCHab {
	lab := x.Lab()
	cyl := lchFunc(lab.L, lab.A, lab.B)
	return LCHab{
		L: cyl.l,
		C: cyl.c,
		H: cyl.h,
	}
}

func (x *XYZ) LCHuv() LCHuv {
	luv := x.Luv()
	cyl := lchFunc(luv.L, luv.U, luv.V)
	return LCHuv{
		L: cyl.l,
		C: cyl.c,
		H: cyl.h,
	}
}
