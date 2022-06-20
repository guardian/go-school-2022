package week5

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	var chans []chan int

	for i := 0; i < 3; i++ {
		ch := make(chan int, 1)
		ch <- i
		chans = append(chans, ch)
	}

	out := FanIn(chans...)

	got := []int{}

	// TODO this will block forever if the function is not implemented
	// correctly so perhaps add timeout.
	for i := 0; i < 3; i++ {
		item, ok := readCh(out)
		if !ok {
			t.Error("timeout reading from result channel")
			return
		}
		got = append(got, item)
	}

	want := []int{0, 1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestThrottle(t *testing.T) {
	i := 0
	fn := func() int { i++; return i }
	tick := make(chan bool, 1)

	fnThrottled := Throttle(fn, tick)

	fnThrottled()
	fnThrottled() // should be collapsed

	if i != 1 {
		t.Errorf("fnThrottled should have run once; got %d", i)
	}

	tick <- true // move things along
	fnThrottled()

	if i != 2 {
		t.Errorf("fnThrottled should have run twice; got %d", i)
	}
}

func TestCancellable(t *testing.T) {
	sender := make(chan int, 10)
	fn := func() chan int {
		ch := make(chan int, 10)

		go func() {
			for item := range sender {
				ch <- item
			}
		}()

		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())

	fnWithCancel := Cancellable(ctx, fn)
	out := fnWithCancel()

	sender <- 0
	sender <- 1

	cancel() // send cancellation

	sender <- 3 // should be ignored now

	got, ok := exhaustChannel(out)
	if !ok {
		t.Error("timeout when exahusting out channel")
		return
	}

	want := []int{0, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("function out chan should have returned %v before cancel; got %v", want, got)
	}
}

func TestGetQuickest(t *testing.T) {
	slow := make(chan int, 1)
	fast := make(chan int, 1)

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "target1":
			<-slow
			fmt.Fprintln(w, "Hello from target1")
		case "mirror1":
			<-fast
			fmt.Fprintln(w, "Hello from mirror1")
		case "mirror2":
			<-slow
			fmt.Fprintln(w, "Hello from mirror2")
		}

	}))
	defer testServer.Close()

	asURL := func(s string) url.URL {
		res, _ := url.Parse(s) // ignore error but don't do this for real!
		return *res
	}

	out := GetQuickest(asURL(testServer.URL+"/target1"), asURL(testServer.URL+"/mirror1"), asURL(testServer.URL+"/mirror2"))
	fast <- 0 // ensure fast wins!

	resp, ok := readCh(out)
	if !ok {
		t.Error("timeout reading from result channel")
		return
	}

	got, _ := io.ReadAll(resp.Response.Body)
	want := "Hello from mirror1"

	if string(got) != want {
		t.Errorf("want %s; got %s", want, string(got))
	}
}

func TestFanOut(t *testing.T) {
	t.SkipNow() // TODO implement this test!
}

func readCh[T any](ch <-chan T) (T, bool) {
	select {
	case item, ok := <-ch:
		if !ok { // chan closed
			return item, false
		}

		return item, true
	case <-time.After(time.Millisecond * 50):
		var notFound T
		return notFound, false
	}
}

func exhaustChannel[T any](ch <-chan T) ([]T, bool) {
	var out []T

	for {
		select {
		case item, ok := <-ch:
			if !ok { // chan closed
				return out, true
			}

			out = append(out, item)
		case <-time.After(time.Millisecond * 50):
			return out, false
		}
	}
}
