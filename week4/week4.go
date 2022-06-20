package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Concurrency in Go!

// CPUs, Threads
// * what are the differences?
// * how do our common languages use threads? (Scala/Node)
//   NODE: callbacks, promises, async/await
//   Scala: lots! (but typically Future, ExecutionContext, 'async' libraries)
// What happens if you 'block'?

// https://go.dev/ref/mem
// 'Do not communicate by sharing memory; instead, share memory by communicating.'

func main() {
	example := flag.Int("example", 0, "choose an example to run!")
	flag.Parse()

	switch *example {
	case 1:
		race()
	case 2:
		top10()
	case 3:
		top10Par()
	case 4:
		launch()
	case 5:
		buffering()
	case 6:
		top10Timeouts()
	default:
		fmt.Println("Specify an example using '-example [n]'.")
	}
}

// Questions:
// * what is wrong with this function? (hint: try using the -race flag)
// * What do we think 'go' actually does? What if x++ were a slow/blocking operation?
func race() {
	x := 0

	go func() {
		x++
	}()

	fmt.Println(x)
}

// What about error handling, timeouts, cancellation?
// How can we improve this function?
func top10() {
	stories, err := getHNTopStories()
	check(err)

	titles := []string{}
	for _, story := range stories[0:10] {
		item, err := getHNItem(story)
		check(err)

		titles = append(titles, item.Title)
	}

	fmt.Printf("%s\n", strings.Join(titles, "\n"))
}

// Questions:
// * why doesn't this terminate?
// * how could we improve this function?
func top10Par() {
	ch := make(chan HNItem)

	stories, err := getHNTopStories()
	check(err)

	for _, story := range stories[0:10] {
		go func(ID int) {
			item, err := getHNItem(ID)
			check(err)

			ch <- item
		}(story)
	}

	for item := range ch {
		fmt.Printf("%s\n", item.Title)
	}
}

func launch() {
	abort := make(chan bool)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- true
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	select {
	case <-abort:
		fmt.Println("Launch aborted!")
	case <-time.After(time.Second * 3):
		fmt.Println("We have liftoff!")
		return
	}
}

// What does this do?
// Why might buffering useful?
func buffering() {
	ch := make(chan int, 1) // channel can only have one item at a time

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

// TODO implement a version of top10 that only waits for each item for a max of
// 500ms.
// Hint: use the `getHNItemCh` function as input here and select on that
// and a timeout.
func top10Timeouts() {
	ch := make(chan HNItem)

	stories, err := getHNTopStories()
	check(err)

	for _, story := range stories[0:10] {
		go func(ID int) {
			select {
			case <-time.After(time.Millisecond * 500):
				fmt.Printf("error: timeout for %d\n", ID)
				return
			case item := <-getHNItemCh(ID):
				check(err)
				ch <- item
			}
		}(story)
	}

	for item := range ch {
		fmt.Printf("%s\n", item.Title)
	}
}

// TODO - the example above is a bit messy. I'll send some (optional) homework
// for anyone keen to explore further.

// Helper code below...

type HNItem struct {
	Title string `json:"title"`
}

func getHNTopStories() ([]int, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var items []int
	err = json.Unmarshal(data, &items)

	return items, err
}

func getHNItem(ID int) (HNItem, error) {
	var item HNItem

	itemURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", ID)
	resp, err := http.Get(itemURL)
	if err != nil {
		return item, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}

	err = json.Unmarshal(data, &item)
	return item, err
}

func getHNItemCh(ID int) chan HNItem {
	ch := make(chan HNItem)

	go func() {
		item, _ := getHNItem(ID)                                     // ignoring error for brevity of example only
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(200))) // for fun
		ch <- item
	}()

	return ch
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
