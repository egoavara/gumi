package svg

import (
	"image/color"
	"regexp"
	"strconv"
	"math"
)

var (
	hexShortRGB  = regexp.MustCompile(`^#[A-Fa-f0-9]{3}$`)
	hexLongRGB   = regexp.MustCompile(`^#[A-Fa-f0-9]{6}$`)
	hexShortRGBA = regexp.MustCompile(`^#[A-Fa-f0-9]{4}$`)
	hexLongRGBA  = regexp.MustCompile(`^#[A-Fa-f0-9]{8}$`)
)

func HexColor(hex string) color.Color {
	var r, g, b, a uint8 = 0,0,0,math.MaxUint8
	if hexShortRGB.MatchString(hex) {
		r = mustParseUint8Hex(hex[1:2])
		g = mustParseUint8Hex(hex[2:3])
		b = mustParseUint8Hex(hex[3:4])
		return color.RGBA{r,g,b,a,}
	}
	if hexLongRGB.MatchString(hex) {
		r = mustParseUint8Hex(hex[1:3])
		g = mustParseUint8Hex(hex[3:5])
		b = mustParseUint8Hex(hex[5:7])
		return color.RGBA{r,g,b,a,}
	}
	if hexShortRGBA.MatchString(hex) {
		r = mustParseUint8Hex(hex[1:2])
		g = mustParseUint8Hex(hex[2:3])
		b = mustParseUint8Hex(hex[3:4])
		a = mustParseUint8Hex(hex[4:5])
		return color.RGBA{r,g,b,a,}
	}
	if hexLongRGBA.MatchString(hex) {
		r = mustParseUint8Hex(hex[1:3])
		g = mustParseUint8Hex(hex[3:5])
		b = mustParseUint8Hex(hex[5:7])
		a = mustParseUint8Hex(hex[7:9])
		return color.RGBA{r,g,b,a,}
	}
	return nil
}
func mustParseUint8Hex(string string) uint8 {
	temp, _ := strconv.ParseUint(string[1:2], 16, 8)

	return uint8(temp)
}
