package v2

import (
	v2 "chess/pieces/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func (a *App) drawOptions(screen *ebiten.Image, piece v2.PieceInterface) {
	a.op.GeoM.Reset()

	options := make(map[int]struct{})
	enPassantOptions := make(map[int]int)

	switch v2.CheckPieceKindFromAny(piece) {
	case v2.PieceKindPawn:
		p := piece.(*v2.Pawn)
		options = p.Options
		enPassantOptions = p.EnPassantOptions
	case v2.PieceKindKnight:
		p := piece.(*v2.Knight)
		options = p.Options
	case v2.PieceKindBishop:
		p := piece.(*v2.Bishop)
		options = p.Options
	case v2.PieceKindRook:
		p := piece.(*v2.Rook)
		options = p.Options
	case v2.PieceKindQueen:
		p := piece.(*v2.Queen)
		options = p.Options
	case v2.PieceKindKing:
		p := piece.(*v2.King)
		options = p.Options
	case v2.PieceKindInvalid:
		panic("invalid piece kind when drawing options")
	}

	for option := range options {
		a.op.GeoM.Reset()

		a.op.GeoM.Translate(108*float64(option%8), 108*math.Floor(float64(option/8)))
		screen.DrawImage(optionI, &a.op)
	}

	for option := range enPassantOptions {
		a.op.GeoM.Reset()

		a.op.GeoM.Translate(108*float64(option%8), 108*math.Floor(float64(option/8)))
		screen.DrawImage(optionI, &a.op)
	}

}

func (a *App) Draw(screen *ebiten.Image) {
	a.op.GeoM.Reset()

	screen.DrawImage(bI, &a.op)
	a.drawPieces(screen)
	if a.selectedPiece != nil {
		a.drawOptions(screen, a.selectedPiece)
	}
}

func (a *App) drawPieces(screen *ebiten.Image) {
	a.op.GeoM.Reset()

	for position, piece := range a.whiteBoard {
		a.op.GeoM.Reset()
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn:
			p := piece.(*v2.Pawn)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wpI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindKnight:
			p := piece.(*v2.Knight)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wknI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindBishop:
			p := piece.(*v2.Bishop)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wbI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindRook:
			p := piece.(*v2.Rook)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wrI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindQueen:
			p := piece.(*v2.Queen)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wqI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindKing:
			p := piece.(*v2.King)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wkiI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind when drawing pieces")
		}
	}

	for position, piece := range a.blackBoard {
		a.op.GeoM.Reset()
		switch v2.CheckPieceKindFromAny(piece) {
		case v2.PieceKindPawn:
			p := piece.(*v2.Pawn)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bpI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindKnight:
			p := piece.(*v2.Knight)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bknI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindBishop:
			p := piece.(*v2.Bishop)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bbI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindRook:
			p := piece.(*v2.Rook)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(brI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindQueen:
			p := piece.(*v2.Queen)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bqI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindKing:
			p := piece.(*v2.King)
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bkiI, &a.op)
			if position != p.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(p.LastPosition%8), 108*math.Floor(float64(p.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case v2.PieceKindInvalid:
			panic("invalid piece kind when drawing pieces")
		}
	}
}
