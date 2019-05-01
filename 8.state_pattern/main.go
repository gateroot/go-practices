package main

import (
	"fmt"
	"state/state"
)

func main() {
	yumi := &state.StatePatternYumichan{}

	badMood := &state.BadMoodState{}
	ordinary := &state.OrdinaryState{}

	fmt.Println("[BAD MOOD]")
	yumi.ChangeState(badMood)
	fmt.Println(yumi.MorningGreet())
	fmt.Println(yumi.GetProtectionForCold())

	fmt.Println("[ORDINARY]")
	yumi.ChangeState(ordinary)
	fmt.Println(yumi.MorningGreet())
	fmt.Println(yumi.GetProtectionForCold())
}
