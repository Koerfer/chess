package v2

import (
	v2 "chess/pieces/v2"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
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

	whiteBoard map[int]any
	blackBoard map[int]any

	whitesTurn    bool
	selectedPiece any
}

func (a *App) Update() error {
	if !a.initiated {
		a.init()
	}
	a.touchIDs = ebiten.AppendTouchIDs(a.touchIDs[:0])

	var board map[int]any
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
			switch v2.CheckPieceKindFromAny(piece) {
			case v2.PieceKindPawn:
				p := piece.(*v2.Pawn)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindKnight:
				p := piece.(*v2.Knight)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindBishop:
				p := piece.(*v2.Bishop)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindRook:
				p := piece.(*v2.Rook)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindQueen:
				p := piece.(*v2.Queen)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindKing:
				p := piece.(*v2.King)
				p.LastPosition = position
				a.selectedPiece = p
			case v2.PieceKindInvalid:
				log.Fatal("invalid piece kind getting piece where clicked")
			}
		}

		if a.selectedPiece == nil {
			return nil
		}

		switch v2.CheckPieceKindFromAny(a.selectedPiece) {
		case v2.PieceKindPawn:
			p := a.selectedPiece.(*v2.Pawn)
			if stop := a.enPassant(p, position, board); stop {
				return nil
			}
			if stop := a.normalPawn(p, position, board); stop {
				return nil
			}
		case v2.PieceKindKnight:
			p := a.selectedPiece.(*v2.Knight)
			if stop := a.normalKnight(p, position, board); stop {
				return nil
			}
		case v2.PieceKindBishop:
			p := a.selectedPiece.(*v2.Bishop)
			if stop := a.normalBishop(p, position, board); stop {
				return nil
			}
		case v2.PieceKindRook:
			p := a.selectedPiece.(*v2.Rook)
			if stop := a.normalRook(p, position, board); stop {
				return nil
			}
		case v2.PieceKindQueen:
			p := a.selectedPiece.(*v2.Queen)
			if stop := a.normalQueen(p, position, board); stop {
				return nil
			}
		case v2.PieceKindKing:
			p := a.selectedPiece.(*v2.King)
			if stop := a.normalKing(p, position, board); stop {
				return nil
			}
		case v2.PieceKindInvalid:
			log.Fatal("invalid piece kind getting piece where clicked")
		}
	}

	return nil
}

func (a *App) enPassant(pawn *v2.Pawn, position int, board map[int]any) bool {
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
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalPawn(pawn *v2.Pawn, position int, board map[int]any) bool {
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
				Piece: &v2.Piece{
					White:        pawn.White,
					LastPosition: pawn.LastPosition,
					Options:      make(map[int]struct{}),
				},
				PinnedToKing:     false,
				PinnedByPosition: 0,
				PinnedByPiece:    nil,
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
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalKnight(knight *v2.Knight, position int, board map[int]any) bool {
	for option := range knight.Options {
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
		delete(board, knight.LastPosition)
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalRook(rook *v2.Rook, position int, board map[int]any) bool {
	for option := range rook.Options {
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
		delete(board, rook.LastPosition)
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalBishop(bishop *v2.Bishop, position int, board map[int]any) bool {
	for option := range bishop.Options {
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
		delete(board, bishop.LastPosition)
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalQueen(queen *v2.Queen, position int, board map[int]any) bool {
	for option := range queen.Options {
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
		delete(board, queen.LastPosition)
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func (a *App) normalKing(king *v2.King, position int, board map[int]any) bool {
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
		a.selectedPiece = nil

		a.whitesTurn = !a.whitesTurn
		a.calculateAllPositions(a.whiteBoard, a.blackBoard)
		return true
	}

	return false
}

func win(board map[int]any, colour bool) bool {
	var checked bool
	for _, piece := range board {
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn:
			p := piece.(*v2.Pawn)
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindKnight:
			p := piece.(*v2.Knight)
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindBishop:
			p := piece.(*v2.Bishop)
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindRook:
			p := piece.(*v2.Rook)
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindQueen:
			p := piece.(*v2.Queen)
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindKing:
			p := piece.(*v2.King)
			checked = p.Checked
			if len(p.Options) != 0 {
				return false
			}
		case v2.PieceKindInvalid:
			log.Fatal("invalid piece kind when checking if win")
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
