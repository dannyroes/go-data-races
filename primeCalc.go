package main

import (
	"fmt"
	"time"
)

// main OMIT
func main() {
	work := make(map[int]string)
	for x := 10000; x < 30000; x++ {
		work[x] = "unprocessed"
	}

	go checkPrimes(work, 10000, 10000)
	go checkPrimes(work, 20000, 10000)

	time.Sleep(200 * time.Millisecond)

	fmt.Printf("%v\n", work[13735])
	fmt.Printf("%v\n", work[23673])

	fmt.Printf("%v\n", work[19157])
	fmt.Printf("%v\n", work[29989])
}

// endmain OMIT
// primes OMIT
func checkPrimes(work map[int]string, startNum, length int) {
	for num := startNum; num < startNum+length; num++ {
		checkTo := num / 2
		for y := 2; y <= checkTo; y++ {
			if num%y == 0 {
				work[num] = "not prime"
				break
			}
		}

		if work[num] == "unprocessed" {
			work[num] = "prime"
		}
	}
}

// endprimes OMIT
