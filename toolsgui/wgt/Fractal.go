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
	// fmt.Println("NEW INDEX FOR POINTS := ", newPointIndex)

	newX := arrPoint[newPointIndex]
	newY := arrPoint[newPointIndex+1]
	// fm.X = newX
	// fm.Y = newY
	// fmt.Println(newX, newY)
	arXY := []int{newX, newY}
	return arXY
}

func (fm *FractalMath) Clear() {
	// fmt.Println("1")
	fm.X = 0
	fm.Y = 0
	fm.ArrayPointsFigure = []int{}

}

func (fm *FractalMath) NewXYdefault(X int, Y int) []int {
	// fmt.Println("DATA :: [ X = ", X, ", Y = ", Y, " ]")

	newX := (X-fm.X)/2 + fm.X
	newY := (Y-fm.Y)/2 + fm.Y
	fm.X = int(newX)
	fm.Y = int(newY)
	// fmt.Println("NEW X,Y := ", fm.X, fm.Y)
	return []int{newX, newY}
}

func (fm *FractalMath) NewXYdefault4(X int, Y int) []int {
	// fmt.Println("DATA :: [ X = ", X, ", Y = ", Y, " ]")

	newX := (X-fm.X)/4 + fm.X
	newY := (Y-fm.Y)/4 + fm.Y
	fm.X = int(newX)
	fm.Y = int(newY)
	// fmt.Println("NEW X,Y := ", fm.X, fm.Y)
	return []int{newX, newY}
}

func (fm *FractalMath) RunFract(countIter int, selectAlgoritm string) {
	// fmt.Println("\nRUNNING ....\n")
	for i := 0; i < countIter; i++ {
		arxy := fm.SelectPoint(fm.ArrayPointsFigure)
		// print(arxy)
		switch selectAlgoritm {
		case "xy/2":
			arxy = fm.NewXYdefault(arxy[0], arxy[1])
			// print(arxy)
		case "xy/4":
			fm.NewXYdefault4(arxy[0], arxy[1])
		}

		// fmt.Println("\n\n")
	}
}

func (fm *FractalMath) RunFractGUI(win draw.Window, selectAlgoritm string, color draw.Color) {
	// fmt.Println("\nRUNNING ....\n")
	arxy := fm.SelectPoint(fm.ArrayPointsFigure)
	// print(arxy)
	switch selectAlgoritm {
	case "xy/2":
		arxy = fm.NewXYdefault(arxy[0], arxy[1])
		fm.X, fm.Y = int(arxy[0]), int(arxy[1])
		// win.DrawPoint(fm.X, fm.Y, draw.White)
		fm.ArrayUsedPoint = append(fm.ArrayUsedPoint, arxy)
		for _, val := range fm.ArrayUsedPoint {
			win.DrawEllipse(val[0], val[1], fm.Size, fm.Size, color)
		}
	}
}
