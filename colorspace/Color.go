package colorspace

import (
	"dithering/colorspace/cie"
)

type Color interface {
	XYZ() cie.XYZ
}

func ToCIELab(col Color) cie.Lab {
	xyz := col.XYZ()
	return xyz.Lab()
}

func ToCIELuv(col Color) cie.LCHab {
	xyz := col.XYZ()
	return xyz.LCHab()
}
