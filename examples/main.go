package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling

	"github.com/Alfex4936/dkssud"
)

func main() {
	// Start the pprof HTTP server in a separate goroutine
	go func() {
		log.Println("Starting pprof server on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Call the function multiple times with different inputs
	inputs := []string{
		"rksekekfk akqkTkdk wkckzkxk vkvkvkgkgk", // Input 1
		"QwertyHangul simple test",               // Input 2
		"zzxxccvvbbaa ddssff",                    // Input 3
		"rkskslrkTksk",                           // Input 4 (testing specific cases)
	}

	// Increase iterations to generate more measurable CPU activity
	for i := 0; i < 10000000; i++ { // Increase iteration count
		for _, input := range inputs {
			dkssud.QwertyToHangul(input)
		}
	}

	// Keep the program running to allow pprof to gather data
	select {} // Block the main goroutine to keep the server running
}
