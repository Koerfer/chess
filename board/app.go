package board

import (
	"chess/engine/v1"
	"chess/pieces"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
	"time"
)

const (
	ScreenWidth  = 864
	ScreenHeight = 864
)

var (
	bI            *ebiten.Image
	optionI       *ebiten.Image
	lastPositionI *ebiten.Image
	newPositionI  *ebiten.Image
	wpI           *ebiten.Image
	wbI           *ebiten.Image
	wkiI          *ebiten.Image
	wqI           *ebiten.Image
	wrI           *ebiten.Image
	wknI          *ebiten.Image
	bpI           *ebiten.Image
	bbI           *ebiten.Image
	bkiI          *ebiten.Image
	bqI           *ebiten.Image
	brI           *ebiten.Image
	bknI          *ebiten.Image
)

type App struct {
	touchIDs  []ebiten.TouchID
	op        ebiten.DrawImageOptions
	initiated bool

	whiteBoard map[int]*pieces.Piece
	blackBoard map[int]*pieces.Piece

	whitesTurn    bool
	selectedPiece *pieces.Piece

	engine v1.Engine
}

func (a *App) Update() error {
	if !a.initiated {
		a.init()
	}
	a.touchIDs = ebiten.AppendTouchIDs(a.touchIDs[:0])

	var board map[int]*pieces.Piece
	switch a.whitesTurn {
	case true:
		board = a.whiteBoard
	case false:
		board = a.blackBoard
	}

	if win(board, a.whitesTurn) {
		return nil
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !a.whitesTurn {
			start := time.Now()
			fmt.Println("Starting")
			option := a.engine.Start(a.whiteBoard, a.blackBoard, 0, a.whitesTurn)
			a.selectedPiece = option.Piece

			if option.EnPassant != 0 {
				if stop := a.enPassant(option.EnPassant, board); stop {
					duration := time.Since(start)
					fmt.Println(duration)
					return nil
				}
			}

			if stop := a.normal(option.MoveTo, board); stop {
				duration := time.Since(start)
				fmt.Println(duration)
				return nil
			}

			duration := time.Since(start)
			fmt.Println(duration)
			return nil
		}

		x, y := ebiten.CursorPosition()
		X := int(math.Floor(float64(x) / (ScreenWidth / 8)))
		Y := int(math.Floor(float64(y) / (ScreenHeight / 8)))
		position := X + Y*8
		if position >= 64 {
			return nil
		}

		if piece, ok := board[position]; ok {
			piece.LastPosition = position
			a.selectedPiece = piece
		}

		if a.selectedPiece == nil {
			return nil
		}

		if stop := a.enPassant(position, board); stop {
			return nil
		}

		if stop := a.normal(position, board); stop {
			return nil
		}
	}

	return nil
}

func (a *App) enPassant(position int, board map[int]*pieces.Piece) bool {
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
		return true
	}

	return false
}

func (a *App) normal(position int, board map[int]*pieces.Piece) bool {
	for option := range a.selectedPiece.Options {
		if position != option {
			continue
		}

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

		if a.selectedPiece.Kind == pieces.King && !a.selectedPiece.HasBeenMoved {
			castled := a.castle(option, board)
			if castled {
				return true
			}
		}

		a.selectedPiece.HasBeenMoved = true

		board[option] = a.selectedPiece
		delete(board, a.selectedPiece.LastPosition)
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func win(board map[int]*pieces.Piece, colour bool) bool {
	var checked bool
	for _, piece := range board {
		if len(piece.Options) != 0 {
			return false
		}
		if piece.Kind == pieces.King {
			checked = piece.Checked
		}
	}

	if !checked {
		fmt.Println("Draw due to stalemate")
		return true
	}

	if colour {
		fmt.Println("Black wins")
	} else {
		fmt.Println("White wins")
	}
	return true
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
