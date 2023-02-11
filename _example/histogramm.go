package main

import (
	"github.com/gonutz/prototype/draw"
<<<<<<<< HEAD:toolsgui/_example/histogramm.go
	tools2D "proto_ui/toolsgui/wgt"
========
	"github.com/yuranignatenko/proto_ui/wgt"
>>>>>>>> main:_example/histogramm.go
)

var dt = []int{0, 10, 9, 2, 8, 7, 11, 15, 17}
var data = []float64{0.3, 0.5, 0.9, 1., 4., 9., 9.9}
var data2 = []int{1, 3, 7, 7, 9, 8, 5, 20, 4, 70, 76, 33, 35, 9, 90}

var hist1 = wgt.NewHistogramma(dt, 50, 50, 300, 100, 100, draw.White)
var hist2 = wgt.NewHistogramma(data2, 50, 200, 300, 100, 100, draw.Green)

var hist3 = wgt.NewHistogrammaFloat(data, 400, 50, 300, 100, 100, draw.Blue)
var hist4 = wgt.NewHistogrammaFloat(data, 400, 200, 300, 100, 100, draw.Yellow)

var hist5 = wgt.NewHistogramma(data2, 50, 350, 300, 100, 100, draw.Red)

var c = 0

func update(win draw.Window) {
	if c < 80 {
		c++
		oldData := hist1.Data
		newData := make([]int, 0)
		for _, val := range oldData {
			newData = append(newData, val+1)
			hist1.Data = newData
		}
	} else {
		hist1.Data = dt
		c = 0
	}
	hist1.View(win, 1, 1, 0, 15)
	hist2.View(win, 2, 1, 1, 15)
	hist3.View(win, 3, 1, 5, 15)
	hist4.View(win, 4, 1, 15, 15)
	hist5.View(win, 5, 1, 35, 25)
}

func main() {
	draw.RunWindow("toolsgui-wgt-example-histogramm", 1000, 600, update)
}
