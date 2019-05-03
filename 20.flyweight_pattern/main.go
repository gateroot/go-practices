package main

import (
	. "flyweight/flyweight"
	"fmt"
)

func main() {
	f := Factory
	あ := f.GetHumanLetter("あ")
	takeAPhoto(あ)
	い := f.GetHumanLetter("い")
	takeAPhoto(い)
	よ := f.GetHumanLetter("よ")
	takeAPhoto(よ)
	り := f.GetHumanLetter("り")
	takeAPhoto(り)
	あ2 := f.GetHumanLetter("あ")
	takeAPhoto(あ2)
	お := f.GetHumanLetter("お")
	takeAPhoto(お)
	し := f.GetHumanLetter("し")
	takeAPhoto(し)
}

func takeAPhoto(letter HumanLetter) {
	fmt.Println(letter.GetLetter())
}
