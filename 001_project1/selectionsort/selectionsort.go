package selectionsort

import (
	insertionSort "001_project1/insertionsort"
	nodeLib "001_project1/nodelibrary"
	"fmt"
)

func findMinNode(inputNode *nodeLib.Node) (minNode *nodeLib.Node, flag bool) {
	minNode = inputNode
	flag = false
	for inputNode != nil {
		if inputNode.Value < minNode.Value {
			minNode = inputNode
			flag = true
		}
		inputNode = inputNode.NextNode
	}

	return minNode, flag
}

//SelectionSort ...
func SelectionSort(linkedList *nodeLib.LinkedList) {
	var currentNode, keyNode *nodeLib.Node
	var minNode *nodeLib.Node
	var found bool
	if linkedList.Cap > 2 {
		keyNode = linkedList.Head.NextNode
		for keyNode != nil {
			currentNode = keyNode.BeforeNode
			keyNode = keyNode.NextNode
			minNode, found = findMinNode(currentNode)
			// fmt.Println(found, "  ", minNode)
			if found {
				// fmt.Println("found")
				// fmt.Println("swap ", currentNode, " and ", minNode)
				linkedList.SwapNode(minNode, currentNode)
				// linkedList.ShowAll()
			}
		}
	} else if linkedList.Cap == 2 {
		keyNode = linkedList.Head.NextNode
		currentNode = linkedList.Head
		if currentNode.Value > keyNode.Value {
			linkedList.SwapNode(currentNode, keyNode)
		}
	} else {
		fmt.Println("Can't sort")
		return
	}
	insertionSort.ReIndex(linkedList)
	fmt.Println("Selection sort : done")
	fmt.Println()
}
