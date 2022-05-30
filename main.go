package main

import "fmt"

func main() {

	// Questions when learning a language:
	// - primitives
	// - collections
	// - control structures
	// - scope
	// - paradigms
	// - ...

	fmt.Println(1 + 2)
	fmt.Printf("%T\n", 1+2)
	fmt.Printf("%T\n", true)
	fmt.Println("A string" + " and another string")

	arr := [2]string{"Emily Dickinson", "Fyodor Dostoevsky"}
	fmt.Println(arr)

	sli := []string{"Emily Dickinson", "Fyodor Dostoevsky"}
	sli = append(sli, "Jane Austen")
	fmt.Println(sli)

	for _, author := range sli {
		fmt.Printf("One of my favourite authors is: %s\n", author)
	}

	fmt.Println(len(sli))

	authors := map[string]string{
		"Emily Dickinson":   "Middlemarch",
		"Fyodor Dostoevsky": "Crime and Punishment",
		"Jane Austen":       "Pride and Prejudice",
	}

	fmt.Println(authors)

	// Custom type + Method (see Celsius example)

	age := 18
	if age > 18 {
		fmt.Println("Can vote!")
	} else {
		fmt.Println("Too young!")
	}

	switch {
	case age > 18:
		fmt.Println("Can vote!")
	default:
		fmt.Println("Too young!")
	}

	type Person struct {
		Age  int
		Name string
	}

	me := Person{Age: 34, Name: "Nic"}
	fmt.Println(me)

	// Scope is block scope
	// And also package scope using Capitalisation.
	// That's it!

	// Task! Create a map of a few of your favourite authors and books and iterate over it to print them out
	// Task! Capitalise a string. Nb. I want you to do this using the 'unicode.ToUpper' function.
	// Task! write fizzbuzz in Go. (https://en.wikipedia.org/wiki/Fizz_buzz)
}

type Celsius float64
type Fahrenheit = float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 1.8) + 32)
}
