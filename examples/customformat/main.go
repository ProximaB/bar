package main

import (
	"fmt"
	"time"

	"github.com/superhawk610/bar"
)

func main() {
	n := 2000
	b := bar.NewWithOpts(
		bar.WithDimensions(n, 30),
		bar.WithFormat(" loading... :percent :bar :rate/s :eta "),
	)

	fmt.Println()
	fmt.Println()

	for i := 0; i < n; i++ {
		b.Tick()
		time.Sleep(500 * time.Millisecond)
	}

	b.Done()

	fmt.Println()
	fmt.Println()
}
