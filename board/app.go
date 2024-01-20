package board

import (
	"chess/pieces"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	ScreenWidth  = 870
	ScreenHeight = 849
)

var (
	bI      *ebiten.Image
	optionI *ebiten.Image
	wpI     *ebiten.Image
	wbI     *ebiten.Image
	wkiI    *ebiten.Image
	wqI     *ebiten.Image
	wrI     *ebiten.Image
	wknI    *ebiten.Image
	bpI     *ebiten.Image
	bbI     *ebiten.Image
	bkiI    *ebiten.Image
	bqI     *ebiten.Image
	brI     *ebiten.Image
	bknI    *ebiten.Image
)

type App struct {
	touchIDs  []ebiten.TouchID
	op        ebiten.DrawImageOptions
	initiated bool

	whiteBoard map[int]*pieces.Piece
	blackBoard map[int]*pieces.Piece

	whitesTurn    bool
	selectedPiece *pieces.Piece
}

func (a *App) Update() error {
	if !a.initiated {
		a.init()
	}
	a.touchIDs = ebiten.AppendTouchIDs(a.touchIDs[:0])

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		X := int(math.Floor(float64(x) / 108.75))
		Y := int(math.Floor(float64(y) / 106.125))
		position := X + Y*8
		if position >= 64 {
			return nil
		}

		var board map[int]*pieces.Piece
		switch a.whitesTurn {
		case true:
			board = a.whiteBoard
		case false:
			board = a.blackBoard
		}
		win(board, a.whitesTurn)

		if piece, ok := board[position]; ok {
			piece.LastPosition = position
			a.selectedPiece = piece
		}

		if a.selectedPiece == nil {
			return nil
		}

		for option, take := range a.selectedPiece.EnPassantOptions {
			if position != option {
				continue
			}

			switch a.whitesTurn {
			case true:
				delete(a.blackBoard, take)
			case false:
				delete(a.whiteBoard, take)
			}

			board[option] = a.selectedPiece
			delete(board, a.selectedPiece.LastPosition)
			a.selectedPiece = nil

			a.whitesTurn = !a.whitesTurn
			a.calculateAllPositions(a.whiteBoard, a.blackBoard)
			return nil
		}

		for option := range a.selectedPiece.Options {
			if position != option {
				continue
			}

			a.TakeOrPromote(position)

			if a.selectedPiece.Kind == pieces.King && !a.selectedPiece.HasBeenMoved {
				castled := a.castle(option, board)
				if castled {
					return nil
				}
			}

			if a.selectedPiece.Kind == pieces.King || a.selectedPiece.Kind == pieces.Rook {
				a.selectedPiece.HasBeenMoved = true
			}

			board[option] = a.selectedPiece
			delete(board, a.selectedPiece.LastPosition)
			a.selectedPiece = nil

			a.whitesTurn = !a.whitesTurn
			a.calculateAllPositions(a.whiteBoard, a.blackBoard)
			return nil
		}
	}

	return nil
}

func win(board map[int]*pieces.Piece, colour bool) bool {
	for _, piece := range board {
		if len(piece.Options) != 0 {
			return false
		}
	}

	if colour {
		fmt.Println("Black wins")
	} else {
		fmt.Println("White wins")
	}
	return true
}

func (a *App) TakeOrPromote(position int) {
	if a.selectedPiece.Kind == pieces.Pawn {
		end := 7
		if a.selectedPiece.White == true {
			end = 0
		}
		if position/8 == end {
			a.selectedPiece.Kind = pieces.Queen // todo: add convert to better Piece logic
		}
	}

	switch a.whitesTurn {
	case true:
		delete(a.blackBoard, position)
	case false:
		delete(a.whiteBoard, position)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
