package bridge

import (
	"fmt"
	"time"
)

type TimeSorter struct {
	SortImpl
}

func (sorter *TimeSorter) Sort(values []int) []int {
	start := time.Now()
	sorted := sorter.SortImpl.Sort(values)
	end := time.Now()
	fmt.Println("time: ", end.Sub(start).Nanoseconds(), "ns")
	return sorted
}
