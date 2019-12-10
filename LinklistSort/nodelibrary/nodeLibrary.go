package nodelibrary

import "fmt"

//Node part of a linked list
type Node struct {
	BeforeNode *Node
	NextNode   *Node
	Index      int
	Value      int
}

//LinkedList ....
type LinkedList struct {
	Head *Node
	Cap  int
}

//AddNode create a new Node and add to the linked list
func (linkedList *LinkedList) AddNode(inputValue int) *Node {
	newNode := &Node{
		Index: 1,
		Value: inputValue,
	}
	if linkedList.Head == nil {
		linkedList.Head = newNode
		linkedList.Cap = 1
	} else {
		currentNode := linkedList.Head
		for currentNode.NextNode != nil {
			currentNode = currentNode.NextNode
		}
		newNode.Index = linkedList.Cap + 1
		newNode.BeforeNode = currentNode
		currentNode.NextNode = newNode
		linkedList.Cap++
	}
	return newNode
}

//ShowAll print all data from a linked list
func (linkedList *LinkedList) ShowAll() {
	if linkedList.Head == nil {
		fmt.Println("empty list")
	} else {
		currentNode := linkedList.Head
		fmt.Println("1")
		fmt.Printf("\t index = %v \n\t value = %v \n\t\t beforeNode = %+v \n\t\t nextNode = %+v \n", currentNode.Index, currentNode.Value, currentNode.BeforeNode, currentNode.NextNode)
		for i := 2; currentNode.NextNode != nil; i++ {
			fmt.Println(i)
			currentNode = currentNode.NextNode
			fmt.Printf("\t index = %v \n\t value = %v \n\t\t beforeNode = %+v \n\t\t nextNode = %+v \n", currentNode.Index, currentNode.Value, currentNode.BeforeNode, currentNode.NextNode)
			// if i > 5 {
			// 	return
			// }
		}
	}
}

//InsertToHead ...
func (linkedList *LinkedList) InsertToHead(inputNode *Node) {
	linkedList.Head.BeforeNode = inputNode

	//change Node before inputNode
	if inputNode.BeforeNode != nil {
		inputNode.BeforeNode.NextNode = inputNode.NextNode
	}

	//change Node after Node
	if inputNode.NextNode != nil {
		inputNode.NextNode.BeforeNode = inputNode.BeforeNode
	}

	//change inputNode
	inputNode.BeforeNode = nil
	inputNode.NextNode = linkedList.Head

	//change head
	linkedList.Head = inputNode

	//re-index
	currentNode := linkedList.Head
	for i := 1; i <= linkedList.Cap; i++ {
		currentNode.Index = i
		currentNode = currentNode.NextNode
	}
}

