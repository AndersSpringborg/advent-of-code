package internal

import (
	"fmt"
	"log"
	"math"
	"sort"
)

type JunctionBox struct {
	id int
	x  int
	y  int
	z  int
}

func (jb *JunctionBox) String() string {
	return fmt.Sprintf("%d", jb.id)
}

type Circut struct {
	junctionBoxes []JunctionBox
}

func (c *Circut) Equals(other *Circut) bool {
	if c.Length() != other.Length() {
		return false
	}
	for _, jb := range c.junctionBoxes {
		if !other.junctionBoxExists(&jb) {
			return false
		}
	}
	return true
}

func (c *Circut) GetIds() []int {
	ids := make([]int, 0)
	for _, jb := range c.junctionBoxes {
		ids = append(ids, jb.id)
	}
	return ids
}

func (c *Circut) Length() int {
	return len(c.junctionBoxes)
}

func (c *Circut) junctionBoxExists(jb *JunctionBox) bool {
	key := jb.String()
	for _, existingJB := range c.junctionBoxes {
		if existingJB.String() == key {
			return true
		}
	}
	return false
}

func (c *Circut) AddBox(a JunctionBox) {
	c.junctionBoxes = append(c.junctionBoxes, a)
}

type JunctionBoxMap struct {
	currentID   int
	Boxes       map[int]*JunctionBox
	connections []*Circut
}

func NewJunctionBoxMap() *JunctionBoxMap {
	return &JunctionBoxMap{
		Boxes: make(map[int]*JunctionBox),
	}
}

func (jbm *JunctionBoxMap) ParseJunctionBox(line string) {
	var x, y, z int
	_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		log.Fatal("Error parsing junction box:", err)
	}

	jbm.Boxes[jbm.currentID] = &JunctionBox{
		id: jbm.currentID,
		x:  x,
		y:  y,
		z:  z,
	}
	jbm.currentID++
}

type Pair struct {
	boxa     *JunctionBox
	boxb     *JunctionBox
	distance float64
}

func (jbm *JunctionBoxMap) FindClosetsJunctionsBoxes() []*Pair {

	connections := make([]*Pair, 0)

	// only calculate each pair once
	for i := 0; i < len(jbm.Boxes); i++ {
		boxA := jbm.Boxes[i]
		for j := i + 1; j < len(jbm.Boxes); j++ {
			boxB := jbm.Boxes[j]
			distance := EuclidianDistance(boxA, boxB)
			pair := &Pair{
				boxa:     boxA,
				boxb:     boxB,
				distance: distance,
			}
			connections = append(connections, pair)
		}
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	return connections
}

func (jbm *JunctionBoxMap) ConnectionAllJunctionBoxesToShortestDistance(n int) {
	for _, box := range jbm.Boxes {
		jbm.connections = append(jbm.connections, &Circut{
			junctionBoxes: []JunctionBox{*box},
		})
	}

	shortestDistance := jbm.FindClosetsJunctionsBoxes()

	for i := 0; i < n; i++ {
		if i%100 == 0 {
			log.Printf("Processing connection %d/%d. Current curcuit count: %d", i, n, len(jbm.connections))
		}

		pair := shortestDistance[i]
		boxA := pair.boxa
		boxB := pair.boxb

		circuitA := jbm.FindCircutContainingBox(boxA)
		circuitB := jbm.FindCircutContainingBox(boxB)
		if !circuitA.Equals(circuitB) {
			// merge circuits
			for _, jb := range circuitB.junctionBoxes {
				circuitA.AddBox(jb)
			}
			// remove circuitB from jbm.connections
			newConnections := make([]*Circut, 0)
			for _, c := range jbm.connections {
				if !c.Equals(circuitB) {
					newConnections = append(newConnections, c)
				}
			}
			jbm.connections = newConnections
		}
	}

}

func (jbm *JunctionBoxMap) GetBiggestCircuitsSum(n int) int {

	// print all circuits with arrow
	for i, circuit := range jbm.connections {
		log.Printf("Circuit %d: - %v", i, circuit.GetIds())
	}

	sort.Slice(jbm.connections, func(a, b int) bool {
		return jbm.connections[a].Length() > jbm.connections[b].Length()
	})

	if len(jbm.connections) < n {
		log.Default().Printf("Amount of circuits: %d is less than requested biggest circuits: %d", len(jbm.connections), n)
		return 0
	}

	first := jbm.connections[0]
	sum := first.Length()

	fmt.Printf("circuit lengths: %d,  %v", first.Length(), first.GetIds())
	for _, circuit := range jbm.connections[1:n] {
		fmt.Printf(", %d,  %v", circuit.Length(), circuit.GetIds())
		sum *= circuit.Length()
	}

	return sum

}

// part 2
func (jbm *JunctionBoxMap) GetLastConnectingPairXMultiplyX() int {
	shortestDistance := jbm.FindClosetsJunctionsBoxes()
	allIds := make(map[int]struct{})
	for _, box := range jbm.Boxes {
		allIds[box.id] = struct{}{}
	}

	i := 0
	for len(allIds) > 0 {
		if i%100 == 0 {
			log.Printf("Processing connection %d. Remaining boxes to connect: %d", i, len(allIds))
		}
		pair := shortestDistance[i]
		boxA := pair.boxa
		boxB := pair.boxb

		if _, existsA := allIds[boxA.id]; existsA {
			delete(allIds, boxA.id)
		}
		if _, existsB := allIds[boxB.id]; existsB {
			delete(allIds, boxB.id)
		}
		i++
	}

	lastPair := shortestDistance[i-1]
	println("Last connecting pair:", lastPair.boxa.id, "and", lastPair.boxb.id)

	return lastPair.boxa.x * lastPair.boxb.x

}
func EuclidianDistance(a *JunctionBox, b *JunctionBox) float64 {
	return math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2)
}

func (jbm *JunctionBoxMap) FindCircutContainingBox(a *JunctionBox) *Circut {
	for _, circut := range jbm.connections {
		if circut.junctionBoxExists(a) {
			return circut
		}
	}

	return nil
}
