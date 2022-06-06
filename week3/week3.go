package week3

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// make the tests pass

// FizzBuzz prints consecutive numbers from 1 until upto (exclusive). Numbers
// divisible by 3 are replaced by 'Fizz' and numbers divisible by 5 replaced by
// 'Buzz'. For numbers divisible by both 3 and 5 'Fizz Buzz' is written.
func FizzBuzz(w io.Writer, upto int) {
	for i := 1; i < upto; i++ {
		prn := strconv.Itoa(i)

		switch {
		case i%3 == 0 && i%5 == 0:
			prn = "Fizz Buzz"
		case i%3 == 0:
			prn = "Fizz"
		case i%5 == 0:
			prn = "Buzz"
		}

		w.Write([]byte(prn + "\n"))
	}
}

// Find returns the first element in items that satisfies a predicate function
// 'p'.
func Find[T any](items []T, p func(T) bool) (T, bool) {
	for _, item := range items {
		if p(item) {
			return item, true
		}
	}

	var t T
	return t, false
}

type Item struct {
	Title string
	URL   string
}

// GetNews executes an HTTP GET request to :host/news, parses the JSON into a
// list of Items, and return it.
//
// - if the HTTP request fails or returns a non-200 response, return an error
//   containing the string 'http request failed'
// - if the JSON parsing fails, return an error containing the
//   string 'json parsing failed'.
//
// (In real life, we would want to use better/custom errors here).
//
// Hint: use io.ReadAll to read the resp.Body and also defer resp.Body.Close()
// to ensure we close the connection.
func GetNews(host string) ([]Item, error) {
	resp, err := http.Get(host + "/news")
	if err != nil {
		return nil, errors.New("http request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http request failed")
	}

	raw, _ := io.ReadAll(resp.Body)
	var items []Item

	err = json.Unmarshal(raw, &items)
	if err != nil {
		return nil, errors.New("json parsing failed")
	}

	return items, nil
}
