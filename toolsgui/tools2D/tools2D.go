package tools2D

import (
	. "fmt"

	"github.com/gonutz/prototype/draw"
)

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
