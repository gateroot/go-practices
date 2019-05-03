package main

import (
	"fmt"
	. "strategy/strategy"
)

type MyClass struct {
	comparator Comparator
}

func (c MyClass) Compare(h1, h2 Human) int {
	return c.comparator.Compare(h1, h2)
}

func main() {
	h1 := NewHuman("John", 172, 82, 25)
	h2 := NewHuman("Bob", 179, 63, 32)

	ageComparator := AgeComparator{}
	heightComparator := HeightComparator{}
	weightComparator := WeightComparator{}

	myClass := MyClass{&ageComparator}
	fmt.Println("age: ", myClass.Compare(h1, h2))
	myClass.comparator = &heightComparator
	fmt.Println("height: ", myClass.Compare(h1, h2))
	myClass.comparator = &weightComparator
	fmt.Println("weight: ", myClass.Compare(h1, h2))
}
