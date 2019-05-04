package main

import (
	"fmt"

	. "go-practices/24.prototype_pattern/prototype"
)

func main() {
	protoTypeMap := make(map[string]Cloneable)
	keeper := NewPrototypeKeeper(protoTypeMap)
	p1 := Paper{"Hello"}
	keeper.AddCloneable("hello", p1)
	p2 := keeper.GetClone("hello").(Paper)
	p1.Name = "World"
	fmt.Println(p2.Name, p1.Name)
}
