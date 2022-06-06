package week3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFizzBuzz(t *testing.T) {
	buf := &bytes.Buffer{}

	FizzBuzz(buf, 10)

	got, _ := io.ReadAll(buf)

	want := `0
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
Fizz Buzz`

	if string(got) != want {
		t.Error(cmp.Diff(string(got), want))
	}
}

func TestFind(t *testing.T) {
	type example[T any] struct {
		items []T
		pred  func(T) bool
		want  T
	}

	items := []string{"foo", "bar", "Baz"}

	examples := []example[string]{
		{items: items, want: "bar", pred: func(elem string) bool { return elem == "bar" }},
	}

	for _, example := range examples {
		if got, _ := Find(example.items, example.pred); got != example.want {
			t.Errorf("got %s, want %s", got, example.want)
		}
	}
}

func TestGetNews(t *testing.T) {
	want := []Item{
		{
			Title: "Twelve Years of Go",
			URL:   "https://go.dev/blog/12years",
		},
		{
			Title: "Fuzzing is Beta Ready",
			URL:   "https://go.dev/blog/fuzz-beta",
		},
	}

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/news" {
			asJSON, _ := json.Marshal(want)
			w.Write(asJSON)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	}))
	defer ts1.Close()

	got, err := GetNews(ts1.URL)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error(cmp.Diff(got, want))
	}
}

func TestGetNewsHTTPError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	_, err := GetNews(ts.URL)

	if err == nil || !strings.Contains(err.Error(), "http request failed") {
		t.Errorf("expected HTTP error but got: %v", err)
	}
}

func TestGetNewsJSONError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{ "unexpected": "json" `)
	}))
	defer ts.Close()

	_, err := GetNews(ts.URL)

	if err == nil || !strings.Contains(err.Error(), "json parsing failed") {
		t.Errorf("expected JSON error but got: %v", err)
	}
}
