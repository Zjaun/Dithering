package colorspace

import (
	"dithering/colorspace/cie"
)

type Color interface {
	XYZ() cie.XYZ
}
