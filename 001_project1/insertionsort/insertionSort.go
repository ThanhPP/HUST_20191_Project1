package insertionsort

import (
	nodeLib "001_project1/nodelibrary"
	"fmt"
)

//ReIndex indexing all the node in the linked list
func ReIndex(linkedList *nodeLib.LinkedList) {
	// t := time.Now()
	currentNode := linkedList.Head
	for i := 1; currentNode != nil; i++ {
		currentNode.Index = i
		currentNode = currentNode.NextNode
	}
	fmt.Println()
	// fmt.Println("Reindex time = ", time.Since(t))
}

// //InsertToHead remove a node and insert it to the head of the linked list
// func InsertToHead(linkedList *nodeLib.LinkedList, inputNode *nodeLib.Node) {
// 	if inputNode == linkedList.Head {
// 		return
// 	}

// 	linkedList.Head.BeforeNode = inputNode
// 	// fmt.Println("linkedList.Head.NextNode.BeforeNode =", inputNode)

// 	//remove Node
// 	inputNode.BeforeNode.NextNode = inputNode.NextNode
// 	if inputNode.NextNode != nil {
// 		inputNode.NextNode.BeforeNode = inputNode.BeforeNode
// 	}

// 	//change inputNode data
// 	inputNode.BeforeNode = nil
// 	inputNode.NextNode = linkedList.Head

// 	//change head

// 	linkedList.Head = inputNode

// }

//Insert insert the input node to the insert node positon
func Insert(linkedList *nodeLib.LinkedList, insertNode *nodeLib.Node, inputNode *nodeLib.Node) {
	if insertNode == inputNode || inputNode.NextNode == insertNode || inputNode == nil || insertNode == nil {
		return
	}

	//remove node
	if inputNode == linkedList.Head {
		linkedList.Head = inputNode.NextNode
		linkedList.Head = nil
	} else {
		inputNode.BeforeNode.NextNode = inputNode.NextNode
		if inputNode.NextNode != nil {
			inputNode.NextNode.BeforeNode = inputNode.BeforeNode
		}
	}

	if insertNode == nil {
		linkedList.AddNode(inputNode.Value)
	}

	if insertNode == linkedList.Head {
		linkedList.Head = inputNode
		inputNode.BeforeNode = nil
		inputNode.NextNode = insertNode

		insertNode.BeforeNode = inputNode
	} else {
		inputNode.BeforeNode = insertNode.BeforeNode
		insertNode.BeforeNode.NextNode = inputNode
		inputNode.NextNode = insertNode
		insertNode.BeforeNode = inputNode
	}

}

//Insert1 divide a linked list to 2 part : sorted and unsorted
func Insert1(linkedList *nodeLib.LinkedList) {
	var tempNode, currentNode, keyNode *nodeLib.Node
	currentNode = linkedList.Head
	keyNode = linkedList.Head
	flag := false
	for keyNode != nil {
		// fmt.Println("keyNode = ", keyNode)
		currentNode = keyNode
		keyNode = keyNode.NextNode
		tempNode = linkedList.Head
		// fmt.Println("--linkedlist.Head = ", linkedList.Head)
		// fmt.Println("--tempNode1 = ", tempNode)
		for (tempNode != keyNode) && (tempNode.Value <= currentNode.Value) {
			// fmt.Println("--tempNode2 = ", tempNode)
			flag = true
			tempNode = tempNode.NextNode
		}
		if !flag {
			if tempNode.Value >= currentNode.Value {
				// fmt.Println("-Inserted To Head")
				linkedList.InsertToHead(currentNode)
			} else {
				break
			}
		} else {
			if tempNode == nil {
				break
			}
			// fmt.Println("-Inserted")
			Insert(linkedList, tempNode, currentNode)
		}
	}
	ReIndex(linkedList)
	fmt.Println("Insertion sort : done")
}

// //SortedInsert add a node to sorted linklist
// func SortedInsert(sortedLinklist *nodeLib.LinkedList, inputNode *nodeLib.Node) {
// 	tempNode := &nodeLib.Node{
// 		Index:      0,
// 		Value:      0,
// 		BeforeNode: nil,
// 		NextNode:   nil,
// 	}
// 	var insertNode *nodeLib.Node
// 	var n int
// 	fmt.Println("sortedLinklist.Head = ", sortedLinklist.Head)
// 	if sortedLinklist.Head == nil {
// 		fmt.Println("tempNode = ", tempNode)
// 		tempNode.Value = inputNode.Value
// 		fmt.Println("tempNode.Value = ", tempNode.Value)
// 		sortedLinklist.Head = tempNode
// 	} else {
// 		if inputNode.Value <= sortedLinklist.Head.Value {
// 			tempNode = inputNode
// 			tempNode.BeforeNode = nil
// 			tempNode.NextNode = sortedLinklist.Head
// 			sortedLinklist.Head.BeforeNode = tempNode
// 			sortedLinklist.Head = tempNode
// 		} else {
// 			fmt.Println("Insert phase - find position")
// 			tempNode = sortedLinklist.Head
// 			for tempNode.NextNode != nil && tempNode.Value <= inputNode.Value {
// 				n++
// 				tempNode = tempNode.NextNode
// 			}
// 			if n > 0 {
// 				fmt.Println("Insert phase - insert")
// 				fmt.Printf("--tempNode = %+v \n", tempNode)
// 				insertNode = inputNode
// 				fmt.Printf("--insertNode = %+v \n", insertNode)
// 				//error

// 			}
// 		}
// 	}
// 	ReIndex(sortedLinklist)
// }

// //InsertionSort ...
// func InsertionSort(linkedList *nodeLib.LinkedList) {
// 	if linkedList.Cap < 2 {
// 		return
// 	}

// 	//create temp list
// 	var sortedList nodeLib.LinkedList
// 	sortedList.Head = linkedList.Head
// 	sortedList.Head.NextNode = nil

// 	var tempNode *nodeLib.Node
// 	var keyNode *nodeLib.Node
// 	keyNode = sortedList.Head
// 	tempNode = sortedList.Head
// 	var currentNode *nodeLib.Node
// 	currentNode = linkedList.Head

// 	var flag bool
// 	for currentNode.NextNode != nil {
// 		if currentNode.Value < keyNode.Value {
// 			InsertToHead(&sortedList, currentNode)
// 		} else {
// 			flag = false
// 			tempNode = keyNode
// 			for tempNode != nil && currentNode.Value > tempNode.Value {
// 				tempNode = tempNode.NextNode
// 				flag = true
// 			}
// 			if flag {
// 				if tempNode == nil {

// 				}
// 			}
// 		}
// 		currentNode = currentNode.NextNode
// 	}

// 	ReIndex(linkedList)
// }
