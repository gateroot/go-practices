package main

import . "chain_of_responsibility/chain_of_responsibility"

func main() {
	hand1 := NewConcreteHandlerA("A", 1, 10)
	hand2 := NewConcreteHandlerA("B", 11, 20)
	hand3 := NewConcreteHandlerA("C", 21, 30)
	hand4 := NewConcreteHandlerB("D")
	//hand5 := NewConcreteHandlerA("E", 31, 40)
	hand1.SetNext(hand2).SetNext(hand3).SetNext(hand4)//.SetNext(hand5)
	hand1.Request(31)
}
