package main

import (
	"github.com/YuranIgnatenko/proto_ui/toolsgui/wgt"

	"github.com/gonutz/prototype/draw"
)

var coord1 = wgt.NewCoordinateSurface2d(50, 50, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord2 = wgt.NewCoordinateSurface2d(300, 50, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord3 = wgt.NewCoordinateSurface2d(550, 50, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord4 = wgt.NewCoordinateSurface2d(70, 200, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord5 = wgt.NewCoordinateSurface2d(300, 200, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)
var coord6 = wgt.NewCoordinateSurface2d(550, 200, 200, 100, draw.DarkBlue, draw.Green, draw.Red, 20, 30, 10)

var cx = 0

func update(win draw.Window) {
	coord1.AddPointNill(1, 1, 5, draw.White, win)
	coord2.AddPointNill(1, 2, 2, draw.Green, win)
	coord1.View(win)
	coord2.View(win)

	coord3.AddLinePoint(1, 1, 2, 3, 7, draw.Green, draw.Green, win)
	coord3.AddLinePoint(2, 3, 4, 2, 7, draw.Green, draw.Green, win)
	coord3.View(win)

	coord4.AddLineNill(1, 1, 2, 3, draw.Blue, win)
	coord4.AddLineNill(2, 3, 4, 2, draw.Blue, win)
	coord4.View(win)

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

}

func main() {
	draw.RunWindow("toolgui-tool2D-example-coordinate", 900, 400, update)
}
