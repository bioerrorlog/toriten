package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	X, Y          int
	Width, Height int
	Color         color.RGBA
	ColorPressed  color.RGBA
	ButtonText    string
	fontFace      font.Face
	textXPadding  int
	textYPadding  int
	IsPressed     bool
	IsClicked     bool // Clicked: the button was pressed in the last frame and is no longer pressed
	isLastPressed bool
}

func NewButton(x, y, width, height int, buttonText string, fontFace font.Face, textXPadding int, textYPadding int, color color.RGBA, colorPressed color.RGBA) (*Button, error) {
	b := &Button{
		X:             x,
		Y:             y,
		Width:         width,
		Height:        height,
		ButtonText:    buttonText,
		fontFace:      fontFace,
		textXPadding:  textXPadding,
		textYPadding:  textYPadding,
		Color:         color,
		ColorPressed:  colorPressed,
		isLastPressed: false,
	}
	return b, nil
}

func (b *Button) Update() error {
	x, y := ebiten.CursorPosition()

	b.IsPressed = x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	// Clicked: the button was pressed in the last frame and is no longer pressed
	if b.isLastPressed && !b.IsPressed {
		b.IsClicked = true
	} else {
		b.IsClicked = false
	}

	// Save the current pressed status for the next frame
	b.isLastPressed = b.IsPressed

	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	fillColor := b.Color
	if b.IsPressed {
		fillColor = b.ColorPressed
	}

	buttonImage := ebiten.NewImage(b.Width, b.Height)
	buttonImage.Fill(fillColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(buttonImage, op)

	// Draw the button text
	textX := b.X + b.textXPadding
	textY := b.Y + b.textYPadding
	text.Draw(screen, b.ButtonText, b.fontFace, textX, textY, color.White)
}
