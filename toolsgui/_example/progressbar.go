package main

import (
	tools2D "proto_ui/toolsgui/wgt"

	"github.com/gonutz/prototype/draw"
)

var cp = 0

var pb1 = tools2D.NewProgressBar(1, 0, 0, 400, 50, 100, draw.White, draw.Green, draw.Blue, 1.2)
var pb2 = tools2D.NewProgressBar(2, 0, 60, 400, 50, 100, draw.Blue, draw.Red, draw.Red, 1.5)
var pb3 = tools2D.NewProgressBar(3, 0, 120, 400, 50, 100, draw.DarkGray, draw.Yellow, draw.Cyan, 1.6)
var pb4 = tools2D.NewProgressBar(4, 0, 180, 400, 50, 100, draw.Green, draw.Purple, draw.Yellow, 2.7)

func update(win draw.Window) {
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
	draw.RunWindow("toolsgui-tools2D-example-progressbar", 400, 240, update)
}
