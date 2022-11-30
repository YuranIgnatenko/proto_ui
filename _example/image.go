package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/yuranignatenko/proto_ui/wgt"
)

var im1 = wgt.NewImage("src/im1.png", 0, 0, 700, 300, 0)
var im2 = wgt.NewImage("src/im2.png", 0, 150, 700, 400, 30)
var im3 = wgt.NewImage("src/im1.png", 0, 300, 700, 300, 180)

func update(win draw.Window) {
	im1.View(win)
	im3.View(win)
	im2.View(win)
}

func main() {
	draw.RunWindow("toolsgui-wgt-example-image", 700, 600, update)
}
