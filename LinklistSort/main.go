package main

import (
	bubbleSort "001_project1/bubblesort"
	insertionSort "001_project1/insertionsort"
	nodeLib "001_project1/nodelibrary"
	selectionSort "001_project1/selectionsort"
	"fmt"
	"sync"
	"time"
)

func genLinkList(linkList *nodeLib.LinkedList, cap int) {
	var zero = struct{}{}
	m := make(map[int]struct{}, cap+1)
	for i := 1; i <= cap; i++ {
		m[i] = zero
	}
	for i := range m {
		linkList.AddNode(i)
	}
}

func main() {
	var insertionSortLinkList nodeLib.LinkedList
	var bubbleSortSortLinkList nodeLib.LinkedList
	var selectionSortLinkList nodeLib.LinkedList
	var createNodeWg sync.WaitGroup

	fmt.Println("________________________________________Create Node________________________________________")
	fmt.Println()
	capacity := 1000
	linkListCounter := 0

	createNodeWg.Add(3)
	createNodeTimer := time.Now()
	go func() {
		genLinkList(&insertionSortLinkList, capacity)
		linkListCounter++
		createNodeWg.Done()
	}()

	go func() {
		genLinkList(&bubbleSortSortLinkList, capacity)
		linkListCounter++
		createNodeWg.Done()
	}()

	go func() {
		genLinkList(&selectionSortLinkList, capacity)
		linkListCounter++
		createNodeWg.Done()
	}()
	linkListCounter++
	createNodeWg.Wait()

	fmt.Printf("\tCreate %d doubly linked list : %d nodes/each ---- in time : %v \n", linkListCounter, capacity, time.Since(createNodeTimer))
	// selectionSortLinkList.ShowAll()
	fmt.Println("______________________________________________________________________________________________")
	fmt.Println()

	var wg sync.WaitGroup
	var insertionTime time.Duration
	var bubbleSortTime time.Duration
	var selectionSortTime time.Duration
	wg.Add(3)
	go func() {
		fmt.Println("start insertion sort")
		insertionTimer := time.Now()
		insertionSort.Insert1(&insertionSortLinkList)
		insertionTime = time.Since(insertionTimer)
		wg.Done()
	}()
	go func() {
		fmt.Println("start bubble sort")
		bubbleSortTimer := time.Now()
		bubbleSort.BubbleSort(&bubbleSortSortLinkList)
		bubbleSortTime = time.Since(bubbleSortTimer)
		wg.Done()
	}()
	go func() {
		fmt.Println("start selection sort")
		selectionSortTimer := time.Now()
		selectionSort.SelectionSort(&selectionSortLinkList)
		selectionSortTime = time.Since(selectionSortTimer)
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("________________________________________Insertion Sort________________________________________")
	fmt.Println()
	fmt.Printf("\tRun insertion sort with capacity : %d (with reindexing) ---- in time : %v \n", insertionSortLinkList.Cap, insertionTime)
	// insertionSortLinkList.ShowAll()
	fmt.Println("______________________________________________________________________________________________")
	fmt.Println()

	fmt.Println("________________________________________Bubble Sort________________________________________")
	fmt.Println()
	fmt.Printf("\tRun bubble sort with capacity : %d (with reindexing) ---- in time : %v \n", bubbleSortSortLinkList.Cap, bubbleSortTime)
	// bubbleSortSortLinkList.ShowAll()
	fmt.Println("______________________________________________________________________________________________")
	fmt.Println()

	fmt.Println("________________________________________Selection Sort________________________________________")
	fmt.Println()
	fmt.Printf("\tRun selection sort with capacity : %d (with reindexing) ---- in time : %v \n", selectionSortLinkList.Cap, selectionSortTime)
	// selectionSortLinkList.ShowAll()
	fmt.Println("______________________________________________________________________________________________")
	fmt.Println()
}
