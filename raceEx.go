package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	for x := 0; x < 5; x++ {
		go func() {
			fmt.Printf("Loop number %d\n", x)
		}()
	}
	time.Sleep(time.Millisecond)
}

// END OMIT
