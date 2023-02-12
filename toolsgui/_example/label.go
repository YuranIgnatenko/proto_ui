package main

import (
	"strconv"

	"github.com/gonutz/prototype/draw"
	tools2D "github.com/YuranIgnatenko/proto_ui/toolsgui/wgt"
)

var ci = 0

var info1 = tools2D.NewLabel("Info label ", 0, 0, draw.Green, 2.5)
var info2 = tools2D.NewLabel("Info label ", 0, 100, draw.Red, 1.7)
var info3 = tools2D.NewLabel("Info label ", 200, 100, draw.Blue, 1.7)
var info4 = tools2D.NewLabel("Info label ", 40, 200, draw.Purple, 2.5)
var info5 = tools2D.NewLabel("Info label ", 200, 200, draw.Yellow, 1.)

func update(win draw.Window) {
	ci++
	info1.View(win)
	info2.View(win)
	info3.Text = strconv.Itoa(ci * 2)
	info3.View(win)
	info4.Text = strconv.Itoa(ci)
	info4.View(win)
	info5.Text = strconv.Itoa(ci + -ci*-3)
	info5.View(win)
}

func main() {
	draw.RunWindow("toolsgui-tools2D-example-label", 300, 300, update)
}
