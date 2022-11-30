package main

import (
	"strconv"

	"github.com/gonutz/prototype/draw"
	"github.com/yuranignatenko/proto_ui/wgt"
)

var fract = wgt.FractalMath{
	X:                 0,
	Y:                 0,
	ArrayPointsFigure: []int{50, 50, 600, 50, 600, 50},
	Size:              3}

var counter = 0

func fcmdbt() {
	fract.ArrayUsedPoint = make([][]int, 0)
	counter = 0
}

var btn = wgt.NewButtonRect(
	"RESET\nDRAW",
	draw.Green, draw.DarkGray,
	draw.Blue, draw.Red, fcmdbt,
	250, 500, 80, 40, 1.2)

var lb = wgt.NewLabel(" ", 250, 550, draw.Green, 2.)

func update(win draw.Window) {
	fract.RunFractGUI(win, "xy/2", draw.LightCyan)
	btn.WaitPressButtonType3(win)
	lb.Text = strconv.Itoa(counter) + "  ITER\n"
	lb.View(win)
	counter++
}

func main() {
	draw.RunWindow("toolsgui-wgt-examplefractal", 1000, 650, update)
}
