package template_method

import "fmt"

type WoodCutPrint interface {
	Draw()
	Cut()
	Print()
}

type WoodCutPrintTemplate struct {
	WoodCutPrint WoodCutPrint
}

func (w WoodCutPrintTemplate) Draw() {
	w.WoodCutPrint.Draw()
}

func (w WoodCutPrintTemplate) Cut() {
	w.WoodCutPrint.Cut()
}

func (w WoodCutPrintTemplate) Print() {
	w.WoodCutPrint.Print()
}

func (w WoodCutPrintTemplate) CreateWookCutPrint() {
	w.Draw()
	w.Cut()
	w.Print()
}

type TanakaWoodCutPrint struct{}

func (w TanakaWoodCutPrint) Draw() {
	fmt.Println("板材にマジックを使って大好きな女の子の絵を描く")
}

func (w TanakaWoodCutPrint) Cut() {
	fmt.Println("彫刻刀を使って細部まで丁寧に板材を彫る")
}

func (w TanakaWoodCutPrint) Print() {
	fmt.Println("版画インクと馬簾を使って豪快にプリントする")
}
