package main

import "log"

func main() {
	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay, _ = saySomething("Hello")
	saySomethingElse, _ = saySomething("Good Bye")

	log.Println(whatToSay)
	log.Println(saySomethingElse)
	log.Println(saySomething("Finally")) // No need to put in variable

	i = 7
	i = 8
	log.Println(i)

}

func saySomething(s string) (string, string) {
	return s, "World"

}
