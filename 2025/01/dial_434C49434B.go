package main

type Dial0x434C49434B struct {
	Index     int
	ZeroIndex int
}

func NewDial0x434C49434B(index int) *Dial0x434C49434B {
	return &Dial0x434C49434B{Index: index}
}

func (d *Dial0x434C49434B) Rotate(direction Direction, steps int) {
	for i := 0; i < steps; i++ {
		if direction == Right {
			d.Index++
			if d.Index == DIALSIZE {
				d.Index = 0
			}
		}
		if direction == Left {
			d.Index--
			if d.Index == -1 {
				d.Index = 99
			}
		}

		if d.Index == 0 {
			d.ZeroIndex++
		}

	}
}
