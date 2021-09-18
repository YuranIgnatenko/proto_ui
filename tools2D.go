package tools2D

import (
	. "fmt"
	"math"
	"strconv"

	"github.com/gonutz/prototype/draw"
)

type CardinateSurface2D struct {
	XDraw, YDraw, WDraw, HDraw int
	ColorLine                  draw.Color
	ColorLineXdiv              draw.Color
	ColorLineYdiv              draw.Color
	OneDivX, OneDivY           int
	SizeDiv                    int
	AllCountDivX               int
	AllCountDivY               int
	DistanceX                  int
	DistanceY                  int
}

func NewCardinateSurface2d(X, Y, W, H int, colorC, colorX, colorY draw.Color, lenXdiv, lenYdiv, size int) CardinateSurface2D {
	return CardinateSurface2D{
		XDraw:         X,
		YDraw:         Y,
		WDraw:         W,
		HDraw:         H,
		ColorLine:     colorC,
		ColorLineXdiv: colorX,
		ColorLineYdiv: colorY,
		OneDivX:       lenXdiv,
		OneDivY:       lenYdiv,
		SizeDiv:       size,
		AllCountDivX:  int((W-X)/lenXdiv) + 1,
		AllCountDivY:  int((H-Y)/lenYdiv) + 1,
		DistanceX:     int(W / int(W/lenXdiv)),
		DistanceY:     int(H / int(H/lenYdiv)),
	}

}

func (cs *CardinateSurface2D) DrawCardinateSurface2D(win draw.Window) {
	win.DrawLine(cs.XDraw, cs.YDraw, cs.XDraw, cs.HDraw+cs.YDraw, draw.White)
	win.DrawLine(cs.XDraw, cs.HDraw+cs.YDraw, cs.WDraw+cs.XDraw, cs.HDraw+cs.YDraw, draw.White)
}

func (cs *CardinateSurface2D) DrawXYDIVCardinateSurface2D(win draw.Window) {

	biasX := 0
	biasY := 0
	for i := 0; i <= cs.AllCountDivX; i++ {
		win.DrawLine(
			cs.XDraw+biasX,
			cs.HDraw+cs.YDraw,
			cs.XDraw+biasX,
			cs.HDraw+cs.SizeDiv+cs.YDraw,
			cs.ColorLineXdiv)
		biasX += cs.DistanceX
	}
	for i := 0; i <= cs.AllCountDivY; i++ {
		win.DrawLine(
			cs.XDraw-cs.SizeDiv,
			cs.HDraw-biasY+cs.YDraw,
			cs.XDraw,
			cs.HDraw-biasY+cs.YDraw,
			cs.ColorLineYdiv)
		biasY += cs.DistanceY
	}
}

func (cs *CardinateSurface2D) UpdateCardinateSurface2D(win draw.Window) {
	cs.DrawCardinateSurface2D(win)
	cs.DrawXYDIVCardinateSurface2D(win)
	cs.AddPointNill(10, 5, 6, draw.Red, win)
	cs.AddPointFill(3, 2, 5, draw.Yellow, win)
	cs.AddLineNill(0, 0, 3, 3, draw.Purple, win)
	cs.AddLinePoint(4, 6, 4, 1, 10, draw.White, draw.Blue, win)
}

func (cs *CardinateSurface2D) AddPointFill(X, Y, size int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := cs.HDraw - (Y * cs.DistanceY)
	win.FillEllipse(xn-int(size/2), yn-int(size/2), size, size, Color)

}

func (cs *CardinateSurface2D) AddPointNill(X, Y, size int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := cs.HDraw - (Y * cs.DistanceY)
	win.DrawEllipse(xn-int(size/2), yn-int(size/2), size, size, Color)

}

func (cs *CardinateSurface2D) AddLineNill(X, Y, x2, y2 int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := cs.HDraw - (Y * cs.DistanceY)
	xn2 := cs.XDraw + cs.DistanceX*x2
	yn2 := cs.HDraw - (y2 * cs.DistanceY)
	win.DrawLine(xn, yn, xn2, yn2, Color)
}

