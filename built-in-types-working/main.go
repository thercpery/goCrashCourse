package main

import (
	"fmt"
	"log"
	"sort"
)

// basic types (numbers, strings, booleans)
var myInt int // used most of the time

/*
These exist but in reality it was not recommended to use this.
var myInt16 int16
var myInt32 int32
var myInt64 int64
*/

var myUint uint // uint only holds positive values or zero

var myFloat float32
var myFloat64 float64 // for larger numbers

// aggregate types (array, struct)
type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Roy"
	log.Println(myString)

	myString = "John" // string are immutable in Golang.

	var myBool = true
	myBool = false

	log.Println(myBool)

	var myStrings [3]int

	myStrings[0] = 1
	myStrings[1] = 8
	myStrings[2] = 7

	fmt.Println("First element in array is", myStrings[0])

	var myCar Car

	myCar.NumberOfTires = 4
	myCar.Luxury = false
	myCar.Make = "Volkswagen"

	myCar2 := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My Car is a %d %s %s.\n", myCar2.Year, myCar2.Make, myCar2.Model)

	var myInt int
	myInt = 10
	fmt.Println(myInt)

	x := 10
	myFirstPointer := &x
	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "elephant")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals { // i is an index, and can be optional
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long.")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func changeValueOfPointer(num *int) {
	*num = 25
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
