package bubblesort

import (
	insertionSort "001_project1/insertionsort"
	nodeLib "001_project1/nodelibrary"
	"fmt"
)

//BubbleSort ...
func BubbleSort(linkedList *nodeLib.LinkedList) {
	var tempNode, keyNode *nodeLib.Node
	if linkedList.Cap > 2 {
		swapLimit := linkedList.Cap
		for swapLimit > 0 {
			tempNode = linkedList.Head
			for tempNode.NextNode != nil {
				if tempNode.Value > tempNode.NextNode.Value {
					linkedList.SwapNode(tempNode, tempNode.NextNode)
				} else {
					tempNode = tempNode.NextNode
				}
			}
			swapLimit--
		}
	} else if linkedList.Cap == 2 {
		keyNode = linkedList.Head.NextNode
		tempNode = linkedList.Head
		if keyNode.Value < tempNode.Value {
			linkedList.SwapNode(keyNode, tempNode)
		}
	} else {
		fmt.Println("Can't sort")
		return
	}
	insertionSort.ReIndex(linkedList)
	fmt.Println("Bubble sort : done")
}
