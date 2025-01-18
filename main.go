package main

import (
	"dithering/colorspace/rgb"
	"dithering/palette"
	"fmt"
)

func main() {

	srgb1 := rgb.SRGB{125, 246, 79}
	srgb2 := rgb.SRGB{56, 8, 86}

	xyz1, xyz2 := srgb1.XYZ(), srgb2.XYZ()
	fmt.Println(xyz1.Lab())
	fmt.Println(xyz2.Lab())

	fmt.Println(palette.Cie76(&xyz1, &xyz2))
	fmt.Println(palette.Cie94(&xyz1, &xyz2, palette.Textiles))
	fmt.Println(palette.Cie2000(&xyz1, &xyz2))

	//lab1, lab2 := cie.Lab{50, 2.6772, -79.7751}, cie.Lab{50, 0, -82.7485}
	//fmt.Println(palette.Cie2000(&lab1, &lab2))

}
