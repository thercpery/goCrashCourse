package main

import (
	"fmt"
	"log"
	"sort"
	// "github.com/eiannone/keyboard"
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
type Animal interface {
	Says() string
	HowManyLegs() int
}

// Dog is the type for dogs
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

/* type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}
*/
/* func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumberOfLegs)
}
*/
var keyPressChan chan rune

func main() {
	/* 	go doSomething("Hello, world")
	   	fmt.Println("This is another message")

	   	for {
	   		// do nothing
	   	}
	*/

	/* keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit.")
	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
	*/
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

	for i, x := range animals { // i is an index, and can be optional (_)
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

	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete(intMap, "four")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	el, ok := intMap["four"]

	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	intMap["two"] = 4
	fmt.Println(intMap["two"])

	z := addTwoNumbers(2, 4)
	fmt.Println(z)

	myTotal := sumMany(2, 3, 4, 5, 88, 7, -5)
	fmt.Println(myTotal)

	/* var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4

	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs()
	*/

	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(&cat)

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

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}

func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is", s)
		until = until + 1
		if until >= 5 {
			break
		}
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says %s and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
