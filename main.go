package main

import (
	"chess/board/v1"
	v2 "chess/board/v2"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

func main() {
	ebiten.SetWindowSize(v1.ScreenWidth, v1.ScreenHeight)
	ebiten.SetWindowTitle("Chess")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	version := 2

	switch version {
	case 1:
		if err := ebiten.RunGame(&v1.App{}); err != nil {
			panic(err)
		}
	case 2:
		if err := ebiten.RunGame(&v2.App{}); err != nil {
			panic(err)
		}
	}
}
