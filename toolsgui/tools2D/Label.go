package tools2D

import "github.com/gonutz/prototype/draw"

type Label struct {
	Text     string
	X, Y     int
	Color    draw.Color
	FontSize float32
}

func NewLabel(Text string, X, Y int, Color draw.Color, fontsize float32) Label {
	return Label{
		Text:     Text,
		X:        X,
		Y:        Y,
		FontSize: fontsize,
		Color:    Color,
	}
}

func (lb *Label) View(win draw.Window) {
	win.DrawScaledText(lb.Text, lb.X, lb.Y, lb.FontSize, lb.Color)
}
