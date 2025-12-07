package internal

import "testing"

func TestManifold2GetTimelinesDP(t *testing.T) {
	manifold := NewManifold2()

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

	expected := 40
	result := manifold.GetTimelinesDP()
	if result != expected {
		t.Errorf("Expected %d timelines, got %d", expected, result)
	}
}

func TestManifold2Initialization(t *testing.T) {
	manifold := NewManifold2()

	if manifold.lineLength != 0 {
		t.Errorf("Expected lineLength to be 0, got %d", manifold.lineLength)
	}

	if len(manifold.linesProcessed) != 0 {
		t.Errorf("Expected empty linesProcessed, got length %d", len(manifold.linesProcessed))
	}
}
