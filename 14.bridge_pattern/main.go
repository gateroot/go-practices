package main

import (
	"fmt"

	. "go-practices/14.bridge_pattern/bridge"
)

func main() {
	quickImpl := &QuickSortImpl{}
	bubbleImpl := &BubbleSortImple{}

	quickSorter := Sorter{quickImpl}
	bubbleSorter := Sorter{bubbleImpl}

	quickTimeSorter := TimeSorter{quickImpl}
	bubbleTimeSorter := TimeSorter{bubbleImpl}

	var values []int
	for i := 10000; i >= 0; i-- {
		values = append(values, i)
	}

	quickSorter.Sort(values)
	bubbleSorter.Sort(values)

	fmt.Println("[quick sort]")
	quickTimeSorter.Sort(values)
	fmt.Println("[bubble sort]")
	bubbleTimeSorter.Sort(values)
}
