package main

import "fmt"

// quoted from [https://www.techscore.com/tech/DesignPattern/AbstractFactory.html/]

type HotPot struct {
	Pot Pot
	Soup Soup
	Main Protein
	Vegetables []Vegetable
	OtherIngredients []Ingredient
}

func (hp HotPot) Print() {
	fmt.Printf("Pot: %s\nSoup: %s\nMain: %s\n", hp.Pot, hp.Soup, hp.Main)
	for i, vegetable := range hp.Vegetables {
		fmt.Printf("Vegetables[%d]: %s\n", i, vegetable)
	}
	for i, ingredient := range hp.OtherIngredients {
		fmt.Printf("Ingredients[%d]: %s\n", i, ingredient)
	}
}

type Pot struct {}

func (p Pot) String() string {
	return "Pot"
}

type Soup interface {
	String() string
}

type ChickenBoneSoup struct {}

func (cbs ChickenBoneSoup) String() string {
	return "Chicken Bone Soup"
}

type Protein interface {
	String() string
}

type Chicken struct {}

func (c Chicken) String() string {
	return "Chicken"
}

type Vegetable interface {
	String() string
}

type ChineseCabbage struct {}

func (cc ChineseCabbage) String() string {
	return "Chinese Cabbage"
}

type Leek struct {}

func (l Leek) String() string {
	return "Leek"
}

type Chrysanthemum struct {}

func (c Chrysanthemum) String() string {
	return "Chrysanthemum"
}

type Ingredient interface {
	String() string
}

type Tofu struct {}

func (t Tofu) String() string {
	return "Tofu"
}

func main() {
	var factory Factory = new(MizutakiFactory)
	hotPot := HotPot{
		Pot: Pot{},
		Soup: factory.getSoup(),
		Main: factory.getMain(),
		Vegetables: factory.getVegetables(),
		OtherIngredients: factory.getOtherIngredient(),
	}
	hotPot.Print()
}


