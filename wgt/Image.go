package wgt

import "github.com/gonutz/prototype/draw"

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
	// Println(im.Image[i])
	win.DrawImageFileTo(im.Image[i], im.X, im.Y, im.W, im.H, im.DEG)
}
