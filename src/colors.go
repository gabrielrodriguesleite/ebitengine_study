package game

import "image/color"

var (
	backGroundColor = color.RGBA{0xfa, 0xf8, 0xef, 0xff}
	frameColor      = color.RGBA{0xbb, 0xad, 0xa0, 0xff}
)

func titleBackgroundColor(value int) color.Color {
	switch value {
	case 0:
		return color.NRGBA{0xee, 0xe4, 0xda, 0x59}
	case 2:
		return color.RGBA{0xee, 0xe4, 0xda, 0xff}
	case 4:
		return color.RGBA{0xed, 0xe0, 0xc8, 0xff}
	}
	panic("tile value not reach")
}
