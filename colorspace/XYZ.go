package colorspace

type XYZ struct {
	X, Y, Z float64
}

func (x *XYZ) ToXYZ() XYZ {
	return *x
}
