package internal

import "testing"

func TestManifoldGetSplits(t *testing.T) {
	manifold := NewManifold(15)

	manifold.ParseLine(".......S.......")
	manifold.ParseLine("...............")
	manifold.ParseLine(".......^.......")
	manifold.ParseLine("...............")
	manifold.ParseLine("......^.^......")
	manifold.ParseLine("...............")
	manifold.ParseLine(".....^.^.^.....")
	manifold.ParseLine("...............")
	manifold.ParseLine("....^.^...^....")
	manifold.ParseLine("...............")
	manifold.ParseLine("...^.^...^.^...")
	manifold.ParseLine("...............")
	manifold.ParseLine("..^...^.....^..")
	manifold.ParseLine("...............")
	manifold.ParseLine(".^.^.^.^.^...^.")
	manifold.ParseLine("...............")

	expected := 21
	if manifold.GetSplits() != expected {
		t.Errorf("Expected %d splits, got %d", expected, manifold.GetSplits())
	}
}

func TestManifoldInitialization(t *testing.T) {
	beamCount := 5
	manifold := NewManifold(beamCount)

	if len(manifold.beams) != beamCount {
		t.Errorf("Expected %d beams, got %d", beamCount, len(manifold.beams))
	}

	if manifold.GetSplits() != 0 {
		t.Errorf("Expected 0 splits, got %d", manifold.GetSplits())
	}

	for i, beam := range manifold.beams {
		if beam {
			t.Errorf("Expected beam %d to be false, got true", i)
		}
	}
}
