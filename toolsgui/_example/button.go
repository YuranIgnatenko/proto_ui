package main

import (
	"fmt"
	"github.com/YuranIgnatenko/proto_ui/toolsgui/wgt"

	"github.com/gonutz/prototype/draw"
)

var btn1 = wgt.NewButtonRect(
	"btn1",
	draw.Purple, draw.Red, draw.DarkBlue, draw.Green,
	cmd1, 50, 50, 100, 50, 2.2,
)

var btn2 = wgt.NewButtonRect(
	"btn2",
	draw.White, draw.DarkGreen, draw.Blue, draw.Black,
	cmd1, 50, 150, 100, 50, 2.2,
)

var btn3 = wgt.NewButtonRect(
	"btn3",
	draw.White, draw.DarkYellow, draw.Green, draw.Gray,
	cmd1, 50, 250, 100, 50, 2.2,
)

var btn4 = wgt.NewButtonImage(
	"src/btn_on.png", "src/btn_off.png", "src/btn_wait.png",
	cmd1, 200, 50, 100, 50,
)

var btn5 = wgt.NewButtonImage(
	"src/btn_on.png", "src/btn_off.png", "src/btn_wait.png",
	cmd1, 200, 150, 100, 50,
)

var btn6 = wgt.NewButtonImage(
	"src/btn2_on.png", "src/btn2_off.png", "src/btn2_wait.png",
	cmd1, 200, 250, 100, 50,
)

func cmd1() {
	fmt.Println("Press button test")
}

func update(win draw.Window) {
	btn1.WaitPressButtonType1(win)
	btn2.WaitPressButtonType2(win)
	btn3.WaitPressButtonType3(win)
	btn4.WaitPressButtonType1(win)
	btn5.WaitPressButtonType2(win)
	btn6.WaitPressButtonType2(win)
}

func main() {
	draw.RunWindow("toolsgui-wgt-example-button", 400, 400, update)
}
