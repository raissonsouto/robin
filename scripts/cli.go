package scripts

import (
	"flag"
	"fmt"
	"os"
)

func initCli() {
	var (
		urlFlag  string
		routines int
		maxSize  int
		smooth   bool
		detach   bool
		help     bool
		//joinRaid   string
		//createRaid string
	)
	flag.StringVar(&urlFlag, "url", "", "The base URL to fuzz")
	flag.IntVar(&routines, "routines", 50, "Number of concurrent requests to make")
	flag.IntVar(&maxSize, "maxSize", 10, "Maximum size of response to accept")
	flag.BoolVar(&smooth, "smooth", false, "Smooth progress bar (useful for slow connections)")
	flag.BoolVar(&detach, "detach", false, "Silent mode (only print successful results)")
	flag.BoolVar(&help, "help", false, "Show help message")
	//flag.StringVar(&joinRaid, "join", "", "Join in a raid with given URL before sending requests")
	//flag.StringVar(&createRaid, "create", "", "Give a link to create a raid before sending requests")
	flag.Parse()

	// Show help message
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// Check required flag
	if urlFlag == "" {
		printErrorUrlNotSpecified()
		os.Exit(1)
	}

	// Print configuration
	if !detach {
		fmt.Println("Starting URL fuzzer with the following configuration:")
		fmt.Printf("  Base URL: %s\n", urlFlag)
		fmt.Printf("  Concurrent requests: %d\n", routines)
		fmt.Printf("  Maximum response size: %d bytes\n", maxSize)
		fmt.Printf("  Smooth progress bar: %v\n", smooth)
		fmt.Printf("  Silent mode: %v\n", silent)
		fmt.Printf("  Join URL: %s\n", joinRaid)
		fmt.Println()
	}

}
