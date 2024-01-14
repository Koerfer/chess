package main

import (
	"chess/board"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

func main() {
	ebiten.SetWindowSize(board.ScreenWidth, board.ScreenHeight)
	ebiten.SetWindowTitle("Chess")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	if err := ebiten.RunGame(&board.App{}); err != nil {
		log.Fatal(err)
	}
}
