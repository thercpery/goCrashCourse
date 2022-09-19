package main

import "fmt"

func main() {
	age := 10           // this is not an expression but a literal
	name := "Jack"      // this is also not an expression but a literal
	rightHanded := true // this is a boolean literal

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded) // now that's an expression
	fmt.Println()

	ageInTenYears := age + 10 // this is an expression

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears) // this is an expression
	fmt.Println()

	isATeenager := age >= 13 // this is an expression

	fmt.Println(name, "is a teenager:", isATeenager) // this is an expression

}
