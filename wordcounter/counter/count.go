package counter

import (
	"bufio"
	"fmt"
	"io"
)

func Count(s io.Reader) int {
	// A scanner is used to read text from a reader (such as a file)
	scanner := bufio.NewScanner(s)
	fmt.Println("Enter words:")
	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)
	// defining a counter
	wc := 0
	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	return wc
}
