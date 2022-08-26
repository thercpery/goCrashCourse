package packageone

import "fmt"

var PackageVar = "packageone's package level variable"

func PrintMe(myVar string, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
