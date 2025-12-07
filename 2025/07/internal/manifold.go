package internal

type Manifold struct {
	beamSize  int
	beams     []bool
	splits    int
	timelines int
}

func NewManifold(beamCount int) *Manifold {
	beams := make([]bool, beamCount)
	return &Manifold{
		beamSize:  beamCount,
		beams:     beams,
		splits:    0,
		timelines: 0,
	}
}

func (m *Manifold) GetSplits() int {
	return m.splits
}

func (m *Manifold) GetTimelines() int {
	return m.timelines
}

func (m *Manifold) ParseLine(line string) {
	beamsBefore := countTrue(m.beams)

	for i := 0; i < len(m.beams); i++ {
		char := line[i]

		if char == '^' {
			if m.beams[i] {
				// Beam was already present, so it splits
				m.splits++
			}
			m.beams[i-1] = true
			m.beams[i+1] = true
			m.beams[i] = false

			beamsAfter := countTrue(m.beams)
			addedTimeline := beamsAfter - beamsBefore
			if addedTimeline > 0 {
				m.timelines += addedTimeline
			}
		} else if char == 'S' {
			m.timelines++
			m.beams[i] = true
		}
	}
}

func countTrue(beams []bool) int {
	count := 0
	for _, beam := range beams {
		if beam {
			count++
		}
	}
	return count
}
