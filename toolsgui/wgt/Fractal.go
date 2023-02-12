package wgt

import (
	"math/rand"

	"github.com/gonutz/prototype/draw"
)

type FractalMath struct {
	ArrayPointsFigure []int
	ArrayUsedPoint    [][]int
	X                 int
	Y                 int
	Size              int
}

func (fm *FractalMath) LoadPointsStart(X int, Y int) {
	fm.X = X
	fm.Y = Y
}

func (fm *FractalMath) LoadArrayPoints(arr []int) {
	fm.ArrayPointsFigure = arr
}

func (fm *FractalMath) SelectPoint(arrPoint []int) []int {
	lenArr := len(arrPoint) / 2

	newPointIndex := rand.Intn(lenArr - 0)

	newX := arrPoint[newPointIndex]
	newY := arrPoint[newPointIndex+1]

	arXY := []int{newX, newY}
	return arXY
}

func (fm *FractalMath) Clear() {
	fm.X = 0
	fm.Y = 0
	fm.ArrayPointsFigure = []int{}

}

func (fm *FractalMath) NewXYdefault(X int, Y int) []int {
	newX := (X-fm.X)/2 + fm.X
	newY := (Y-fm.Y)/2 + fm.Y
	fm.X = int(newX)
	fm.Y = int(newY)
	return []int{newX, newY}
}

func (fm *FractalMath) NewXYdefault4(X int, Y int) []int {
	newX := (X-fm.X)/4 + fm.X
	newY := (Y-fm.Y)/4 + fm.Y
	fm.X = int(newX)
	fm.Y = int(newY)
	return []int{newX, newY}
}

func (fm *FractalMath) RunFract(countIter int, selectAlgoritm string) {
	for i := 0; i < countIter; i++ {
		arxy := fm.SelectPoint(fm.ArrayPointsFigure)
		switch selectAlgoritm {
		case "xy/2":
			fm.NewXYdefault(arxy[0], arxy[1])
		case "xy/4":
			fm.NewXYdefault4(arxy[0], arxy[1])
		}
	}
}

func (fm *FractalMath) RunFractGUI(win draw.Window, selectAlgoritm string, color draw.Color) {
	arxy := fm.SelectPoint(fm.ArrayPointsFigure)
	switch selectAlgoritm {
	case "xy/2":
		arxy = fm.NewXYdefault(arxy[0], arxy[1])
		fm.X, fm.Y = int(arxy[0]), int(arxy[1])
		fm.ArrayUsedPoint = append(fm.ArrayUsedPoint, arxy)
		for i, val := range fm.ArrayUsedPoint {
			if i%2 != 0 {
				win.DrawEllipse(val[0], val[1], fm.Size, fm.Size, color)
			}
		}
	}
}