func (cs *CardinateSurface2D) AddLinePoint(X, Y, x2, y2, size int, ColorLine, colorPoint draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := cs.HDraw - (Y * cs.DistanceY)
	xn2 := cs.XDraw + cs.DistanceX*x2
	yn2 := cs.HDraw - (y2 * cs.DistanceY)
	win.FillEllipse(xn-int(size/2), yn-int(size/2), size, size, colorPoint)
	win.FillEllipse(xn2-int(size/2), yn2-int(size/2), size, size, colorPoint)
	win.DrawLine(xn, yn, xn2, yn2, ColorLine)

}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type Biffur struct {
	// R*X*(1-X)
	// rx(1-X)
	r int
}

func NewBiffur(r int) Biffur {
	return Biffur{
		r: r,
	}
}

func (b *Biffur) GetAlgoritm(x_population int) int {
	return int(b.r * x_population * (1 - x_population))
}

func (b *Biffur) GetLoop(count_population int) {
	Population := 4
	for i := 0; i < count_population; i++ {
		Population = b.GetAlgoritm(Population)
		Println("value population : ", Population)
	}
}

func (b *Biffur) GetLoopGUI(win draw.Window, count_population int) {
	Population := 4

	Population = b.GetAlgoritm(Population)

	// Println("value population : ", Population)

}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ButtonCircle struct {
	Text                                             string
	X, Y, W, H                                       int
	FontSize                                         float32
	FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color
	Command                                          func()
}

func NewButtonCircle(Text string, FgActive, FgDisable, FgFontActive, FgFontDisable draw.Color, command func(), X, Y, W, H int, FontSize float32) ButtonCircle {
	return ButtonCircle{
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

func (b *ButtonCircle) WaitPressButtonType1(window draw.Window) {
	centerX, centerY := b.W/2, b.H/2

	mouseX, mouseY := window.MousePosition()
	mouseInCircle := math.Hypot(float64(mouseX-centerX), float64(mouseY-centerY)) < 20
	colorBg := b.FgDisable
	colorFont := b.FgFontDisable

	if mouseInCircle {
		colorBg = b.FgActive
		colorFont = b.FgFontActive
	}

	window.FillEllipse(b.X/2, b.Y/2, b.W, b.H, colorBg)
	window.DrawScaledText(b.Text, b.X, b.Y, b.FontSize, colorFont)

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		dx, dy := click.X-centerX, click.Y-centerY
		squareDist := dx*dx + dy*dy
		if squareDist <= 20*20 {
			// close the window and end the application
			b.Command()
		}
	}
}

func (b *ButtonCircle) WaitPressButtonType2(window draw.Window) {
	W, H := window.Size()
	centerX, centerY := W/2, H/2

	// draw a button in the center of the screen
	mouseX, mouseY := window.MousePosition()
	mouseInCircle := math.Hypot(float64(mouseX-centerX), float64(mouseY-centerY)) < 20
	Color := draw.DarkRed
	if mouseInCircle {
		Color = draw.Red
	}

	if mouseInCircle {
		window.DrawScaledText(b.Text, centerX, centerY, 1.6, b.FgActive)
	} else {
		window.FillEllipse(centerX-20, centerY-20, 40, 40, Color)
		window.DrawEllipse(centerX-20, centerY-20, 40, 40, draw.White)

	}

	// check all mouse clicks that happened during this frame
	for _, click := range window.Clicks() {
		dx, dy := click.X-centerX, click.Y-centerY
		squareDist := dx*dx + dy*dy
		if squareDist <= 20*20 {
			// close the window and end the application
			b.Command()
		}
	}
}

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
}

func NewButtonImage(im1, im2, im3 string, FgFontActive, FgFontDisable draw.Color,
	command func(), X, Y, W, H int) ButtonImage {
	return ButtonImage{
		FgActive:  im1,
		FgDisable: im2,
		FgClick:   im3,
		Command:   command,
		X:         X,
		Y:         Y,
		W:         W,
		H:         H,
	}
}

