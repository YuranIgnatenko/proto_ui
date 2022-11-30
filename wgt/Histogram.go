package wgt

import (
	"strconv"

	"github.com/gonutz/prototype/draw"
)

type Histogramma struct {
	Data           []int
	MaxValueColumn int
	X, Y, W, H     int
	Color          draw.Color
}

func NewHistogramma(data []int, X, Y, W, H, maxValCol int, Color draw.Color) Histogramma {
	return Histogramma{
		Data:           data,
		X:              X,
		Y:              Y,
		W:              W,
		H:              H,
		Color:          Color,
		MaxValueColumn: maxValCol,
	}
}

func NewHistogrammaFloat(data []float64, X, Y, W, H, maxValCol int, Color draw.Color) Histogramma {
	newDt := make([]int, 0)
	for i := 0; i < len(data); i++ {
		// Println(int(data[i] * 10))
		newDt = append(newDt, int(data[i]*10))
	}
	return Histogramma{
		Data:           newDt,
		X:              X,
		Y:              Y,
		W:              W,
		H:              H,
		Color:          Color,
		MaxValueColumn: maxValCol,
	}
}

func (hist *Histogramma) View(win draw.Window, typeView int, fontSize float32, padding int, divLine int) {

	oneColumnW := int(hist.W / len(hist.Data))
	oneColumnH := int(hist.H / hist.MaxValueColumn)

	xn := hist.X
	biasX := 5
	biasYt1 := int(hist.MaxValueColumn / 10)
	bY := biasYt1
	biasYt2 := int(hist.MaxValueColumn / 4)
	bY2 := biasYt2
	if typeView == 1 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
	}
	if typeView == 2 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 10; i++ {
			win.DrawLine(hist.X-biasX, hist.H+hist.Y-oneColumnH*bY, hist.X, hist.H+hist.Y-oneColumnH*bY, hist.Color)
			bY += biasYt1

		}
	}
	if typeView == 3 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X-biasX, hist.H+hist.Y-oneColumnH*bY2, hist.X, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			bY2 += biasYt2

		}
	}
	if typeView == 4 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2, hist.X+hist.W+padding*len(hist.Data)+2+divLine, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			win.DrawScaledText(strconv.Itoa(i*biasYt2), hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2+2, fontSize, hist.Color)
			bY2 += biasYt2

		}
	}
	if typeView == 5 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.FillRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2, hist.X+hist.W+padding*len(hist.Data)+2+divLine, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			win.DrawScaledText(strconv.Itoa(i*biasYt2), hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2+2, fontSize, hist.Color)
			bY2 += biasYt2

		}
	}

}
