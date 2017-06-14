package main

import (
	"fmt"
)

func main() {
	// main OMIT
	work1 := make(map[int]string)
	for x := 10000; x < 20000; x++ {
		work1[x] = "unprocessed"
	}
	work2 := make(map[int]string)
	for x := 20000; x < 30000; x++ {
		work2[x] = "unprocessed"
	}

	workerChan := make(chan map[int]string)
	out1 := checkPrimes(workerChan)
	out2 := checkPrimes(workerChan)
	workerChan <- work1
	workerChan <- work2

	result := <-out1
	result2 := <-out2
	for num, status := range result2 {
		result[num] = status
	}

	close(workerChan)
	// endmain OMIT
	fmt.Printf("%v\n", result[13735])
	fmt.Printf("%v\n", result[23673])

	fmt.Printf("%v\n", result[19157])
	fmt.Printf("%v\n", result[29989])
}

// primes OMIT
func checkPrimes(workChan <-chan map[int]string) <-chan map[int]string {
	out := make(chan map[int]string)
	go func() {
		defer close(out)
		for work := range workChan {
			for num := range work {
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
			out <- work
		}
	}()
	return out
}

// endprimes OMIT