func (b *ButtonImage) View(window draw.Window) {
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

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type Label struct {
	Text     string
	X, Y     int
	Color    draw.Color
	FontSize float32
}

func NewLabel(Text string, X, Y int, Color draw.Color, fontsize float32) Label {
	return Label{
		Text:     Text,
		X:        X,
		Y:        Y,
		FontSize: fontsize,
		Color:    Color,
	}
}

func (lb *Label) View(win draw.Window) {
	win.DrawScaledText(lb.Text, lb.X, lb.Y, lb.FontSize, lb.Color)
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ImagePic struct {
	Image           string
	X, Y, W, H, DEG int
}

func NewImage(image string, x, y, w, h, deg int) ImagePic {
	return ImagePic{
		Image: image,
		X:     x,
		Y:     y,
		W:     w,
		H:     h,
		DEG:   deg,
	}
}

func (im *ImagePic) View(win draw.Window) {
	win.DrawImageFileTo(im.Image, im.X, im.Y, im.W, im.H, im.DEG)
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ImagePicAnimation struct {
	Image           []string
	X, Y, W, H, DEG int
}

func NewImagePicAnimation(image []string, x, y, w, h int, deg int) ImagePicAnimation {
	return ImagePicAnimation{
		Image: image,
		X:     x,
		Y:     y,
		W:     w,
		H:     h,
		DEG:   deg,
	}
}

func (im *ImagePicAnimation) View(win draw.Window, i int) {
	Println(im.Image[i])
	win.DrawImageFileTo(im.Image[i], im.X, im.Y, im.W, im.H, im.DEG)
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type ProgressBar struct {
	TypeBar, X, Y, W, H, MaxPercent int
	Color1                          draw.Color
	Color2                          draw.Color
	Color3                          draw.Color
}

func NewProgressBar(t, X, Y, W, H, MaxPercent int, Color1, Color2, Color3 draw.Color) ProgressBar {
	return ProgressBar{
		TypeBar:    t,
		X:          X,
		Y:          Y,
		W:          W,
		H:          H,
		MaxPercent: MaxPercent,
		Color1:     Color1,
		Color2:     Color2,
		Color3:     Color3,
	}
}

func (pb *ProgressBar) View(percent int, win draw.Window) {
	border := 5
	lengthBar := pb.W
	percPx := int(lengthBar / 100)
	centerBar := lengthBar / 2

	if pb.TypeBar == 1 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
	}
	if pb.TypeBar == 2 {
		win.FillRect(pb.X, pb.Y, percent*percPx, pb.H, pb.Color2)
	}
	if pb.TypeBar == 3 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", centerBar+pb.X, pb.Y+int(pb.H/2), pb.Color3)
	}
	if pb.TypeBar == 4 {
		win.FillRect(pb.X, pb.Y, percent*percPx, pb.H, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", centerBar+pb.X, pb.Y+int(pb.H/2), pb.Color3)

	}
	if pb.TypeBar == 5 {
		win.DrawRect(pb.X, pb.Y, pb.W, pb.H, pb.Color1)
		win.FillRect(pb.X+border, pb.Y+border, percent*percPx-border*2, pb.H-border*2, pb.Color2)
		win.DrawText(strconv.Itoa(percent)+" % ", pb.X+pb.W+border, pb.Y+int(pb.H/2), pb.Color3)
	}
}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type Histogramma struct {
	Data           []int
	MaxValueColumn int
	X, Y, W, H     int
	Color          draw.Color
}

func NewHistogramma(data []int, X, Y, W, H, maxValCol int, Color draw.Color) Histogramma {
	return Histogramma{
		Data:           data,
		X:              X,
		Y:              Y,
		W:              W,
		H:              H,
		Color:          Color,
		MaxValueColumn: maxValCol,
	}
}

func (hist *Histogramma) View(win draw.Window, typeView int, fontSize float32, padding int, divLine int) {

	oneColumnW := int(hist.W / len(hist.Data))
	oneColumnH := int(hist.H / hist.MaxValueColumn)

	xn := hist.X
	biasX := 5
	biasYt1 := int(hist.MaxValueColumn / 10)
	bY := biasYt1
	biasYt2 := int(hist.MaxValueColumn / 4)
	bY2 := biasYt2
	if typeView == 1 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
	}
	if typeView == 2 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 10; i++ {
			win.DrawLine(hist.X-biasX, hist.H+hist.Y-oneColumnH*bY, hist.X, hist.H+hist.Y-oneColumnH*bY, hist.Color)
			bY += biasYt1

		}
	}
	if typeView == 3 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X-biasX, hist.H+hist.Y-oneColumnH*bY2, hist.X, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			bY2 += biasYt2

		}
	}
	if typeView == 4 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.DrawRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2, hist.X+hist.W+padding*len(hist.Data)+2+divLine, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			win.DrawScaledText(strconv.Itoa(i*biasYt2), hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2+2, fontSize, hist.Color)
			bY2 += biasYt2

		}
	}
	if typeView == 5 {
		win.DrawRect(hist.X, hist.Y, hist.W+padding*len(hist.Data), hist.H, hist.Color)
		for _, value := range hist.Data {
			win.FillRect(xn, (hist.H+hist.Y)-oneColumnH*value, oneColumnW, oneColumnH*value, hist.Color)
			xn += padding + oneColumnW
		}
		for i := 0; i < 4; i++ {
			win.DrawLine(hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2, hist.X+hist.W+padding*len(hist.Data)+2+divLine, hist.H+hist.Y-oneColumnH*bY2, hist.Color)
			win.DrawScaledText(strconv.Itoa(i*biasYt2), hist.X+hist.W+padding*len(hist.Data)+2, hist.H+hist.Y-oneColumnH*bY2+2, fontSize, hist.Color)
			bY2 += biasYt2

		}
	}

}

