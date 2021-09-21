package tools2D

import (
	"strings"

	"github.com/gonutz/prototype/draw"
)

type EntryRect struct {
	X, Y, W, H                                                   int
	FontSize                                                     float32
	FgActive, FgDisable, FgFontActive, FgFontDisable, ColorFlash draw.Color
	ArrayNewText                                                 []string
	FlagisClick                                                  bool
	tempFlag                                                     bool
	MAXLENSYNB                                                   int
	// Command                                          func()
}

func NewEntryRect(text string, maxlensymb int, FgActive, FgDisable, FgFontActive, FgFontDisable, colorFlash draw.Color, X, Y, W, H int) EntryRect {
	var ent EntryRect
	res := ent.SetText(text)
	return EntryRect{
		FgActive:      FgActive,
		FgDisable:     FgDisable,
		FgFontActive:  FgFontActive,
		FgFontDisable: FgFontDisable,
		ColorFlash:    colorFlash,
		FlagisClick:   false,
		tempFlag:      false,
		ArrayNewText:  res,
		X:             X,
		Y:             Y,
		W:             W,
		H:             H,
		FontSize:      float32(H/19) + 0.5,
		MAXLENSYNB:    maxlensymb,
		// Command:       command,
	}
}

func (b *EntryRect) SetText(text string) []string {
	newLne := strings.Split(text, "")
	b.ArrayNewText = newLne
	return newLne
}

func (b *EntryRect) GetText() string {
	newLne := strings.Join(b.ArrayNewText[:], "")
	return newLne
}

func (b *EntryRect) View(window draw.Window) {
	// centerX, centerY := b.W/2, b.H/2

	newLne := strings.Join(b.ArrayNewText[:], "")

	mouseX, mouseY := window.MousePosition()
	mouseInRect := (b.W+b.X > mouseX && mouseX > b.X) && (b.H+b.Y > mouseY && mouseY > b.Y)

	colorBg := b.FgDisable
	colorFont := b.FgFontDisable

	// Println(b.ArrayNewText)

	if mouseInRect {
		colorBg = b.ColorFlash
		colorFont = b.FgDisable
	}

	if mouseInRect && b.FlagisClick {
		colorBg = b.ColorFlash
		colorFont = b.FgFontActive
	}
	if mouseInRect == false && b.FlagisClick {
		colorBg = b.FgActive
		colorFont = b.FgFontActive
	}

	window.FillRect(b.X, b.Y, b.W, b.H, colorBg)
	window.DrawScaledText(newLne, b.X, b.Y, b.FontSize, colorFont)

	// check all mouse clicks that happened during this frame

	arKeys := []draw.Key{
		draw.KeyA, draw.KeyB, draw.KeyC,
		draw.KeyD, draw.KeyE, draw.KeyF,
		draw.KeyG, draw.KeyH, draw.KeyI,
		draw.KeyJ, draw.KeyK, draw.KeyL,
		draw.KeyM, draw.KeyN, draw.KeyO,
		draw.KeyP, draw.KeyQ, draw.KeyR,
		draw.KeyS, draw.KeyT, draw.KeyU,
		draw.KeyV, draw.KeyW, draw.KeyX,
		draw.KeyY, draw.KeyZ, draw.Key1,
		draw.Key2, draw.Key3, draw.Key4,

		draw.Key5, draw.Key6, draw.Key7,
		draw.Key8, draw.Key9, draw.Key0,
		draw.KeyBackspace, draw.KeySpace,
		draw.KeyRightAlt, draw.KeyLeftAlt,
	}

	arNamesKeys := []string{
		"A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z", "1", "2",
		"3", "4", "5", "6", "7", "8", "9",
		"0", "BACKSPACE", " ", ".", ".",
	}

	// Println(len(arKeys), len(arNamesKeys))

	if b.tempFlag == false {
		b.tempFlag = true
	}

	for _, click := range window.Clicks() {
		clickInRect := (b.W+b.X > click.X && click.X > b.X) && (b.H+b.Y > click.Y && click.Y > b.Y)
		if clickInRect && b.tempFlag {
			b.FlagisClick = true
			b.tempFlag = true
			colorBg = b.FgActive
			colorFont = b.FgFontActive
		}
		if clickInRect != true {
			b.FlagisClick = false
			b.tempFlag = false
		}
	}

	if b.FlagisClick {
		for ind, key := range arKeys {
			// Println(key.String)
			if window.WasKeyPressed(key) && arNamesKeys[ind] != "BACKSPACE" {
				if len(b.ArrayNewText) < b.MAXLENSYNB {
					b.ArrayNewText = append(b.ArrayNewText, arNamesKeys[ind])
				}
			}
			if window.WasKeyPressed(key) && arNamesKeys[ind] == "BACKSPACE" && len(b.ArrayNewText) != 0 {
				newLne := b.ArrayNewText[0 : len(b.ArrayNewText)-1]
				b.ArrayNewText = make([]string, 0)
				for _, val := range newLne {
					b.ArrayNewText = append(b.ArrayNewText, val)
				}
			}
		}
	} // b.Command()

}
