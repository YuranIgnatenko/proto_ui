package main

import (
	"github.com/YuranIgnatenko/proto_ui/toolsgui/wgt"

	"github.com/gonutz/prototype/draw"
)

var ent1 = wgt.NewEntryRect("Entry", 10,
	draw.White, draw.DarkGray, draw.Black,
	draw.Black, draw.Cyan, 0, 0, 300, 50)

var ent2 = wgt.NewEntryRect("Entry", 10,
	draw.Blue, draw.DarkRed, draw.White,
	draw.Black, draw.Yellow, 0, 60, 300, 50)

var ent3 = wgt.NewEntryRect("Entry", 10,
	draw.Brown, draw.DarkCyan, draw.Blue,
	draw.White, draw.Red, 0, 120, 300, 50)

var info1 = wgt.NewLabel("Info label ", 0, 200, draw.Green, 1.7)
var info2 = wgt.NewLabel("Info label ", 0, 250, draw.Red, 1.7)
var info3 = wgt.NewLabel("Info label ", 0, 300, draw.Blue, 1.7)

func update(win draw.Window) {
	ent1.View(win)
	ent2.View(win)
	ent3.View(win)

	info1.Text = (ent1.GetText())
	info2.Text = (ent2.GetText())
	info3.Text = (ent3.GetText())

	info1.View(win)
	info2.View(win)
	info3.View(win)
}

func main() {
	draw.RunWindow("toolsgui-wgt-example-entry", 300, 400, update)
}
