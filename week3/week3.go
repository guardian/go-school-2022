package week3

import (
	"fmt"
	"io"
)

// Make the tests pass by running `go test ./week3` until it's happy.
//
// Use the cheatsheet to help:
// https://docs.google.com/document/d/1kHskYWl64DsMmI-dRKLKr-SN5B865up4tL7AoM2biHg/edit?usp=sharing.

// FizzBuzz prints consecutive numbers from 0 until upto (exclusive). Numbers
// divisible by 3 are replaced by 'Fizz' and numbers divisible by 5 replaced by
// 'Buzz'. For numbers divisible by both 3 and 5 'Fizz Buzz' is written.
//
// Hints: you can use strconv.Itoa to convert an int to a string. fmt.Fprintf is
// also a good option here. Consider using a switch to check the cases.
func FizzBuzz(w io.Writer, upto int) {
	fmt.Fprint(w, "TODO")
}

// Find returns the first element in items that satisfies a predicate function
// 'p'. bool is set to true if an element is found, or false if not.
//
// Hints: for/range is your friend here. You can use `var t T` to initialise the
// unknown type to its 'zero value'.
func Find[T any](items []T, p func(T) bool) (T, bool) {
	var t T
	return t, false // FIXME
}

type Item struct {
	Title string
	URL   string
}

// GetNews executes an HTTP GET request to :host/news, parses the JSON into a
// list of Items, and return it.
//
// - if the HTTP request fails, return an error containing the string 'http
// request failed' - if the JSON parsing fails, return an error containing the
// string 'json parsing failed'.
//
// (In real life, we would want to use better/custom errors here).
//
// Hints: use the `http` package for your GET request. Use the `encoding/json`
// package to 'Unmarshal' the response into a slice of the `Item` type defined
// above. Use the io.ReadAll helper to read the response body. &T (where T is
// some type) can be used to get a pointer for it.
func GetNews(host string) ([]Item, error) {
	return []Item{}, nil // TODO
}
