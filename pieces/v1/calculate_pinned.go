package v1

func (p *Piece) calculatePinnedOptions(position int) {
	if p.PinnedToKing {
		for option := range p.Options {
			if option == p.PinnedByPosition {
				continue
			}
			if p.PinnedByPosition%8 == position%8 { // same column
				if p.PinnedByPosition/8 < position/8 {
					if option%8 == position%8 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition/8 > position/8 {
					if option%8 == position%8 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition/8 == position/8 { // same row
				if p.PinnedByPosition%8 < position%8 {
					if option/8 == position/8 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition%8 > position%8 {
					if option/8 == position/8 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition%9 == position%9 { // same left->right column
				if p.PinnedByPosition < position {
					if option%9 == position%9 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition > position {
					if option%9 == position%9 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			} else if p.PinnedByPosition%7 == position%7 { // same right->left column
				if p.PinnedByPosition < position {
					if option%7 == position%7 && option < position && option > p.PinnedByPosition {
						continue
					}
				}
				if p.PinnedByPosition > position {
					if option%7 == position%7 && option > position && option < p.PinnedByPosition {
						continue
					}
				}
			}
			delete(p.Options, option)
		}
	}
}
