package main

import (
	"fmt"
	"time"
)

// START OMIT
var appState string

func main() {
	appState = "Starting"
	writeLog("Initializing services")
	time.Sleep(200 * time.Millisecond)

	appState = "Running" // HLrace
	writeLog("Working...")
	time.Sleep(200 * time.Millisecond)

	appState = "Ending"
	writeLog("Shutting down services")
	time.Sleep(200 * time.Millisecond)
}

func writeLog(message string) {
	go func() {
		fmt.Printf("App State: %v; Message: %v\n", appState, message) // HLrace
	}()
}

// END OMIT
