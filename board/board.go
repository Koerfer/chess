package board

import (
	"chess/engine/v1"
)

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	a.initWhiteBoard()
	a.initBlackBoard()
	a.initImages()
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
	a.engine = v1.Engine{}
	a.engine.Init(3)
}
