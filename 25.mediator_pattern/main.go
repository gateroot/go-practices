package main

import . "go-practices/25.mediator_pattern/mediator"

func main() {
	saito := NewSaito()
	tanaka := &Tanaka{}
	sato := &Sato{}
	saito.AddColleague(tanaka)
	saito.AddColleague(sato)
	tanaka.SetMediator(saito)
	sato.SetMediator(saito)
	sato.Twitter("Tanaka", "Hello")
}
