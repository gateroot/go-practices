package main

type Factory interface {
	getSoup() Soup
	getMain() Protein
	getVegetables() []Vegetable
	getOtherIngredient() []Ingredient
}

type MizutakiFactory struct {}

func (mf MizutakiFactory) getSoup() Soup {
	return new(ChickenBoneSoup)
}

func (mf MizutakiFactory) getMain() Protein {
	return new(Chicken)
}

func (mf MizutakiFactory) getVegetables() []Vegetable {
	return []Vegetable{
		new(ChineseCabbage),
		new(Leek),
		new(Chrysanthemum),
	}
}

func (mf MizutakiFactory) getOtherIngredient() []Ingredient {
	return []Ingredient{
		new(Tofu),
	}
}
