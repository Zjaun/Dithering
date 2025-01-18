package cie

type Luv struct {
	L, U, V float64
}

func LuvFunc(num float64, x *XYZ) float64 {
	return num / (x.X + 15*x.Y + 3*x.Z)
}

func (l *Luv) XYZ() XYZ {
	u0, v0 := LuvFunc(4*D65.X, &D65), LuvFunc(9*D65.Y, &D65)
	y := YFunc(l.L)
	a := 1.0 / 3 * (52*l.L/(l.U+13*l.L*u0) - 1)
	b := -5 * y
	d := y * (39*l.L/(l.V+13*l.L*v0) - 5)
	x := (d - b) / (a + 1.0/3)
	return XYZ{
		X: x,
		Y: y,
		Z: x*a + b,
	}
}
