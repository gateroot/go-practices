package strategy

type Human struct {
	name   string
	height int
	weight int
	age    int
}

func NewHuman(name string, height, weight, age int) Human {
	return Human{name, height, weight, age}
}
