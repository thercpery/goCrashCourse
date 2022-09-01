package main

import "log"

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
}
