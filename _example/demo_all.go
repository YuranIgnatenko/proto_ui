package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/yuranignatenko/proto_ui/wgt"
)

func cmd1() {

}

var btn1 = wgt.NewButtonRect(
	"btn1",
	draw.Purple, draw.Red, draw.DarkBlue, draw.Green,
	cmd1, 10, 10, 80, 20, 1.0,
)

var coord5 = wgt.NewCoordinateSurface2d(30, 40, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord6 = wgt.NewCoordinateSurface2d(30, 200, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)

var ent1 = wgt.NewEntryRect("Entry", 10,
	draw.White, draw.DarkGray, draw.Black,
	draw.Black, draw.Cyan, 10, 340, 300, 50)

var ent2 = wgt.NewEntryRect("Entry", 10,
	draw.Blue, draw.DarkRed, draw.White,
	draw.Black, draw.Yellow, 10, 420, 300, 50)

var info1 = wgt.NewLabel("Info label ", 10, 480, draw.Green, 1.7)
var info2 = wgt.NewLabel("Info label ", 10, 520, draw.Red, 1.7)

// var fract = wgt.FractalMath{
// 	X:                 0,
// 	Y:                 0,
// 	ArrayPointsFigure: []int{},
// 	ArrayUsedPoint:    [][]int{},
// 	Size:              4}

var dt = []int{0, 10, 9, 2, 8, 7, 11, 15, 17}
var data = []float64{0.3, 0.5, 0.9, 1., 4., 9., 9.9}
var data2 = []int{1, 3, 7, 7, 9, 8, 5, 20, 4, 70, 76, 33, 35, 9, 90}

var hist1 = wgt.NewHistogramma(dt, 300, 50, 200, 100, 100, draw.White)
var hist2 = wgt.NewHistogramma(data2, 300, 200, 200, 100, 100, draw.Green)

var hist3 = wgt.NewHistogrammaFloat(data, 400, 320, 200, 100, 100, draw.Blue)
var hist4 = wgt.NewHistogrammaFloat(data, 650, 320, 100, 100, 100, draw.Yellow)

var hist5 = wgt.NewHistogramma(data2, 400, 450, 100, 100, 100, draw.Red)

var im2 = wgt.NewImage("src/im2.png", 500, 0, 200, 200, 30)
var im1 = wgt.NewImage("src/im1.png", 400, 350, 400, 200, 0)

var cx = 0
var c = 0
var cp = 0

var pb1 = wgt.NewProgressBar(1, 650, 0, 550, 50, 100, draw.White, draw.Green, draw.Blue, 0.9)
var pb2 = wgt.NewProgressBar(2, 650, 60, 500, 40, 100, draw.Blue, draw.Red, draw.Red, 0.9)
var pb3 = wgt.NewProgressBar(3, 650, 120, 450, 30, 100, draw.DarkGray, draw.Yellow, draw.Cyan, 0.9)
var pb4 = wgt.NewProgressBar(4, 650, 180, 300, 20, 100, draw.Green, draw.Purple, draw.Yellow, 0.9)

func update(win draw.Window) {
	im1.View(win)
	im2.View(win)

	btn1.WaitPressButtonType1(win)

	if cx < 10 {
		cx++
	} else {
		cx = 0
	}

	coord5.AddLineNill(cx+1, 1, cx+2, 3, draw.Blue, win)
	coord5.AddLineNill(cx+2, 3, cx+4, 2, draw.Blue, win)
	coord5.View(win)

	coord6.AddPointNill(cx, 1, 5, draw.White, win)
	coord6.AddPointNill(cx, 2, 2, draw.Green, win)
	coord6.View(win)

	ent1.View(win)
	ent2.View(win)

	info1.Text = (ent1.GetText())
	info2.Text = (ent2.GetText())

	info1.View(win)
	info2.View(win)

	// fract.RunFractGUI(win, "xy/2", draw.LightCyan)
	if c < 80 {
		c++
		oldData := hist1.Data
		newData := make([]int, 0)
		for _, val := range oldData {
			newData = append(newData, val+1)
			hist1.Data = newData
		}
	} else {
		hist1.Data = dt
		c = 0
	}
	hist1.View(win, 1, 1, 0, 15)
	hist2.View(win, 2, 1, 1, 15)
	hist3.View(win, 3, 1, 5, 15)
	hist4.View(win, 4, 1, 15, 15)
	hist5.View(win, 5, 1, 35, 25)

	if cp == 100 {
		cp = 0
	} else {
		cp++
	}

	pb1.View(cp, win)
	pb2.View(cp, win)
	pb3.View(cp, win)
	pb4.View(cp, win)

}

func main() {
	// fract.LoadArrayPoints([]int{500, 10, 800, 10, 800, 800})
	draw.RunWindow("toolsgui-wgt-example-button", 1200, 600, update)
}
