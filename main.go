package main

import (
	"fmt"
)

func main() {
	var (
		one   = "1"
		two   = "2"
		three = "3"
	)

	// fs.ReadDir(fs.Fs, "hello.txt")
	// --------------------Types of date ---------------------- //
	str := "String"
	intiger := 10

	floats := 9.99

	boolean := true
	// ------------------------- // ---------------- //

	// Controls

	if str == "String" {
		fmt.Println("Is equal")
	} else {
		fmt.Println("No is equal")
	}

	// Types avanced

	lists := [3]int{1, 2, 3}

	var lista_second []int

	lista_second = append(lista_second, 1)
	lista_second = append(lista_second, 2)
	lista_second = append(lista_second, 3)
	lista_second = append(lista_second, 4)

	type Pet struct {
		Name     string
		Nickname string
	}

	lola := Pet{
		Name:     "Lola",
		Nickname: "Lolancia",
	}

	fmt.Println(lola)
	fmt.Println(lola.Name)
	fmt.Println(lola.Nickname)
	fmt.Println("Hello word")
	fmt.Println(lists)
	fmt.Println(lista_second)
	fmt.Println(one, two, three)

	fmt.Println(str, intiger, floats, boolean)
	great_people()
}

func great_people() {
	fmt.Println("Hi people!!!")
}
