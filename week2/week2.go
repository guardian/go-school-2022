package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 1.8) + 32)
}

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
	fmt.Println(`a "quote" like this.`)

	arr := [2]string{"Emily Dickinson", "Fyodor Dostoevsky"}
	fmt.Println(arr)

	sli := []string{"Emily Dickinson", "Fyodor Dostoevsky"}
	sli = append(sli, "Jane Austen")
	fmt.Println(sli)
	fmt.Printf("%T\n", sli)

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
	if age >= 18 {
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

	c := Celsius(100)
	//ToFahrenheit(c)
	c.ToFahrenheit()
}
