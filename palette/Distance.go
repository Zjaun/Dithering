package palette

import (
	"dithering/colorspace"
	"dithering/utils"
	"math"
)

type function func(ref, samp colorspace.Color) float64

type Distance int
type application int

type factor struct {
	kl, k1, k2 float64
}

const (
	SRGB Distance = iota
	CIE76
	CMC
	CIE94ARTS
	CIE94TEXTILE
	CIEDE2000
)

const (
	GraphicArts application = iota
	Textiles
)

var factors = []factor{
	{1, 0.045, 0.015},
	{2, 0.048, 0.014},
}

var Functions = []function{
	Cie76, cie94arts, cie94textile, Cie2000,
}

// This function is symmetrical; order of the parameters does not matter
//
// Uses squared distance to save unnecessary computing
//func Srgb(ref, samp colorspace.Color) float64 {
//	refXyz, sampXyz := ref.XYZ(), samp.XYZ()
//	refRgb, sampRgb := refXyz.Luv(), sampXyz.Luv()
//	return math.Pow(refRgb.R-sampRgb.R, 2) +
//		math.Pow(refRgb.G-sampRgb.G, 2) +
//		math.Pow(refRgb.B-sampRgb.B, 2)
//}

// This function is symmetrical; order of the parameters does not matter
func Cie76(ref, samp colorspace.Color) float64 {
	refXyz, sampXyz := ref.XYZ(), samp.XYZ()
	refLab, sampLab := refXyz.Lab(), sampXyz.Lab()
	return math.Sqrt(
		math.Pow(refLab.L-sampLab.L, 2) +
			math.Pow(refLab.A-sampLab.A, 2) +
			math.Pow(refLab.B-sampLab.B, 2))
}

// This function is asymmetrical; order of the parameters does matter
func Cie94(ref, samp colorspace.Color, app application) float64 {
	fact := factors[app]
	refXyz, sampXyz := ref.XYZ(), samp.XYZ()
	refLab, sampLab := refXyz.Lab(), sampXyz.Lab()
	dL := refLab.L - sampLab.L
	C1 := math.Sqrt(refLab.A*refLab.A + refLab.B*refLab.B)
	C2 := math.Sqrt(sampLab.A*sampLab.A + sampLab.B*sampLab.B)
	dC, da, db := C1-C2, refLab.A-sampLab.A, refLab.B-sampLab.B
	dH := da*da + db*db - dC*dC
	SC, SH := 1+fact.k1*C1, 1+fact.k2*C1
	add1 := (dL * dL) / (fact.kl * fact.kl)
	add2 := (dC * dC) / (SC * SC)
	add3 := dH / (SH * SH)
	return math.Sqrt(add1 + add2 + add3)
}

func cie94arts(ref, samp colorspace.Color) float64 {
	return Cie94(ref, samp, GraphicArts)
}

func cie94textile(ref, samp colorspace.Color) float64 {
	return Cie94(ref, samp, Textiles)
}

// This function is asymmetrical; order of the parameters does matter
func Cie2000(ref, samp colorspace.Color) float64 {
	refXyz, sampXyz := ref.XYZ(), samp.XYZ()
	refLab, sampLab := refXyz.Lab(), sampXyz.Lab()
	Lp := (refLab.L + sampLab.L) / 2
	C1 := math.Sqrt(refLab.A*refLab.A + refLab.B*refLab.B)
	C2 := math.Sqrt(sampLab.A*sampLab.A + sampLab.B*sampLab.B)
	C := (C1 + C2) / 2
	c7 := math.Pow(C, 7)
	tf7 := math.Pow(25, 7)
	G := 0.5 * (1 - math.Sqrt(c7/(c7+tf7)))
	g1 := 1 + G
	ap1 := refLab.A * (g1)
	ap2 := sampLab.A * (g1)
	Cp1 := math.Sqrt(ap1*ap1 + refLab.B*refLab.B)
	Cp2 := math.Sqrt(ap2*ap2 + sampLab.B*sampLab.B)
	Cp := (Cp1 + Cp2) / 2
	Cp7 := math.Pow(Cp, 7)
	hp1 := utils.PiecewiseArctan(refLab.B, ap1)
	hp2 := utils.PiecewiseArctan(sampLab.B, ap2)
	Hp := utils.AverageHue(hp1, hp2)
	t1, t2, t3, t4 := utils.Cos(Hp-30), utils.Cos(2*Hp), utils.Cos(3*Hp+6), utils.Cos(4*Hp-63)
	T := 1 - 0.17*t1 + 0.24*t2 + 0.32*t3 - 0.20*t4
	dhp := utils.DiffHue(hp1, hp2)
	dLp := sampLab.L - refLab.L
	dCp := Cp2 - Cp1
	dHp := 2 * math.Sqrt(Cp1*Cp2) * utils.Sin(dhp/2.0)
	Lp2 := math.Pow(Lp-50, 2)
	SL := 1 + (0.015*Lp2)/math.Sqrt(20+Lp2)
	SC := 1 + 0.045*Cp
	SH := 1 + 0.015*Cp*T
	idt := math.Pow(Hp-275, 2) / math.Pow(25, 2)
	dt := 30 * math.Exp(-idt)
	RC := 2 * math.Sqrt(Cp7/(Cp7+tf7))
	RT := -RC * utils.Sin(2*dt)
	add1 := (dLp * dLp) / (SL * SL)
	add2 := (dCp * dCp) / (SC * SC)
	add3 := (dHp * dHp) / (SH * SH)
	return math.Sqrt(add1 + add2 + add3 + RT*(dCp/SC)*(dHp/SH))
}
