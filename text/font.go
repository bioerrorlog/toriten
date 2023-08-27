package text

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFont(path string, size float64, dpi float64) (font.Face, error) {
	ttfBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ttf, err := opentype.Parse(ttfBytes)
	if err != nil {
		return nil, err
	}
	fontFace, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}
	return fontFace, nil
}
