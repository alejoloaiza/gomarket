package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	processLines(os.Stdin)
}

func processLines(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		newline := scanner.Text()
		sline := strings.Split(newline, " ")
		if len(sline) > 3 {

		}
	}
}
