package main

import (
	"myapp/packageone"
)

var myVar = "package-level variable"

func main() {
	var blockVar = "block-level variable"
	packageone.PrintMe(myVar, blockVar)
}
