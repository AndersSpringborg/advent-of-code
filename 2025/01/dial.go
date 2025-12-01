package main

type Dial struct {
	Index     uint
	ZeroIndex uint
}

type Direction int

const (
	Right Direction = iota
	Left  Direction = iota
)

const DIALSIZE = 100

func NewDial(index uint) *Dial {
	return &Dial{Index: index}
}

func (d *Dial) Rotate(direction Direction, steps uint) {
	if direction == Right {
		d.Index = (d.Index + steps) % DIALSIZE
	} else if direction == Left {
		d.Index = (d.Index + DIALSIZE - (steps % DIALSIZE)) % DIALSIZE
	}

	if d.Index == 0 {
		d.ZeroIndex++
	}
}
