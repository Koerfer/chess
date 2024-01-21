package board

import "chess/engine"

func (a *App) init() {
	defer func() {
		a.initiated = true
	}()

	a.initWhiteBoard()
	a.initBlackBoard()
	a.initImages()
	a.calculateAllPositions(a.whiteBoard, a.blackBoard)
	a.engine = engine.Engine{}
	a.engine.Init()
}
