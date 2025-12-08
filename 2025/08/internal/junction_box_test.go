package internal

import "testing"

func TestGetBiggestCircuts(t *testing.T) {
	jbm := NewJunctionBoxMap()

	jbm.ParseJunctionBox("162,817,812")
	jbm.ParseJunctionBox("57,618,57")
	jbm.ParseJunctionBox("906,360,560")
	jbm.ParseJunctionBox("592,479,940")
	jbm.ParseJunctionBox("352,342,300")
	jbm.ParseJunctionBox("466,668,158")
	jbm.ParseJunctionBox("542,29,236")
	jbm.ParseJunctionBox("431,825,988")
	jbm.ParseJunctionBox("739,650,466")
	jbm.ParseJunctionBox("52,470,668")
	jbm.ParseJunctionBox("216,146,977")
	jbm.ParseJunctionBox("819,987,18")
	jbm.ParseJunctionBox("117,168,530")
	jbm.ParseJunctionBox("805,96,715")
	jbm.ParseJunctionBox("346,949,466")
	jbm.ParseJunctionBox("970,615,88")
	jbm.ParseJunctionBox("941,993,340")
	jbm.ParseJunctionBox("862,61,35")
	jbm.ParseJunctionBox("984,92,344")
	jbm.ParseJunctionBox("425,690,689")

	jbm.ConnectionAllJunctionBoxesToShortestDistance(10)
	biggestCircuts := jbm.GetBiggestCircuitsSum(3)

	if biggestCircuts != 40 {
		t.Errorf("Expected biggest circuts sum to be %d, got %d", 40, biggestCircuts)
	}

}