//SwapNode swap the position between 2 node in a linked list
func (linkedList *LinkedList) SwapNode(inputNode1 *Node, inputNode2 *Node) {
	if inputNode1 == inputNode2 || inputNode1 == nil || inputNode2 == nil {
		return
	}

	//check index
	if inputNode1.Index > inputNode2.Index {
		tempNode := inputNode1
		inputNode1 = inputNode2
		inputNode2 = tempNode
	}
	var tempNode Node
	tempNode = *inputNode1

	//if head
	if inputNode1 == linkedList.Head {
		if inputNode1.NextNode != inputNode2 {
			if inputNode2.NextNode != nil {
				inputNode2.NextNode.BeforeNode = inputNode1
			}
			inputNode1.NextNode.BeforeNode = inputNode2
			inputNode1.NextNode = inputNode2.NextNode
			inputNode1.BeforeNode = inputNode2.BeforeNode
			inputNode1.Index = inputNode2.Index

			inputNode2.BeforeNode.NextNode = inputNode1
			inputNode2.NextNode = tempNode.NextNode
			inputNode2.BeforeNode = tempNode.BeforeNode
			inputNode2.Index = tempNode.Index
			linkedList.Head = inputNode2
		} else {
			if inputNode2.NextNode != nil {
				inputNode2.NextNode.BeforeNode = inputNode1
			}
			inputNode1.NextNode = inputNode2.NextNode
			inputNode1.BeforeNode = inputNode2
			inputNode1.Index = inputNode2.Index

			inputNode2.NextNode = inputNode1
			inputNode2.BeforeNode = tempNode.BeforeNode
			inputNode2.Index = tempNode.Index
			linkedList.Head = inputNode2
		}
		return
	}

	//if not head
	if inputNode1.NextNode == inputNode2 {
		// fmt.Println("inputNode1.NextNode == inputNode2")
		// fmt.Println(tempNode)
		if inputNode2.NextNode != nil {
			inputNode2.NextNode.BeforeNode = inputNode1
		}
		inputNode1.BeforeNode.NextNode = inputNode2
		inputNode1.BeforeNode = inputNode2
		inputNode1.NextNode = inputNode2.NextNode
		inputNode1.Index = inputNode2.Index

		inputNode2.NextNode = inputNode1
		inputNode2.BeforeNode = tempNode.BeforeNode
		inputNode2.Index = tempNode.Index

		return
	}

	if inputNode2.NextNode != nil {
		inputNode2.NextNode.BeforeNode = inputNode1
	}
	inputNode1.BeforeNode.NextNode = inputNode2
	inputNode2.BeforeNode.NextNode = inputNode1
	inputNode1.NextNode.BeforeNode = inputNode2

	inputNode1.BeforeNode = inputNode2.BeforeNode
	inputNode1.NextNode = inputNode2.NextNode
	inputNode1.Index = inputNode2.Index

	inputNode2.NextNode = tempNode.NextNode
	inputNode2.BeforeNode = tempNode.BeforeNode
	inputNode2.Index = tempNode.Index
	return

	// //change the head of the linked list or node before node 1
	// beforeNode1 := inputNode1.BeforeNode
	// if inputNode1 == linkedList.Head {
	// 	linkedList.Head = inputNode2
	// } else {
	// 	//change the pointer to first node
	// 	beforeNode1.NextNode = inputNode2
	// }

	// //change the pointer of the node before node 2
	// if inputNode2.Index-inputNode1.Index > 1 {
	// 	inputNode2.BeforeNode.NextNode = inputNode1
	// }
	// if inputNode2.BeforeNode.BeforeNode == inputNode1 {
	// 	inputNode2.BeforeNode.BeforeNode = inputNode2
	// }

	// //change the pointer of the after node 2
	// if inputNode2.NextNode != nil {
	// 	inputNode2.NextNode.BeforeNode = inputNode1
	// }

	// //swap pointer between two nodes
	// if inputNode2.Index-inputNode1.Index == 1 {
	// 	var tempNode Node
	// 	tempNode.BeforeNode = inputNode1.BeforeNode
	// 	tempNode.NextNode = inputNode1.NextNode
	// 	tempNode.Index = inputNode1.Index

	// 	inputNode1.BeforeNode = inputNode2
	// 	inputNode1.NextNode = inputNode2.NextNode
	// 	inputNode1.Index = inputNode2.Index

	// 	inputNode2.BeforeNode = tempNode.BeforeNode
	// 	inputNode2.NextNode = inputNode1
	// 	inputNode2.Index = tempNode.Index
	// } else {
	// 	var tempNode Node
	// 	tempNode.BeforeNode = inputNode1.BeforeNode
	// 	tempNode.NextNode = inputNode1.NextNode
	// 	tempNode.Index = inputNode1.Index
	// 	// fmt.Printf("temp node : %+v \n", tempNode)

	// 	inputNode1.BeforeNode = inputNode2.BeforeNode
	// 	inputNode1.NextNode = inputNode2.NextNode
	// 	inputNode1.Index = inputNode2.Index

	// 	inputNode2.BeforeNode = tempNode.BeforeNode
	// 	inputNode2.NextNode = tempNode.NextNode
	// 	inputNode2.Index = tempNode.Index
	// }
}

// //FindNode find a node with given index
// func (linkedList *LinkedList) FindNode(index int) *Node {
// 	if linkedList.Head == nil || linkedList.Cap < index || index < 1 {
// 		fmt.Println("Cannot find")
// 		return nil
// 	}
// 	currentNode := linkedList.Head
// 	if currentNode.Index == index {
// 		return currentNode
// 	}
// 	for currentNode.NextNode != nil {
// 		currentNode = currentNode.NextNode
// 		if currentNode.Index == index {
// 			return currentNode
// 		}
// 	}
// 	return nil
// }

// //InsertionSort ...
// func (linkedList *LinkedList) InsertionSort() {
// 	var currentNode *Node
// 	currentNode = linkedList.Head
// 	fmt.Println(currentNode)
// }

// //InsertNode ...
// func (linkedList *LinkedList) InsertNode(index int, inputNode *Node) {
// 	if inputNode == linkedList.Head {
// 		if index == 1 {
// 			return
// 		}

// 		//change head
// 		linkedList.Head = inputNode.NextNode
// 		inputNode.NextNode.BeforeNode = nil

// 		//change node before index node
// 		oldIndexNode := linkedList.FindNode(index)
// 		oldIndexNode.BeforeNode.NextNode = inputNode

// 		//change node after index node
// 		if oldIndexNode.NextNode != nil {
// 			oldIndexNode.NextNode.BeforeNode = oldIndexNode.BeforeNode
// 		}
// 	}

// 	if index == 1 {
// 		linkedList.InsertToHead(inputNode)
// 	} else {
// 		//remove node
// 		inputNode.BeforeNode.NextNode = inputNode.NextNode
// 		if inputNode.NextNode != nil {
// 			inputNode.NextNode.BeforeNode = inputNode.BeforeNode
// 		}

// 		//change node data
// 		oldIndexNode := linkedList.FindNode(index)
// 		inputNode.BeforeNode = oldIndexNode.BeforeNode
// 		inputNode.NextNode = oldIndexNode

// 		//insert node
// 		oldIndexNode.BeforeNode.NextNode = inputNode
// 		oldIndexNode.BeforeNode = inputNode

// 		//re-index
// 		if index < inputNode.Index {
// 			currentNode := linkedList.Head
// 			for i := 1; i <= linkedList.Cap; i++ {
// 				currentNode.Index = i
// 				currentNode = currentNode.NextNode
// 			}
// 		}
// 	}
// }
