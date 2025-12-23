package utils

import (
	"image/color"

	"tinygo.org/x/drivers"
	"tinygo.org/x/tinyfont"
)

func CenterText(display drivers.Displayer, f tinyfont.Fonter, x, y, w, h int16, str string, c color.RGBA) {
	_, lw := tinyfont.LineWidth(f, str)
	lh := int16(tinyfont.GetGlyph(f, rune(str[0])).Info().Height)
	lx := max(w-int16(lw), 0) / 2
	ly := max(h-lh, 0) / 2
	// 👇 y coord for text is baseline
	tinyfont.WriteLine(display, f, lx, ly+lh, str, c)
}
