package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, world!")

	// Variables - DO NOT DECLARE VARIABLES AND NOT USE THEM OR GOLANG WILL RETURN AN ERROR

	// keyword "var"
	/* var whatToSay string
	whatToSay = "Hello World!" */

	// shorthand variables/assignment operator
	/* whatToSay := "Hello World!"

	sayHelloWorld(whatToSay) */

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}
}

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
