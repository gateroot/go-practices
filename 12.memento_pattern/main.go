package main

import (
	"fmt"

	. "go-practices/12.memento_pattern/memento"
)

func main() {
	memory := make(map[string]*Memento)

	calc := Calc{}
	for i := 1; i <= 5; i++ {
		calc.Plus(i)
	}
	fmt.Println(calc.GetTemp())
	memory["5までの足し算"] = calc.CreateMemento()

	calc2 := Calc{}
	calc2.SetMemento(memory["5までの足し算"])
	for i := 6; i <= 10; i++ {
		calc.Plus(i)
	}
	fmt.Println(calc.GetTemp())
	memory["10までの足し算"] = calc.CreateMemento()
}
