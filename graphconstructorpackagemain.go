package main

import (
	g "goGraphs/graphConstructor"
	//  "github.com/damilarelana/introToGo/mathpackage"
)

func main() {

	//initialize a slice of Nodes via what ReadData returns
	testNodesSlice := g.ReadData() // this would inherently call the scanf function

	// iterate and print each node
	for _, testNode := range testNodesSlice {
		g.PrintData(&testNode)
	}
}
