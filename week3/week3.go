package week3

import "io"

// make the tests pass

// FizzBuzz prints consecutive numbers from 0 until upto (exclusive). Numbers
// divisible by 3 are replaced by 'Fizz' and numbers divisible by 5 replaced by
// 'Buzz'. For numbers divisible by both 3 and 5 'Fizz Buzz' is written.
func FizzBuzz(w io.Writer, upto int) {
	w.Write([]byte("TODO"))
}

// Find returns the first element in items that satisfies a predicate function
// 'p'.
func Find[T any](items []T, p func(T) bool) T {
	return items[0] // FIXME
}

type Item struct {
	Title string
	URL   string
}

// GetNews executes an HTTP GET request to :host/news, parses the JSON into a list of Items, and return it.
//
// - if the HTTP request fails, return an error containing the string 'http request failed'
// - if the JSON parsing fails, return an error containing the string 'json parsing failed'.
//
// (In real life, we would want to use better/custom errors here).
func GetNews(host string) ([]Item, error) {
	return []Item{}, nil // TODO
}
