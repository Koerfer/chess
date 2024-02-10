package v2

import (
	engine "chess/engine/v2"
	v2 "chess/pieces/v2"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
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

	whiteBoard map[int]v2.PieceInterface
	blackBoard map[int]v2.PieceInterface

	whitesTurn    bool
	selectedPiece v2.PieceInterface

	engine engine.Engine
}

func (a *App) Update() error {
	if !a.initiated {
		a.init()
	}
	a.touchIDs = ebiten.AppendTouchIDs(a.touchIDs[:0])

	var board map[int]v2.PieceInterface
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
		x, y := ebiten.CursorPosition()
		X := int(math.Floor(float64(x) / (ScreenWidth / 8)))
		Y := int(math.Floor(float64(y) / (ScreenHeight / 8)))
		position := X + Y*8
		if position >= 64 {
			return nil
		}

		if piece, ok := board[position]; ok {
			piece.SetLastPosition(position)
			a.selectedPiece = piece
		}

		if a.selectedPiece == nil {
			return nil
		}

		switch v2.CheckPieceKindFromAny(a.selectedPiece) {
		case v2.PieceKindPawn:
			p := a.selectedPiece.(*v2.Pawn)
			if stop := a.enPassant(p, position, board); stop {
				a.recalculateBoard()
				a.engine.Init(a.whiteBoard, a.blackBoard)
				return nil
			}
			if stop := a.normalPawn(p, position, board); stop {
				a.recalculateBoard()
				a.engine.Init(a.whiteBoard, a.blackBoard)
				return nil
			}
		case v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			if stop := a.normal(a.selectedPiece, position, board); stop {
				a.recalculateBoard()
				a.engine.Init(a.whiteBoard, a.blackBoard)
				return nil
			}
		case v2.PieceKindKing:
			p := a.selectedPiece.(*v2.King)
			if stop := a.normalKing(p, position, board); stop {
				a.recalculateBoard()
				a.engine.Init(a.whiteBoard, a.blackBoard)
				return nil
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind getting piece where clicked")
		}
	}

	return nil
}

func (a *App) recalculateBoard() {
	a.selectedPiece = nil

	a.whitesTurn = !a.whitesTurn
	for _, piece := range a.whiteBoard {
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn, v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtectedBy(make(map[int]v2.PieceInterface))
			piece.SetAttackedBy(make(map[int]v2.PieceInterface))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindKing:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindInvalid:
			panic("invalid piece kind when calculating options")
		}
	}
	for _, piece := range a.blackBoard {
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn, v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtectedBy(make(map[int]v2.PieceInterface))
			piece.SetAttackedBy(make(map[int]v2.PieceInterface))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindKing:
			piece.SetOptions(make(map[int]struct{}))
			piece.SetProtecting(make(map[int]v2.PieceInterface))
		case v2.PieceKindInvalid:
			panic("invalid piece kind when calculating options")
		}
	}
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
}

func (a *App) enPassant(pawn *v2.Pawn, position int, board map[int]v2.PieceInterface) bool {
	for option, take := range pawn.EnPassantOptions {
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
		delete(board, pawn.LastPosition)
		return true
	}

	return false
}

func (a *App) normalPawn(pawn *v2.Pawn, position int, board map[int]v2.PieceInterface) bool {
	for option := range pawn.Options {
		if position != option {
			continue
		}

		end := 7
		if pawn.White == true {
			end = 0
		}
		if position/8 == end {
			board[position] = &v2.Queen{
				White:        pawn.White,
				LastPosition: pawn.LastPosition,
				Options:      make(map[int]struct{}),

				PinnedToKing:     false,
				PinnedByPosition: 0,
				PinnedByPiece:    nil,
				Protecting:       make(map[int]v2.PieceInterface),
				Value:            9,
				AttackedBy:       make(map[int]v2.PieceInterface),
				ProtectedBy:      make(map[int]v2.PieceInterface),
			}
		} else {
			board[option] = pawn
		}

		switch a.whitesTurn {
		case true:
			delete(a.blackBoard, position)
		case false:
			delete(a.whiteBoard, position)
		}

		delete(board, pawn.LastPosition)
		return true
	}

	return false
}

func (a *App) normal(piece v2.PieceInterface, position int, board map[int]v2.PieceInterface) bool {
	for option := range piece.GetOptions() {
		if position != option {
			continue
		}

		switch a.whitesTurn {
		case true:
			delete(a.blackBoard, position)
		case false:
			delete(a.whiteBoard, position)
		}

		board[option] = a.selectedPiece
		delete(board, piece.GetLastPosition())
		return true
	}

	return false
}

func (a *App) normalKing(king *v2.King, position int, board map[int]v2.PieceInterface) bool {
	for option := range king.Options {
		if position != option {
			continue
		}

		switch a.whitesTurn {
		case true:
			delete(a.blackBoard, position)
		case false:
			delete(a.whiteBoard, position)
		}

		if !king.HasBeenMoved {
			castled := a.castle(king, option, board)
			if castled {
				return true
			}
		}

		king.HasBeenMoved = true

		board[option] = a.selectedPiece
		delete(board, king.LastPosition)
		return true
	}

	return false
}

func win(board map[int]v2.PieceInterface, colour bool) bool {
	var checked bool
	for _, piece := range board {
		if len(piece.GetOptions()) != 0 {
			return false
		}
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn, v2.PieceKindKnight, v2.PieceKindBishop, v2.PieceKindRook, v2.PieceKindQueen:
			if len(piece.GetOptions()) != 0 {
				return false
			}
		case v2.PieceKindKing:
			p := piece.(*v2.King)
			checked = p.Checked
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind when checking if win")
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
