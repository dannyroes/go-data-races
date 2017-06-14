package main

import (
	"fmt"
	"sync"
	"time"
)

// MUTEX OMIT
var appState string
var stateMutex *sync.RWMutex // HL

func main() {
	stateMutex = &sync.RWMutex{} // HL

	stateMutex.Lock() // HL
	appState = "Starting"
	stateMutex.Unlock() // HL

	writeLog("Initializing services")
	time.Sleep(200 * time.Millisecond)

	// MUTEXEND OMIT

	stateMutex.Lock() // HL
	appState = "Running"
	stateMutex.Unlock() // HL

	writeLog("Working...")
	time.Sleep(200 * time.Millisecond)

	stateMutex.Lock()
	appState = "Ending"
	stateMutex.Unlock()

	writeLog("Shutting down services")
	time.Sleep(200 * time.Millisecond)
}

// START OMIT

func writeLog(message string) {
	go func() {
		stateMutex.RLock() // HL
		fmt.Printf("App State: %v; Message: %v\n", appState, message)
		stateMutex.RUnlock() // HL
	}()
}

// END OMIT
