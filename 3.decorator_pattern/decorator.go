package main

import "fmt"

type Icecream interface {
	GetName() string
	HowSweet() string
}

type VanillaIcecream struct {}

func (v VanillaIcecream) GetName() string {
	return "バニラアイスクリーム"
}

func (v VanillaIcecream) HowSweet() string {
	return "バニラ味"
}

type GreenTeaIcecream struct {}

func (g GreenTeaIcecream) GetName() string {
	return "抹茶アイスクリーム"
}

func (g GreenTeaIcecream) HowSweet() string {
	return "抹茶味"
}

type CasheNutsToppingIcecream struct {
	Icecream Icecream
}

func (c CasheNutsToppingIcecream) GetName() string {
	return "カシューナッツ" + c.Icecream.GetName()
}

func (c CasheNutsToppingIcecream) HowSweet() string {
	return c.Icecream.HowSweet()
}

func main() {
	for _, i := range []Icecream{VanillaIcecream{}, GreenTeaIcecream{}, CasheNutsToppingIcecream{VanillaIcecream{}}} {
		fmt.Printf("%s: %s\n", i.GetName(), i.HowSweet())
	}
}