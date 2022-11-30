package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/yuranignatenko/proto_ui/wgt"
)

var cp = 0

var pb1 = wgt.NewProgressBar(1, 0, 0, 400, 50, 100, draw.White, draw.Green, draw.Blue)
var pb2 = wgt.NewProgressBar(2, 0, 60, 400, 50, 100, draw.Blue, draw.Red, draw.Red)
var pb3 = wgt.NewProgressBar(3, 0, 120, 400, 50, 100, draw.DarkGray, draw.Yellow, draw.Cyan)
var pb4 = wgt.NewProgressBar(4, 0, 180, 400, 50, 100, draw.Green, draw.Purple, draw.Yellow)

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
	draw.RunWindow("toolsgui-wgt-example-progressbar", 400, 240, update)
}
