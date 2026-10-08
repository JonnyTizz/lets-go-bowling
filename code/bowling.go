package bowling

type Bowling struct {
}

func NewBowling() *Bowling {
	return &Bowling{}
}

func (b Bowling) Roll(pins int) {
}

func (b *Bowling) Score() {
}
