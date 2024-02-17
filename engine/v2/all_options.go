package v2

func (e *Engine) getAllOptions() {
	for _, piece := range e.whiteBoard {
		for option := range piece.GetOptions() {
			e.allWhiteOptions[option] = struct{}{}
		}
	}

	for _, piece := range e.blackBoard {
		for option := range piece.GetOptions() {
			e.allBlackOptions[option] = struct{}{}
		}
	}

	e.evalOptions()
}

func (e *Engine) evalOptions() {
	e.whiteEval += float64(len(e.allWhiteOptions)) * 0.05
	e.blackEval += float64(len(e.allBlackOptions)) * 0.05
}
