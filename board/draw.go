package board

import (
	"chess/pieces"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func (a *App) drawOptions(screen *ebiten.Image, piece *pieces.Piece) {
	a.op.GeoM.Reset()

	for option := range piece.Options {
		a.op.GeoM.Reset()

		a.op.GeoM.Translate(108*float64(option%8), 108*math.Floor(float64(option/8)))
		screen.DrawImage(optionI, &a.op)
	}

	for option := range piece.EnPassantOptions {
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
		switch piece.Kind {
		case pieces.Pawn:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wpI, &a.op)
		case pieces.Knight:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wknI, &a.op)
		case pieces.Bishop:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wbI, &a.op)
		case pieces.Rook:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wrI, &a.op)
		case pieces.Queen:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wqI, &a.op)
		case pieces.King:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(wkiI, &a.op)
		}
	}

	for position, piece := range a.blackBoard {
		a.op.GeoM.Reset()
		switch piece.Kind {
		case pieces.Pawn:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bpI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case pieces.Knight:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bknI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case pieces.Bishop:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bbI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case pieces.Rook:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(brI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case pieces.Queen:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bqI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		case pieces.King:
			a.op.GeoM.Translate(108*float64(position%8), 108*math.Floor(float64(position/8)))
			screen.DrawImage(bkiI, &a.op)
			if position != piece.LastPosition {
				screen.DrawImage(newPositionI, &a.op)
				a.op.GeoM.Reset()
				a.op.GeoM.Translate(108*float64(piece.LastPosition%8), 108*math.Floor(float64(piece.LastPosition/8)))
				screen.DrawImage(lastPositionI, &a.op)
			}
		}
	}
}
