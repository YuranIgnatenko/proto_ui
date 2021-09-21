package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/toolsgui/tools2D"
)

var im1 = tools2D.NewImage("src/im1.png", 0, 0, 700, 300, 0)
var im2 = tools2D.NewImage("src/im2.png", 0, 150, 700, 400, 30)
var im3 = tools2D.NewImage("src/im1.png", 0, 300, 700, 300, 180)

func update(win draw.Window) {
	im1.View(win)
	im3.View(win)
	im2.View(win)
}

func main() {
	draw.RunWindow("toolsgui-tools2D-example-image", 700, 600, update)
}
