package main

import (
	. "template_method/template_method"
)

func main() {
	tanaka := TanakaWoodCutPrint{}
	template := WoodCutPrintTemplate{tanaka}
	template.CreateWookCutPrint()
}
