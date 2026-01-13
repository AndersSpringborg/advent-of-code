package internal

import (
	"fmt"
	"testing"
)

func TestServerGraph_ParseLine(t *testing.T) {
	sg := NewServerGraph()
	sg.ParseLine("aaa: you hhh")
	sg.ParseLine("you: bbb ccc")
	sg.ParseLine("bbb: ddd eee")
	sg.ParseLine("ccc: ddd eee fff")
	sg.ParseLine("ddd: ggg")
	sg.ParseLine("eee: out")
	sg.ParseLine("fff: out")
	sg.ParseLine("ggg: out")
	sg.ParseLine("hhh: ccc fff iii")
	sg.ParseLine("iii: out")

	fmt.Println(sg.ToString())
	totalPaths := sg.TotalPathsFromYouToOut()
	expectedPaths := 5

	if totalPaths != expectedPaths {
		t.Errorf("Expected total paths %d, got %d", expectedPaths, totalPaths)
	}
}
func TestServerGraph_ParseLineDac(t *testing.T) {
	sg := NewServerGraph()
	sg.ParseLine("svr: aaa bbb")
	sg.ParseLine("aaa: fft")
	sg.ParseLine("fft: ccc")
	sg.ParseLine("bbb: tty")
	sg.ParseLine("tty: ccc")
	sg.ParseLine("ccc: ddd eee")
	sg.ParseLine("ddd: hub")
	sg.ParseLine("hub: fff")
	sg.ParseLine("eee: dac")
	sg.ParseLine("dac: fff")
	sg.ParseLine("fff: ggg hhh")
	sg.ParseLine("ggg: out")
	sg.ParseLine("hhh: out")

	fmt.Println(sg.ToString())
	totalPaths := sg.TotalPathsFromYouToOutWithFCCandDAC()
	expectedPaths := 2

	if totalPaths != expectedPaths {
		t.Errorf("Expected total paths %d, got %d", expectedPaths, totalPaths)
	}
}
