package graphConstructor

import "fmt"

// Node i.e. graph vertices is defined here "from perspective of a node and it's child nodes"
// we do NOT consider the upwards context of it's parent node
type Node struct {
	Value int
	Left  *Node // i.e. the left child node is of a type Node
	Right *Node // i.e. the right child node is also of a type Node
}

// ReadData function reads from a structured *.in file using Scanf
// assigns nodes values at their appropriate index, and returns a slice of nodes
func ReadData() []Node {
	var N int           // is the first row in *.in file. It defines number of nodes
	fmt.Scanf("%d", &N) // reads and assigns to inputted value to the exact memory location of N, i.e. no copies

	// defines/initializes slice of nodes to be constructured and returned. Also helps sort out nil child nodes
	nodes := make([]Node, N)

	// This iterates through the structure *.in file [i.e. other rows after row containing N]
	// to read the structure node data; `value of the node` `left child node index` `right child node index`
	// `for _, nodeIndexValue := range  make([]int, N)` does not work, as make([]int, N) only generates a slice N zero values
	// `for nodeIndexValue := 0; nodeIndexValue < N; nodeIndexValue++`` is equivalent to
	// `for nodeIndexValue, _ := range make([]int, N)` or `for nodeIndexValue := range make([]int, N)`

	for nodeIndexValue := range make([]int, N) { // we create a throw-away range of values [with N indices], which we then ranged over the indices
		var nodeValue, indexLeftChildNode, indexRightChildNode int
		fmt.Scanf("%d %d %d", &nodeValue, &indexLeftChildNode, &indexRightChildNode) // read row data from the *.in file
		nodes[nodeIndexValue].Value = nodeValue                                      // assign the Value of that node struct at index "indexValue"
		if indexLeftChildNode >= 0 {                                                 // check if the index is negative i.e. no child nodes exists
			nodes[nodeIndexValue].Left = &nodes[indexLeftChildNode] //assign a pointer to the child node at indexLeftChildNode
		}
		if indexRightChildNode >= 0 { // check if the index is negative i.e. no child nodes exists
			nodes[nodeIndexValue].Right = &nodes[indexRightChildNode] //assign a pointer to the child node at indexRightChildNode
		}
	}
	return nodes
}

// PrintData function helps to print the already constructed graph node values and child nodes
// it takes a single node of type *Node i.e. it does not take a slice of nodes
func PrintData(node *Node) {
	fmt.Print("Node Value: ", node.Value) // print the current node value
	if node.Left != nil {                 // check if the left child node even exists before printing
		fmt.Print(" , Left Child Node: ", node.Left.Value)
	}
	if node.Right != nil { // check if the right child node even exists before printing
		fmt.Print(" , Right Child Node: ", node.Right.Value)
	}
	fmt.Println() // print a new line

}
