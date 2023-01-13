package main

import "fmt"

func main() {
	//even name variable using const, stil changed by variable name in increment func
	name := "Eko"
	counter := 0

	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
