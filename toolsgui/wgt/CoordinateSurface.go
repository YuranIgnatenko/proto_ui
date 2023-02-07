package wgt

import "github.com/gonutz/prototype/draw"

type CoordinateSurface2D struct {
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

func NewCoordinateSurface2d(X, Y, W, H int, colorC, colorX, colorY draw.Color, lenXdiv, lenYdiv, size int) CoordinateSurface2D {
	return CoordinateSurface2D{
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
		AllCountDivX:  int((W) / lenXdiv),
		AllCountDivY:  int((H) / lenYdiv),
		DistanceX:     int(W / int(W/lenXdiv)),
		DistanceY:     int(H / int(H/lenYdiv)),
	}

}

func (cs *CoordinateSurface2D) DrawCardinateSurface2D(win draw.Window) {
	win.DrawLine(cs.XDraw, cs.YDraw, cs.XDraw, cs.HDraw+cs.YDraw, draw.White)
	win.DrawLine(cs.XDraw, cs.HDraw+cs.YDraw, cs.WDraw+cs.XDraw, cs.HDraw+cs.YDraw, draw.White)
}

func (cs *CoordinateSurface2D) DrawXYDIVCardinateSurface2D(win draw.Window) {
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

func (cs *CoordinateSurface2D) View(win draw.Window) {
	cs.DrawCardinateSurface2D(win)
	cs.DrawXYDIVCardinateSurface2D(win)
	// cs.AddPointNill(10, 5, 6, draw.Red, win)
	// cs.AddPointFill(3, 2, 5, draw.Yellow, win)
	// cs.AddLineNill(0, 0, 3, 3, draw.Purple, win)
	// cs.AddLinePoint(4, 6, 4, 1, 10, draw.White, draw.Blue, win)
}

func (cs *CoordinateSurface2D) AddPointFill(X, Y, size int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := (cs.HDraw + cs.YDraw) - (Y * cs.DistanceY)
	win.FillEllipse(xn-int(size/2), yn-int(size/2), size, size, Color)

}

func (cs *CoordinateSurface2D) AddPointNill(X, Y, size int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := (cs.HDraw + cs.YDraw) - (Y * cs.DistanceY)
	win.DrawEllipse(xn-int(size/2), yn-int(size/2), size, size, Color)

}

func (cs *CoordinateSurface2D) AddLineNill(X, Y, x2, y2 int, Color draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := (cs.HDraw + cs.YDraw) - (Y * cs.DistanceY)
	xn2 := cs.XDraw + cs.DistanceX*x2
	yn2 := (cs.HDraw + cs.YDraw) - (y2 * cs.DistanceY)
	win.DrawLine(xn, yn, xn2, yn2, Color)
}

func (cs *CoordinateSurface2D) AddLinePoint(X, Y, x2, y2, size int, ColorLine, colorPoint draw.Color, win draw.Window) {
	xn := cs.XDraw + cs.DistanceX*X
	yn := (cs.HDraw + cs.YDraw) - (Y * cs.DistanceY)
	xn2 := cs.XDraw + cs.DistanceX*x2
	yn2 := (cs.HDraw + cs.YDraw) - (y2 * cs.DistanceY)
	win.FillEllipse(xn-int(size/2), yn-int(size/2), size, size, colorPoint)
	win.FillEllipse(xn2-int(size/2), yn2-int(size/2), size, size, colorPoint)
	win.DrawLine(xn, yn, xn2, yn2, ColorLine)

}
