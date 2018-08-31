package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

const (
	CommandOrder string = "ORDER"
)

func main() {
	watcher := observer.Observer{

		// Register a handler function for every next available item.
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},

		// Register a handler for any emitted error.
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},

		// Register a handler when a stream is completed.
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	it, _ := iterable.New([]interface{}{1, 2, 3, 4, errors.New("bang"), 5})
	source := observable.From(it)
	sub := source.Subscribe(watcher)

	// wait for the channel to emit a Subscription
	<-sub
	//processLines(os.Stdin)
}

func processLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		newline := scanner.Text()
		sline := strings.Split(newline, " ")
		if len(sline) >= Localconfig.MinWordsCommand {
			if sline[0] == CommandOrder {

			}
		}
	}
}