/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
/// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// var counter = 0
// var card = NewCardinateSurface2d(
// 	600, 20, 200, 200, draw.White, draw.White, draw.White, 20, 20, 5)

// func f() {
// 	Println("cLICK")
// }

// var btn = NewButtonRect("BTN1",
// 	draw.Green, draw.DarkGray,
// 	draw.Blue, draw.Red,
// 	f, 600, 600, 80, 80, 1.2)

// var lb = NewLabel("info label 1", 350, 50, draw.Yellow, 1.5)
// var pb1 = NewProgressBar(1, 850, 10, 200, 50, 100, draw.White, draw.Purple, draw.Green)
// var pb2 = NewProgressBar(2, 850, 100, 200, 50, 100, draw.White, draw.Purple, draw.Green)
// var pb3 = NewProgressBar(3, 850, 200, 200, 50, 100, draw.White, draw.Purple, draw.Green)
// var pb4 = NewProgressBar(4, 850, 300, 200, 50, 100, draw.White, draw.Purple, draw.Green)
// var pb5 = NewProgressBar(5, 850, 400, 200, 50, 100, draw.Gray, draw.DarkCyan, draw.White)
// var i = 0
// var hist = NewHistogramma([]int{10, 50, 60, 80, 90, 20, 70, 43, 20, 20, 54}, 850, 500, 100, 100, 100, draw.White)

// func update(win draw.Window) {
// 	card.UpdateCardinateSurface2D(win)

// 	btn.WaitPressButtonType3(win)
// 	lb.Text = strconv.Itoa(counter) + "  ITER\n"
// 	lb.View(win)

// 	pb1.View(i, win)
// 	pb2.View(i, win)
// 	pb3.View(i, win)
// 	pb4.View(i, win)
// 	pb5.View(i, win)
// 	counter++
// 	i++

// 	if i > 100 {
// 		i = 0
// 	}
// 	hist.View(win)
// }

// func main() {
// 	draw.RunWindow("2d tools", 1200, 700, update)
// }
