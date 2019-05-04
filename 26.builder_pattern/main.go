package main

import (
	"fmt"

	. "go-practices/26.builder_pattern/builder"
)

func main() {
	builder := NewSaltWaterBuilder()
	director := NewDirector(builder)
	director.Construct()
	saltWater := builder.GetResult()
	fmt.Println(saltWater)
}
