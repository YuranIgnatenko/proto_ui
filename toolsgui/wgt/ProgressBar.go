package wgt

import (
	"strconv"

	"github.com/gonutz/prototype/draw"
)

type ProgressBar struct {
	TypeBar, X, Y, W, H, MaxPercent int
	Color1                          draw.Color
	Color2                          draw.Color
	Color3                          draw.Color
}

func NewProgressBar(typeBar, X, Y, W, H, MaxPercent int, Color1, Color2, Color3 draw.Color) ProgressBar {
	return ProgressBar{
		TypeBar:    typeBar,
		X:          X,
		Y:          Y,
		W:          W,
		H:          H,
		MaxPercent: MaxPercent,
		Color1:     Color1,
		Color2:     Color2,
		Color3:     Color3,
	}
}

func (pb *ProgressBar) View(percent int, win draw.Window) {
	border := 5
	lengthBar := pb.W
	percPx := int(lengthBar / 100)
	centerBar := lengthBar / 2

	if pb.TypeBar == 1 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
	}
	if pb.TypeBar == 2 {
		win.FillRect(pb.X, pb.Y, percent*percPx, pb.H, pb.Color2)
	}
	if pb.TypeBar == 3 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", centerBar+pb.X, pb.Y+int(pb.H/2), pb.Color3)
	}
	if pb.TypeBar == 4 {
		win.FillRect(pb.X, pb.Y, percent*percPx, pb.H, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", centerBar+pb.X, pb.Y+int(pb.H/2), pb.Color3)

	}
	if pb.TypeBar == 5 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", pb.X+pb.W+border, pb.Y+int(pb.H/2), pb.Color3)
	}
}
