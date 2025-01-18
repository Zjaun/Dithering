package utils

import "math"

func radToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// parameters must be in degrees, angle returned is in degrees
func Atan2(b, a float64) float64 {
	return radToDeg(math.Atan2(degToRad(b), degToRad(a)))
}

// parameters must be in degrees, angle returned is in radians
func Sin(x float64) float64 {
	return math.Sin(degToRad(x))
}

// parameters must be in degrees, angle returned is in radians
func Cos(x float64) float64 {
	return math.Cos(degToRad(x))
}

// Ensures that angles returned by this function is in the interval of [0, 360]
func PiecewiseArctan(b, a float64) (val float64) {
	if temp := Atan2(b, a); temp >= 0 {
		val = temp
	} else {
		val = temp + 360
	}
	return
}

// angle returned is in degrees
func AverageHue(h1, h2 float64) (val float64) {
	if temp := math.Abs(h1 - h2); temp > 180 {
		val = (h1 + h2 + 360) / 2
	} else {
		val = (h1 + h2) / 2
	}
	return
}

// angle returned is in degrees
func DiffHue(h1, h2 float64) (val float64) {
	if temp := h2 - h1; math.Abs(temp) <= 180 {
		val = temp
	} else if math.Abs(temp) > 180 && h2 <= h1 {
		val = temp + 360
	} else {
		val = temp - 360
	}
	return
}
