package tools2D

import (
	"github.com/gonutz/prototype/draw"
)

// type ButtonCircle struct {
// 	Text                                             string
// 	X, Y, W, H                                       int
// 	FontSize                                         float32
// 	FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color
// 	Command                                          func()
// }

// func NewButtonCircle(Text string, FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color, command func(), X, Y, W, H int, FontSize float32) ButtonCircle {
// 	return ButtonCircle{
// 		Text:          Text,
// 		FgActive:      FgActive,
// 		FgDisable:     FgDisable,
// 		FgFontActive:  FgFontActive,
// 		FgFontDisable: FgFontDisable,
// 		Command:       command,
// 		X:             X,
// 		Y:             Y,
// 		W:             W,
// 		H:             H,
// 		FontSize:      FontSize,
// 	}
// }

// func (b *ButtonCircle) WaitPressButtonType1(window draw.Window) {
// 	centerX, centerY := b.W/2, b.H/2

// 	mouseX, mouseY := window.MousePosition()
// 	mouseInCircle := math.Hypot(float64(mouseX-centerX), float64(mouseY-centerY)) < 20
// 	colorBg := b.FgDisable
// 	colorFont := b.FgFontDisable

// 	if mouseInCircle {
// 		colorBg = b.FgActive
// 		colorFont = b.FgFontActive
// 	}

// 	window.FillEllipse(b.X/2, b.Y/2, b.W, b.H, colorBg)
// 	window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)

// 	// check all mouse clicks that happened during this frame
// 	for _, click := range window.Clicks() {
// 		dx, dy := click.X-centerX, click.Y-centerY
// 		squareDist := dx*dx + dy*dy
// 		if squareDist <= 20*20 {
// 			// close the window and end the application
// 			b.Command()
// 		}
// 	}
// }

// func (b *ButtonCircle) WaitPressButtonType2(window draw.Window) {
// 	centerX, centerY := b.W/2, b.H/2

// 	// draw a button in the center of the screen
// 	mouseX, mouseY := window.MousePosition()
// 	mouseInCircle := math.Hypot(float64(mouseX-centerX), float64(mouseY-centerY)) < 20

// 	if mouseInCircle {
// 		window.FillEllipse(centerX-20, centerY-20, b.W, b.H, b.FgActive)
// 		window.DrawScaledText(b.Text, centerX, centerY, 1.6, b.FgFontActive)
// 	} else {
// 		window.DrawEllipse(centerX-20, centerY-20, b.W, b.H, b.FgDisable)
// 		window.DrawScaledText(b.Text, centerX, centerY, 1.6, b.FgFontDisable)
// 	}

// 	// check all mouse clicks that happened during this frame
// 	for _, click := range window.Clicks() {
// 		dx, dy := click.X-centerX, click.Y-centerY
// 		squareDist := dx*dx + dy*dy
// 		if squareDist <= 20*20 {
// 			// close the window and end the application
// 			b.Command()
// 		}
// 	}
// }

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ButtonRect struct {
	Text                                             string
	X, Y, W, H                                       int
	FontSize                                         float32
	FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color
	Command                                          func()
}

func NewButtonRect(Text string, FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color, command func(), X, Y, W, H int, FontSize float32) ButtonRect {
	return ButtonRect{
		Text:          Text,
		FgActive:      FgActive,
		FgDisable:     FgDisable,
		FgFontActive:  FgFontActive,
		FgFontDisable: FgFontDisable,
		Command:       command,
		X:             X,
		Y:             Y,
		W:             W,
		H:             H,
		FontSize:      FontSize,
	}
}

func (b *ButtonRect) WaitPressButtonType1(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2

	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	colorBg := b.FgDisable
	colorFont := b.FgFontDisable

	if mouseInRect {
		colorBg = b.FgActive
		colorFont = b.FgFontActive
	}

	window.FillRect(b.X, b.Y, b.W, b.H, colorBg)
	window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		if clickInRect {
			b.Command()
		}
	}
}

func (b *ButtonRect) WaitPressButtonType2(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2

	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	colorBg := b.FgDisable
	colorFont := b.FgFontDisable

	if mouseInRect {
		colorBg = b.FgActive
		colorFont = b.FgFontActive
		window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)
	} else {
		window.FillRect(b.X, b.Y, b.W, b.H, colorBg)

	}

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		if clickInRect {
			b.Command()
		}
	}
}

func (b *ButtonRect) WaitPressButtonType3(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2

	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	colorBg := b.FgDisable
	colorFont := b.FgFontDisable

	if mouseInRect {
		colorBg = b.FgActive
		colorFont = b.FgFontActive

		window.DrawRect(b.X, b.Y, b.W, b.H, colorBg)
		window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)
	} else {
		window.DrawRect(b.X, b.Y, b.W, b.H, colorBg)
		window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)
	}

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		if clickInRect {
			b.Command()
		}
	}
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ButtonImageText struct {
	Text                         string
	X, Y, W, H                   int
	FontSize                     float32
	FgActive, FgDisable, FgClick string
	FgFontActive, FgFontDisable  draw.Color
	Command                      func()
}

func NewButtonImageText(Text, im1, im2, im3 string, FgFontActive, FgFontDisable draw.Color,
	command func(), X, Y, W, H int, FontSize float32) ButtonImageText {
	return ButtonImageText{
		Text:          Text,
		FgActive:      im1,
		FgDisable:     im2,
		FgClick:       im3,
		FgFontActive:  FgFontActive,
		FgFontDisable: FgFontDisable,
		Command:       command,
		X:             X,
		Y:             Y,
		W:             W,
		H:             H,
		FontSize:      FontSize,
	}
}

func (b *ButtonImageText) View(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2
	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	img := b.FgDisable
	colorFont := b.FgFontDisable

	if mouseInRect {
		img = b.FgActive
		colorFont = b.FgFontActive
	}
	window.DrawImageFileTo(img, b.X, b.Y, b.W, b.H, 0)
	window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)
		if clickInRect {
			window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)

			b.Command()
		}
	}
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
type ButtonImage struct {
	X, Y, W, H                   int
	FontSize                     float32
	FgActive, FgDisable, FgClick string
	FgFontActive, FgFontDisable  draw.Color
	Command                      func()
	isPress                      bool
	counterClick                 int
}

func NewButtonImage(im1, im2, im3 string, command func(), X, Y, W, H int) ButtonImage {
	return ButtonImage{
		FgActive:     im1,
		FgDisable:    im2,
		FgClick:      im3,
		Command:      command,
		X:            X,
		Y:            Y,
		W:            W,
		H:            H,
		isPress:      false,
		counterClick: 0,
	}
}

func (b *ButtonImage) WaitPressButtonType1(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2
	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	img := b.FgDisable

	if mouseInRect {
		img = b.FgActive
	}
	window.DrawImageFileTo(img, b.X, b.Y, b.W, b.H, 0)

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)
		if clickInRect {
			window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)
			b.Command()
		}
	}
}

func (b *ButtonImage) WaitPressButtonType2(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2
	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)
	img := b.FgDisable

	if mouseInRect {
		img = b.FgClick
	}
	window.DrawImageFileTo(img, b.X, b.Y, b.W, b.H, 0)

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)
		if clickInRect {
			window.DrawImageFileTo(b.FgClick, b.X, b.Y, b.W, b.H, 0)

			b.Command()
			b.counterClick = 1
		}
	}

	if b.counterClick == 1 {
		b.FgActive, b.FgDisable = b.FgDisable, b.FgActive
		b.counterClick = 2
	}

}
