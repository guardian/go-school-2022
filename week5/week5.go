package week5

import (
	"context"
	"net/http"
	"net/url"
)

// FanIn should merge channels into a single read-only channel.
// Hint: use select.
func FanIn[T any](chs ...chan T) <-chan T {
	ch := make(chan T)
	return ch
}

// Throttle should return a function that, however often called, executes no
// more than once per wait time.
func Throttle[A, B any](fn func() A, tick chan B) func() A {
	return func() A {
		var a A
		return a
	}
}

// Cancellable wraps a function to provide cancellation (via a context). Note,
// this is NOT recommended in practice; functions themselves should accept a
// context directly (by convention as the first arg) so that they can perform
// any cleanup related to cancellation.
//
// Hint: make sure you *close* the channel upon cancel.
func Cancellable[A any](ctx context.Context, fn func() chan A) func() chan A {
	return func() chan A {
		ch := make(chan A)
		return ch
	}
}

type Result struct {
	Response http.Response
	Error    error
}

// GetQuickest should fetch each target/mirror in parallel and return a channel
// with the fastest response only or an error if all fail.
func GetQuickest(target1 url.URL, mirror1 url.URL, mirror2 url.URL) <-chan Result {
	ch := make(chan Result, 1)
	return ch
}

// FanOut should make an HTTP GET request to each target and push the raw
// response onto a (returned) channel. It should use workerCount as the number
// of concurrent workers to use.
func FanOut(targets []url.URL, workerCount int) <-chan []byte {
	ch := make(chan []byte)
	return ch
}
