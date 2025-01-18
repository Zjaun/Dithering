package cie

import (
	"dithering/utils"
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

// returned angle is measured in degrees
func lchFunc(a, b, c float64) lch {
	return lch{
		l: a,
		c: math.Sqrt(math.Pow(b, 2) + math.Pow(c, 2)),
		h: utils.PiecewiseArctan(b, c),
	}
}

func (x *XYZ) Lab() Lab {
	rw := &D65
	xr, yr, zr := x.X/rw.X, x.Y/rw.Y, x.Z/rw.Z
	fx, fy, fz := LabFunc(xr), LabFunc(yr), LabFunc(zr)
	return Lab{
		L: 116*fy - 16,
		A: 500 * (fx - fy),
		B: 200 * (fy - fz),
	}
}

func (x *XYZ) Luv() Luv {
	rw := &D65
	yr := x.Y / rw.Y
	up, vp := LuvFunc(4*x.X, x), LuvFunc(9*x.Y, x)
	upr, vpr := LuvFunc(4*rw.X, rw), LuvFunc(9*rw.Y, rw)
	var l float64
	if yr > E {
		l = 116*math.Cbrt(yr) - 16
	} else {
		l = K * yr
	}
	return Luv{
		L: l,
		U: 13 * l * (up - upr),
		V: 13 * l * (vp - vpr),
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
